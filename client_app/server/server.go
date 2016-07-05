package server

import (
	"flag"
	"fmt"
	"io"
	"net"

	pb "github.com/alvindaiyan/goTalk2/proto/client_proto"
	"google.golang.org/grpc"

	"google.golang.org/grpc/grpclog"
)

var (
	port       = flag.Int("port", 10000, "The server port")
	grpcServer *grpc.Server
)

type chat_server struct {
}

func (*chat_server) Chat(stream pb.Chat_ChatServer) error {
	for {
		in, err := stream.Recv()
		// end of the streaming
		if err == io.EOF {
			grpclog.Println("server -- finished stream")
			return nil
		}
		if err != nil {
			grpclog.Printf("returned with error %v", err)
			return err
		}
		content := in.Content
		title := in.Title
		if title == "" {
			title = "Unknown"
		}
		grpclog.Printf("server -- received message:\n%v: %v", title, content)
		revMsg := "received"
		stream.Send(&pb.Msg{Content: revMsg})
	}
}

func Shutdown() {
	grpclog.Println("shutdown server")
	grpcServer.Stop()
}

func InitChatServer() {
	grpclog.Println("start server...")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		grpclog.Fatal("failed to listen: %v", err)
	}

	grpcServer = grpc.NewServer()
	pb.RegisterChatServer(grpcServer, new(chat_server))
	grpcServer.Serve(lis)
	grpclog.Println("server shutdown...")

}
