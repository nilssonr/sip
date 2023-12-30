package sip

import (
	"strings"
	"sync"
)

// type Message interface {}

type Message interface {
	Accept() (*Accept, bool)
	AcceptEncoding() (*AcceptEncoding, bool)
	AcceptLanguage() (*AcceptLanguage, bool)
	AlertInfo() (*AlertInfo, bool)
	Allow() (*Allow, bool)
	AuthenticationInfo() (*AuthenticationInfo, bool)
	Authorization() (*Authorization, bool)
	CallID() (*CallID, bool)
	CallInfo() (*CallInfo, bool)
	Contact() (*Contact, bool)
	ContentDisposition() (*ContentDisposition, bool)
	ContentEncoding() (*ContentEncoding, bool)
	ContentLanguage() (*ContentLanguage, bool)
	ContentLength() (*ContentLength, bool)
	ContentType() (*ContentType, bool)
	CSeq() (*CSeq, bool)
	Date() (*Date, bool)
	ErrorInfo() (*ErrorInfo, bool)
	Expires() (*Expires, bool)
	From() (*From, bool)
	InReplyTo() (*InReplyTo, bool)
	MaxForwards() (*MaxForwards, bool)
	MinExpires() (*MinExpires, bool)
	MimeVersion() (*MimeVersion, bool)
	Organization() (*Organization, bool)
	Priority() (*Priority, bool)
	ProxyAuthenticate() (*ProxyAuthenticate, bool)
	ProxyAuthorization() (*ProxyAuthorization, bool)
	ProxyRequire() (*ProxyRequire, bool)
	RecordRoute() (*RecordRoute, bool)
	ReplyTo() (*ReplyTo, bool)
	Require() (*Require, bool)
	RetryAfter() (*RetryAfter, bool)
	Route() (*Route, bool)
	Server() (*Server, bool)
	Subject() (*Subject, bool)
	Supported() (*Supported, bool)
	Timestamp() (*Timestamp, bool)
	To() (*To, bool)
	Unsupported() (*Unsupported, bool)
	UserAgent() (*UserAgent, bool)
	Via() ([]*Via, bool)
	Warning() (*Warning, bool)
	WWWAuthenticate() (*WWWAuthenticate, bool)

	Method() string

	AppendHeader(header Header)
	GetHeaders(name string) []Header
}

func IsRequest(msg Message) bool {
	return msg.Method() == ""
}

type defaultMessage struct {
	method string
	host   string
	port   string
	*headers
}

// Accept implements Message.
func (msg defaultMessage) Accept() (*Accept, bool) {
	return getHeader[Accept]("accept", msg)
}

// AcceptEncoding implements Message.
func (msg defaultMessage) AcceptEncoding() (*AcceptEncoding, bool) {
	return getHeader[AcceptEncoding]("accept-encoding", msg)
}

// AcceptLanguage implements Message.
func (msg defaultMessage) AcceptLanguage() (*AcceptLanguage, bool) {
	return getHeader[AcceptLanguage]("accept-language", msg)
}

// AlertInfo implements Message.
func (msg defaultMessage) AlertInfo() (*AlertInfo, bool) {
	return getHeader[AlertInfo]("alert-info", msg)
}

func (msg defaultMessage) Allow() (*Allow, bool) {
	return getHeader[Allow]("allow", msg)
}

// AuthenticationInfo implements Message.
func (msg defaultMessage) AuthenticationInfo() (*AuthenticationInfo, bool) {
	return getHeader[AuthenticationInfo]("authentication-info", msg)
}

// Authorization implements Message.
func (msg defaultMessage) Authorization() (*Authorization, bool) {
	return getHeader[Authorization]("authorization", msg)
}

func (msg defaultMessage) CallID() (*CallID, bool) {
	return getHeader[CallID]("call-id", msg)
}

// CallInfo implements Message.
func (msg defaultMessage) CallInfo() (*CallInfo, bool) {
	return getHeader[CallInfo]("call-info", msg)
}

// Contact implements Message.
func (msg defaultMessage) Contact() (*Contact, bool) {
	return getHeader[Contact]("contact", msg)
}

// ContentDisposition implements Message.
func (msg defaultMessage) ContentDisposition() (*ContentDisposition, bool) {
	return getHeader[ContentDisposition]("content-disposition", msg)
}

// ContentEncoding implements Message.
func (msg defaultMessage) ContentEncoding() (*ContentEncoding, bool) {
	return getHeader[ContentEncoding]("content-encoding", msg)
}

// ContentLanguage implements Message.
func (msg defaultMessage) ContentLanguage() (*ContentLanguage, bool) {
	return getHeader[ContentLanguage]("content-language", msg)
}

// ContentLength implements Message.
func (msg defaultMessage) ContentLength() (*ContentLength, bool) {
	return getHeader[ContentLength]("content-length", msg)
}

// ContentType implements Message.
func (msg defaultMessage) ContentType() (*ContentType, bool) {
	return getHeader[ContentType]("content-type", msg)
}

func (msg defaultMessage) CSeq() (*CSeq, bool) {
	return getHeader[CSeq]("cseq", msg)
}

// Date implements Message.
func (msg defaultMessage) Date() (*Date, bool) {
	return getHeader[Date]("date", msg)
}

// ErrorInfo implements Message.
func (msg defaultMessage) ErrorInfo() (*ErrorInfo, bool) {
	return getHeader[ErrorInfo]("error-info", msg)
}

// Expires implements Message.
func (msg defaultMessage) Expires() (*Expires, bool) {
	return getHeader[Expires]("expires", msg)
}

// From implements Message.
func (msg defaultMessage) From() (*From, bool) {
	return getHeader[From]("from", msg)
}

// InReplyTo implements Message.
func (msg defaultMessage) InReplyTo() (*InReplyTo, bool) {
	return getHeader[InReplyTo]("in-reply-to", msg)
}

// MaxForwards implements Message.
func (msg defaultMessage) MaxForwards() (*MaxForwards, bool) {
	return getHeader[MaxForwards]("max-forwards", msg)
}

// MimeVersion implements Message.
func (msg defaultMessage) MimeVersion() (*MimeVersion, bool) {
	return getHeader[MimeVersion]("mime-version", msg)
}

// MinExpires implements Message.
func (msg defaultMessage) MinExpires() (*MinExpires, bool) {
	return getHeader[MinExpires]("min-expires", msg)
}

// Organization implements Message.
func (msg defaultMessage) Organization() (*Organization, bool) {
	return getHeader[Organization]("organization", msg)
}

// Priority implements Message.
func (msg defaultMessage) Priority() (*Priority, bool) {
	return getHeader[Priority]("priority", msg)
}

// ProxyAuthenticate implements Message.
func (msg defaultMessage) ProxyAuthenticate() (*ProxyAuthenticate, bool) {
	return getHeader[ProxyAuthenticate]("proxy-authenticate", msg)
}

// ProxyAuthorization implements Message.
func (msg defaultMessage) ProxyAuthorization() (*ProxyAuthorization, bool) {
	return getHeader[ProxyAuthorization]("proxy-authorization", msg)
}

// ProxyRequire implements Message.
func (msg defaultMessage) ProxyRequire() (*ProxyRequire, bool) {
	return getHeader[ProxyRequire]("proxy-require", msg)
}

// RecordRoute implements Message.
func (msg defaultMessage) RecordRoute() (*RecordRoute, bool) {
	return getHeader[RecordRoute]("record-route", msg)
}

// ReplyTo implements Message.
func (msg defaultMessage) ReplyTo() (*ReplyTo, bool) {
	return getHeader[ReplyTo]("reply-to", msg)
}

// Require implements Message.
func (msg defaultMessage) Require() (*Require, bool) {
	return getHeader[Require]("require", msg)
}

// RetryAfter implements Message.
func (msg defaultMessage) RetryAfter() (*RetryAfter, bool) {
	return getHeader[RetryAfter]("retry-after", msg)
}

// Route implements Message.
func (msg defaultMessage) Route() (*Route, bool) {
	return getHeader[Route]("route", msg)
}

// Server implements Message.
func (msg defaultMessage) Server() (*Server, bool) {
	return getHeader[Server]("server", msg)
}

// Subject implements Message.
func (msg defaultMessage) Subject() (*Subject, bool) {
	return getHeader[Subject]("subject", msg)
}

// Supported implements Message.
func (msg defaultMessage) Supported() (*Supported, bool) {
	return getHeader[Supported]("supported", msg)
}

// Timestamp implements Message.
func (msg defaultMessage) Timestamp() (*Timestamp, bool) {
	return getHeader[Timestamp]("timestamp", msg)
}

// To implements Message.
func (msg defaultMessage) To() (*To, bool) {
	return getHeader[To]("to", msg)
}

// Unsupported implements Message.
func (msg defaultMessage) Unsupported() (*Unsupported, bool) {
	return getHeader[Unsupported]("unsupported", msg)
}

// UserAgent implements Message.
func (msg defaultMessage) UserAgent() (*UserAgent, bool) {
	return getHeader[UserAgent]("user-agent", msg)
}

// Via implements Message.
func (msg defaultMessage) Via() ([]*Via, bool) {
	headers := msg.GetHeaders("via")
	vias := make([]*Via, 0)
	for _, header := range headers {
		if via, ok := header.(*Via); ok {
			vias = append(vias, via)
		}
	}
	return vias, len(vias) > 0
}

// WWWAuthenticate implements Message.
func (msg defaultMessage) WWWAuthenticate() (*WWWAuthenticate, bool) {
	return getHeader[WWWAuthenticate]("www-authenticate", msg)
}

// Warning implements Message.
func (msg defaultMessage) Warning() (*Warning, bool) {
	return getHeader[Warning]("warning", msg)
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

func getHeader[T any](name string, msg Message) (*T, bool) {
	headers := msg.GetHeaders(name)

	if len(headers) == 0 {
		return nil, false
	}

	allow, ok := headers[0].(T)
	if !ok {
		return nil, false
	}

	return &allow, true
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
