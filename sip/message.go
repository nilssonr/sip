package sip

import (
	"strings"
	"sync"
)

// type Message interface {}

type Message interface {
	Method() string
	CSeq() (*CSeq, bool)
	Allow() (*Allow, bool)

	AppendHeader(header Header)
	GetHeaders(name string) []Header
}

type defaultMessage struct {
	method string
	host   string
	port   string
	*headers
}

func newMessage(rl *RequestLine) Message {
	return defaultMessage{
		method: rl.Method,
		host:   rl.Host,
		port:   rl.Port,
		headers: &headers{
			headers: map[string][]Header{},
			mu:      sync.RWMutex{},
		},
	}
}

func (msg defaultMessage) Method() string {
	return msg.method
}

// Allow implements Message.
func (msg defaultMessage) Allow() (*Allow, bool) {
	headers := msg.GetHeaders("allow")
	if len(headers) == 0 {
		return nil, false
	}
	allow, ok := headers[0].(*Allow)
	if !ok {
		return nil, false
	}
	return allow, true
}

// CSeq implements Message.
func (msg defaultMessage) CSeq() (*CSeq, bool) {
	headers := msg.GetHeaders("cseq")
	if len(headers) == 0 {
		return nil, false
	}
	cseq, ok := headers[0].(*CSeq)
	if !ok {
		return nil, false
	}
	return cseq, true
}

type headers struct {
	headers map[string][]Header
	mu      sync.RWMutex
}

func (h *headers) AppendHeader(header Header) {
	h.mu.Lock()
	defer h.mu.Unlock()

	name := strings.ToLower(header.Name())
	if _, ok := h.headers[name]; !ok {
		h.headers[name] = []Header{header}
	} else {
		h.headers[name] = append(h.headers[name], header)
	}
}

func (h *headers) GetHeaders(name string) []Header {
	name = strings.ToLower(name)
	h.mu.Lock()
	defer h.mu.Unlock()

	if headers, exists := h.headers[name]; exists {
		return headers
	}
	return []Header{}
}
