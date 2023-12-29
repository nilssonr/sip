package main

import (
	"fmt"

	"github.com/nilssonr/sip/transport"
)

func main() {
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
