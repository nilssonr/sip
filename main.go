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

			if callID, ok := msg.CallID(); ok {
				fmt.Println("CallID: ", callID)
			}

			if allow, ok := msg.Allow(); ok {
				fmt.Println("Allow: ", allow)
			}

			if via, ok := msg.Via(); ok {
				for _, v := range via {
					fmt.Println("VIA: ", v.Host)
				}
			}
		}
	}()
	transl.Listen("tcp", "0.0.0.0:5060")

	select {}
}
