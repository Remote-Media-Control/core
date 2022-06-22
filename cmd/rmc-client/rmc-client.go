package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/Remote-Media-Control/core/proto"

	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr string
	port int
	cl   pb.RemotePlayerClient
)

func runReq(rq *pb.RunRequest) {
	sendRequest(func(ctx context.Context, cl pb.RemotePlayerClient) error {
		r, err := cl.Run(ctx, rq)
		log.Printf("Ret: %s", r.String())
		return err
	})
}

func infoReq() {
	sendRequest(func(ctx context.Context, cl pb.RemotePlayerClient) error {
		r, err := cl.Info(ctx, &pb.Empty{})
		log.Printf("Ret: %s", r.String())
		return err
	})
}

func sendRequest(cb func(context.Context, pb.RemotePlayerClient) error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", addr, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	cl = pb.NewRemotePlayerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// r, err := cl.Run(ctx, rq)
	err = cb(ctx, cl)
	if err != nil {
		log.Fatalf("could not send run req: %v", err)
	}

}

func makeApp() *cli.App {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "addr",
			Value:       "127.0.0.1",
			Usage:       "Server addr",
			Destination: &addr,
		},
		cli.IntFlag{
			Name:        "port",
			Value:       50051,
			Usage:       "Server port",
			Destination: &port,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "run some key",
			Subcommands: []cli.Command{
				{
					Name:    "next",
					Aliases: []string{"n"},
					Usage:   "next",
					Action: func(c *cli.Context) error {
						runReq(&pb.RunRequest{Key: pb.Key_NEXT})
						return nil
					},
				},
				{
					Name:    "prev",
					Aliases: []string{"pr"},
					Usage:   "previous",
					Action: func(c *cli.Context) error {
						runReq(&pb.RunRequest{Key: pb.Key_PREV})
						return nil
					},
				},
				{
					Name:    "play",
					Aliases: []string{"p"},
					Usage:   "play",
					Action: func(c *cli.Context) error {
						runReq(&pb.RunRequest{Key: pb.Key_PLAY})
						return nil
					},
				},
				{
					Name:    "pause",
					Aliases: []string{"pa"},
					Usage:   "pause",
					Action: func(c *cli.Context) error {
						runReq(&pb.RunRequest{Key: pb.Key_PAUSE})
						return nil
					},
				},
			},
		},
		{
			Name:    "info",
			Aliases: []string{"i"},
			Usage:   "info about current playback",
			Action: func(c *cli.Context) error {
				infoReq()
				return nil
			},
		},
	}

	return app
}

func main() {
	err := makeApp().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
