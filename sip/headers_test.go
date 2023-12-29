package sip

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestParseCSeq(t *testing.T) {
// 	tests := []struct {
// 		Input    []byte
// 		Expected *CSeq
// 	}{
// 		{
// 			Input:    []byte("123 INVITE"),
// 			Expected: &CSeq{Sequence: 123, Method: "INVITE"},
// 		},
// 		{
// 			Input:    []byte("1                            INVITE"),
// 			Expected: &CSeq{Sequence: 1, Method: "INVITE"},
// 		},
// 	}

// 	for _, test := range tests {
// 		header, err := parseCSeq(test.Input)
// 		assert.Nil(t, err)
// 		assert.Equal(t, header.Sequence, test.Expected.Sequence)
// 		assert.Equal(t, header.Method, test.Expected.Method)
// 	}
// }

// func BenchmarkParseCSeq(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		parseCSeq([]byte("123 INVITE"))
// 	}
// }

// func TestParseContact(t *testing.T) {
// 	tests := []struct {
// 		Input    []byte
// 		Expected *Contact
// 	}{
// 		{
// 			Input: []byte("<sip:bob@192.0.2.4>"),
// 			Expected: &Contact{
// 				Scheme: "sip",
// 				User:   "bob",
// 				Host:   "192.0.2.4",
// 			},
// 		},
// 		{
// 			Input: []byte("Lee.Foote <sips:lee.foote@example.com>;tsp=wss"),
// 			Expected: &Contact{
// 				Scheme:    "sips",
// 				Name:      "Lee.Foote",
// 				User:      "lee.foote",
// 				Host:      "example.com",
// 				Transport: "wss",
// 			},
// 		},
// 		{
// 			Input: []byte("\"Mr. Watson\" <tel:watson@worcester.bell-telephone.com>;q=0.7; expires=3600;transport=tcp"),
// 			Expected: &Contact{
// 				Scheme:    "tel",
// 				Name:      "Mr. Watson",
// 				User:      "watson",
// 				Host:      "worcester.bell-telephone.com",
// 				Transport: "tcp",
// 				Expires:   3600,
// 				Q:         "0.7",
// 			},
// 		},
// 	}

// 	for _, test := range tests {
// 		header, err := parseContact(test.Input)
// 		assert.Nil(t, err)
// 		assert.Equal(t, header.Scheme, test.Expected.Scheme)
// 		assert.Equal(t, header.Name, test.Expected.Name)
// 		assert.Equal(t, header.User, test.Expected.User)
// 		assert.Equal(t, header.Host, test.Expected.Host)
// 		assert.Equal(t, header.Port, test.Expected.Port)
// 		assert.Equal(t, header.Transport, test.Expected.Transport)
// 		assert.Equal(t, header.Q, test.Expected.Q)
// 		assert.Equal(t, header.Expires, test.Expected.Expires)
// 	}
// }

// func BenchmarkParseContact(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		parseContact([]byte("\"Mr. Watson\" <sip:watson@worcester.bell-telephone.com>;q=0.7; expires=3600"))
// 	}
// }

// func TestParseFrom(t *testing.T) {
// 	tests := []struct {
// 		Input    []byte
// 		Expected *From
// 	}{
// 		{
// 			Input: []byte("Bob <sips:bob@biloxi.example.com>;tag=ja743ks76zlflH"),
// 			Expected: &From{
// 				Scheme: "sips",
// 				Name:   "Bob",
// 				User:   "bob",
// 				Host:   "biloxi.example.com",
// 				Port:   "",
// 				Tag:    "ja743ks76zlflH",
// 			},
// 		},
// 		{
// 			Input: []byte("Bob <sips:bob@biloxi.example.com:5060>;tag=JueHGuidj28dfga"),
// 			Expected: &From{
// 				Scheme: "sips",
// 				Name:   "Bob",
// 				User:   "bob",
// 				Host:   "biloxi.example.com",
// 				Port:   "5060",
// 				Tag:    "JueHGuidj28dfga",
// 			},
// 		},
// 		{
// 			Input: []byte(`"J Rosenberg"       <sip:jdrosen@example.com>
// 			;
// 			tag=98asjd8`),
// 			Expected: &From{
// 				Scheme: "sip",
// 				Name:   "J Rosenberg",
// 				User:   "jdrosen",
// 				Host:   "example.com",
// 				Tag:    "98asjd8",
// 			},
// 		},
// 	}

// 	for _, test := range tests {
// 		header, err := parseFrom(test.Input)
// 		assert.Nil(t, err)
// 		assert.Equal(t, header.Scheme, test.Expected.Scheme)
// 		assert.Equal(t, header.Name, test.Expected.Name)
// 		assert.Equal(t, header.User, test.Expected.User)
// 		assert.Equal(t, header.Port, test.Expected.Port)
// 		assert.Equal(t, header.Tag, test.Expected.Tag)
// 		assert.Equal(t, header.UserType, test.Expected.UserType)
// 	}
// }

// func BenchmarkFrom(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		parseFrom([]byte("Alice <sip:alice@atlanta.com>;tag=1928301774"))
// 	}
// }

// func TestParseRequestLine(t *testing.T) {
// 	tests := []struct {
// 		Input    []byte
// 		Expected *RequestLine
// 	}{
// 		{
// 			Input: []byte("INVITE sip:vivekg@chair-dnrc.example.com:5060 SIP/2.0"),
// 			Expected: &RequestLine{
// 				Method: "INVITE",
// 				Scheme: "sip",
// 				User:   "vivekg",
// 				Host:   "chair-dnrc.example.com",
// 				Port:   "5060",
// 			},
// 		},
// 		{
// 			Input: []byte("SIP/2.0 200 OK"),
// 			Expected: &RequestLine{
// 				Method:            "",
// 				StatusCode:        StatusOK,
// 				StatusDescription: StatusText(StatusOK),
// 			},
// 		},
// 	}

// 	for _, test := range tests {
// 		line, err := parseRequestLine(test.Input)
// 		assert.Nil(t, err)
// 		assert.Equal(t, line.Method, test.Expected.Method)
// 	}
// }
