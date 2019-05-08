package main

import (
	"fmt"
	"github.com/kitt-technology/ping/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
)

func main() {
	// Initialise connection with service B (ping)
	conn, err := grpc.Dial(os.Getenv("PING_SERVICE"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	pingClient := ping.NewPingClient(conn)
	defer conn.Close()

	// HTTP endpoint to test connection
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		resp, err := pingClient.Ping(context.Background(), &ping.PingRequest{Message: "ping"})
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Fprintf(w, resp.Message)
	})

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
