package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/Remote-Media-Control/core/actor"
	pb "github.com/Remote-Media-Control/core/proto"

	"google.golang.org/grpc"
)

var (
	addr = flag.String("addr", "127.0.0.1", "The server addr")
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedRemotePlayerServer
}

func (s *server) Run(ctx context.Context, in *pb.RunRequest) (*pb.Empty, error) {
	key := in.GetKey()
	log.Printf("Recv: %v", key)
	actor.PressKey(key)
	return &pb.Empty{}, nil
}

func (s *server) Info(ctx context.Context, in *pb.Empty) (*pb.InfoResponse, error) {
	return &pb.InfoResponse{
		SongName:   "unimplemented",
		ArtistName: "unimplemented",
		Status:     pb.PlayerStatus_UNKNOWN,
	}, nil
}

func runServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *addr, *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterRemotePlayerServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	flag.Parse()
	runServer()
}
