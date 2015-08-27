package client

import (
	"io"

	pb "github.com/goTalk2/proto/client_proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc/grpclog"

	"google.golang.org/grpc"
)

var (
	serverAddr *string
	title      string
)

func Chat(letters ...string) error {
	// get connection
	conn := connect(serverAddr)
	defer conn.Close()
	client := pb.NewChatClient(conn)

	stream, err := client.Chat(context.Background())
	if err != nil {
		grpclog.Println("%v.Chat(_) = _, %v", client, err) // better logging
		return err
	}

	// receive msg
	waitc := make(chan struct{})
	var recevieErr error
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done
				close(waitc)
				return
			}
			if err != nil {
				grpclog.Printf("Failed to receive a msg : %v", err) // need better logging
				recevieErr = err
				return
			}
			grpclog.Printf("client -- server status: %s", in.Content)
		}
	}()

	if recevieErr != nil {
		return recevieErr
	}

	// send msg
	for _, str := range letters {
		grpclog.Printf("client -- send msg: %v", str)
		if err := stream.Send(&pb.Msg{Content: str, Title: title}); err != nil {
			grpclog.Printf("%v.Send(%v) = %v", stream, str, err) // need better logging
			return err
		}
	}

	// close send
	stream.CloseSend()
	<-waitc
	return nil
}

func InitChatClient(t string, srvAddr *string) {
	title = t
	serverAddr = srvAddr
}

func connect(srvAddr *string) *grpc.ClientConn {
	conn, err := grpc.Dial(*srvAddr)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	grpclog.Println("client started...")
	return conn

}
