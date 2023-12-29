package main

import (
	"fmt"
	"log/slog"

	"github.com/nilssonr/sip/transport"
)

func main() {

	slog.Info("baws")
	slog.Group("hey")
	slog.Info("baw")

	transl := transport.NewLayer()
	go func() {
		for msg := range transl.Messages() {
			fmt.Println("Got message")
			fmt.Printf("Method: %s\n", msg.Method())
		}
	}()
	transl.Listen("tcp", "0.0.0.0:5060")

	select {}
}
