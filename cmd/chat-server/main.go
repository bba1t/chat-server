package main

import (
	"context"
	desc "github.com/bba1t/chat-server/pkg/chat_v1"
	"github.com/brianvoe/gofakeit"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"strconv"
)

const (
	grpcHost = "localhost:"
	grpcPort = 50051
)

type server struct {
	desc.UnimplementedChatV1Server
}

func main() {
	lis, err := net.Listen("tcp", grpcHost+":"+strconv.Itoa(grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Println("Create request, names: ", req.GetUsernames())
	return &desc.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) Delete(_ context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	log.Println("Delete request, id: ", req.GetId())
	return &empty.Empty{}, nil
}

func (s *server) SendMessage(_ context.Context, req *desc.SendMessageRequest) (*empty.Empty, error) {
	log.Println("SendMessage request: ", req.GetFrom(), req.GetText(), req.GetTimestamp())
	return &empty.Empty{}, nil
}
