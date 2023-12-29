package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/nilssonr/sip/sip"
)

func main() {
	ltcp, err := net.Listen("tcp4", "0.0.0.0:5060")
	if err != nil {
		panic(err)
	}
	defer ltcp.Close()

	for {
		conn, err := ltcp.Accept()
		if err != nil {
			panic(err)
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	reader := bufio.NewReader(conn)
	buffer := bytes.Buffer{}

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println("client hung up")
				break
			}

			panic(err)
		}

		buffer.Write(line)

		if strings.HasSuffix(buffer.String(), "\r\n\r\n") {
			fmt.Println(buffer.String())
			msg, err := sip.Parse(buffer.Bytes())
			if err != nil {
				panic(err)
			}

			fmt.Printf("%+v\n", msg)
			break
		}
	}

	// handleClient(conn)
}
