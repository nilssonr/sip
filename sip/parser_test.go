package sip

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	tests := []struct {
		Input    []byte
		Expected Message
	}{
		{
			Input: []byte("INVITE sip:bob@biloxi.com SIP/2.0\r\n" +
				"Via: SIP/2.0/UDP pc33.atlanta.com;branch=z9hG4bKnashds8\r\n" +
				"Max-Forwards: 70\r\n" +
				"To: Bob <sip:bob@biloxi.com>;user=bawsman;tag=h123\r\n" +
				"From: Alice <sip:alice@atlanta.com>;tag=1928301774\r\n" +
				"Call-ID: a84b4c76e66710\r\n" +
				"CSeq: 314159 INVITE\r\n" +
				"Contact: <sip:alice@pc33.atlanta.com>\r\n" +
				"Content-Type: application/sdp\r\n" +
				"Content-Length: 142\r\n`"),
			Expected: defaultMessage{},
		},
	}

	for _, test := range tests {
		_, err := Parse(test.Input)
		assert.Nil(t, err)
	}
}
