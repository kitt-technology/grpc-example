package main

import (
	"fmt"
	"github.com/timrcoulson/ping/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
)
func main() {
	conn, err := grpc.Dial(os.Getenv("PING_SERVICE"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	pingClient := ping.NewPingClient(conn)
	defer conn.Close()

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
