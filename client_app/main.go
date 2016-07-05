package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/alvindaiyan/goTalk2/client_app/client"
	"github.com/alvindaiyan/goTalk2/client_app/server"
)

var (
	msgc = make(chan string) // the message channel

	serverAddr = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
	myTitle    = flag.String("title", "", "The name show to your friend")
)

// an input from command line
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
	for {
		// start the app
		waitc := make(chan struct{}) // a wait lock

		// start the server thread
		go func() {
			fmt.Println("start the server")
			server.InitChatServer()
			defer close(waitc)
		}()

		// start the client thread
		go func() {
			for {
				msg := <-msgc // a message to send
				client.InitChatClient(*myTitle, serverAddr)
				err := client.Chat(msg)
				if err != nil {
					// restart the client
					fmt.Printf("send Err: %v", err)
				}
			}
		}()

		// start the input thread
		go input()

		<-waitc
		// finished in this round restart the app
		fmt.Println("restart the app")
	}
}
