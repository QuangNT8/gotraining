// Multithreaded Hello World server.
// Uses Goroutines.  We could also use channels (a native form of
// inproc), but I stuck to the example.
//
// Author:  Brendan Mc.
// Requires: http://github.com/alecthomas/gozmq

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"zmq/config"

	// zmq "github.com/alecthomas/gozmq"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	var wg sync.WaitGroup

	// load config
	err := config.NewConfig("./config.yml")
	if err != nil {
		panic(err)
	}

	getconfig := config.GetDbUrl()

	fmt.Println("Host: ", getconfig["Host"])
	fmt.Println("Name: ", getconfig["Name"])

	//  Frontend socket talks to clients over ipc
	frontend, _ := zmq.NewSocket(zmq.ROUTER)
	defer frontend.Close()
	frontend.Bind("ipc:///tmp/zmq_server.ipc")

	//  Backend socket talks to workers over inproc
	backend, _ := zmq.NewSocket(zmq.DEALER)
	defer backend.Close()
	backend.Bind("inproc://backend")

	// Launch pool of worker threads
	for i := 0; i != 5; i = i + 1 {
		go worker(&wg, i)
	}

	//  Connect backend to frontend via a proxy
	err = zmq.Proxy(frontend, backend, nil)
	fmt.Println("Proxy interrupted:", err)

	wg.Wait()
}

func worker(wg *sync.WaitGroup, workerID int) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	worker, _ := zmq.NewSocket(zmq.DEALER)
	defer worker.Close()
	worker.Connect("inproc://backend")
	wg.Add(1)
	for {
		select {
		case <-ctx.Done():
			cancel()
			fmt.Println("closing websocket")
			wg.Done()
			return
		default:
		}
		msg, err := worker.RecvMessage(0)
		if err == nil {
			_, flowInfo := pop(msg)
			jsonMap := make(map[string]interface{})
			err := json.Unmarshal([]byte(flowInfo[0]), &jsonMap)
			if err != nil {
				fmt.Println("Unexpected data:", flowInfo[0])
				panic(err)
			}
			for key, value := range jsonMap {
				if key == "flow" {
					data, err := json.Marshal(value)
					if err != nil {
						panic(err)
					}
					fmt.Println("Receive flow: ", data)

				} else if key == "interface" {
					fmt.Println("Receive flow: ", value)
				} else {
					fmt.Println("key: ", key)
					fmt.Println("Receive flow: ", value)
				}
			}
		}
	}

}

func pop(msg []string) (head, tail []string) {
	if msg[1] == "" {
		head = msg[:2]
		tail = msg[2:]
	} else {
		head = msg[:1]
		tail = msg[1:]
	}
	return
}
