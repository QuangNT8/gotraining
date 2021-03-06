package server

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"RESTovergRPC/directory"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func StartGRPC(wg *sync.WaitGroup, ctx context.Context, dbUrl map[string]string) {
	listen, err := net.Listen("tcp", "localhost:5566")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// register service
	grpcServer := grpc.NewServer()
	ds, err := NewDirectoryServer(dbUrl)
	defer ds.Close()
	if err != nil {
		log.Fatalf("couldn't connect to backend: %v", err)
	}
	directory.RegisterDirectoryServer(grpcServer, ds)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	_, cancel := context.WithCancel(ctx)
	defer cancel()
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")

			grpcServer.GracefulStop()
			cancel()
			wg.Done()
		}
	}()

	// start gRPC server
	log.Println("starting gRPC server...")
	grpcServer.Serve(listen)
}

func StartHTTP(wg *sync.WaitGroup) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Connect to the GRPC server
	conn, err := grpc.Dial("localhost:5566", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	// Register grpc-gateway
	rmux := runtime.NewServeMux()
	client := directory.NewDirectoryClient(conn)
	err = directory.RegisterDirectoryHandlerClient(ctx, rmux, client)
	if err != nil {
		log.Fatal(err)
	}
	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	server := http.Server{
		Addr:    "localhost:8081",
		Handler: rmux,
	}

	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")
			server.Shutdown(ctx)
			cancel()
			wg.Done()
		}
	}()
	log.Println("REST server ready...")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
