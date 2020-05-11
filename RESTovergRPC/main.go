package main

import (
	"RESTovergRPC/server"
	"context"
	"os"
	"sync"
	//"github.com/annp1987/RESTovergRPC/server"
)

var (
	host     = os.Getenv("DB_USERS_HOST")
	port     = os.Getenv("DB_USERS_PORT")
	user     = os.Getenv("DB_USERS_USER")
	dbname   = os.Getenv("DB_USERS_NAME")
	password = os.Getenv("DB_USERS_PASSWORD")
)

func Init() {
	if host == "" {
		host = "127.0.0.1"
	}
	if port == "" {
		port = "5432"
	}
	if user == "" {
		user = "postgres"
	}
	if dbname == "" {
		dbname = "quang_database"
	}
	if password == "" {
		password = "1"
	}
}

func main() {
	var wg sync.WaitGroup
	ctx := context.Background()
	// for test locally with go run main.go
	Init()
	dburl := map[string]string{
		"Host":     host,
		"Port":     port,
		"User":     user,
		"Type":     dbname,
		"Password": password,
	}
	wg.Add(1)
	go server.StartGRPC(&wg, ctx, dburl)
	wg.Add(1)
	go server.StartHTTP(&wg)

	// Block forever
	wg.Wait()

}
