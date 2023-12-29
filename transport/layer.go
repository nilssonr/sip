package transport

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

type layer struct {
	messages chan sip.Message
}

func NewLayer() layer {
	return layer{
		messages: make(chan sip.Message),
	}
}

// Listen implements Layer.
func (l *layer) Listen(network string, address string) error {
	listener, err := net.Listen(network, address)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			break
		}

		go func(r io.ReadWriter) {
			reader := bufio.NewReader(r)
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
					msg, err := sip.Parse(buffer.Bytes())
					if err != nil {
						panic(err)
					}

					l.messages <- msg
				}
			}
		}(conn)
	}

	return nil
}

// Messages implements Layer.
func (l *layer) Messages() <-chan sip.Message {
	return l.messages
}
