package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/goTalk2/client_app/client"
	"github.com/goTalk2/client_app/server"
)

var (
	msgc       = make(chan string)
	serverAddr = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
)

func input() {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		msgc <- text
	}
}

func main() {
	flag.Parse()
	fmt.Println("start the program")
	// start the app
	for {
		waitc := make(chan struct{}) // a wait lock

		// start the server thread
		go func() {
			server.InitChatServer()
			close(waitc)
		}()

		// start the client thread
		client.InitChatClient(serverAddr)

		go func() {
			for {
				msg := <-msgc // a message to send
				client.Chat(msg)
			}
			close(waitc) // unlock the main process and start over
		}()

		// start the input thread
		go input()
		<-waitc
		server.Shutdown()
		client.Shutdown()
		fmt.Println("restart the app")
	}
}
