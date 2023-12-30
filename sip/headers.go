package sip

type Header interface {
	Name() string
}

type RequestLine struct {
	Method            string
	Scheme            string
	StatusCode        int
	StatusDescription string
	User              string
	Host              string
	Port              string
	UserType          string
}

// Accept follows the syntax defined in [H14.1].  The semantics are also
// identical, with the exception that if no Accept header field is present, the
// server SHOULD assume a default value of application/sdp.
//
// An empty Accept header field means that no formats are acceptable.
//
// See: https://datatracker.ietf.org/doc/html/rfc3261#section-20.1
type Accept string

func (Accept) Name() string { return "Accept" }

// AcceptEncoding is similar to Accept, but restricts the content-codings [H3.5]
// that are acceptable in the response.  See [H14.3]. The semantics in SIP are
// identical to those defined in [H14.3].
//
// An empty Accept-Encoding header field is permissible.  It is equivalent to
// Accept-Encoding: identity, that is, only the identity encoding, meaning no
// encoding, is permissible.
//
// If no Accept-Encoding header field is present, the server SHOULD assume a
// default value of identity.
//
// This differs slightly from the HTTP definition, which indicates that when
// not present, any encoding can be used, but the identity encoding is preferred.
//
// See: https://datatracker.ietf.org/doc/html/rfc3261#section-20.2
type AcceptEncoding string

func (AcceptEncoding) Name() string { return "Accept-Encoding" }

// AcceptLanguage is used in requests to indicate the preferred languages for
// reason phrases, session descriptions, or status responses carried as message
// bodies in the response.
//
// See: https://datatracker.ietf.org/doc/html/rfc3261#section-20.3
type AcceptLanguage string

func (AcceptLanguage) Name() string { return "Accept-Language" }

// AlertInfo specifies alternative ring tones.
//
// See: https://datatracker.ietf.org/doc/html/rfc3261#section-20.4
type AlertInfo string

func (AlertInfo) Name() string { return "Alert-Info" }

// Allow lists the set of methods supported by the UA generating the message.

// All methods, including ACK and CANCEL, understood by the UA MUST be
// included in the list of methods in the Allow header field, when
// present.  The absence of an Allow header field MUST NOT be
// interpreted to mean that the UA sending the message supports no
// methods.   Rather, it implies that the UA is not providing any
// information on what methods it supports.

// Supplying an Allow header field in responses to methods other than
// OPTIONS reduces the number of messages needed.
//
// See: https://datatracker.ietf.org/doc/html/rfc3261#section-20.5
type Allow []string

func (Allow) Name() string { return "Allow" }

// AuthenticationInfo rovides for mutual authentication with HTTP Digest.
// A UAS MAY include this header field in a 2xx response to a request that
// was successfully authenticated using digest based on the Authorization header
// field.
//
// See: https://datatracker.ietf.org/doc/html/rfc3261#section-20.6
type AuthenticationInfo string

func (AuthenticationInfo) Name() string { return "Authentication-Info" }

// Authorization contains authentication credentials of a UA.
//
// See: https://datatracker.ietf.org/doc/html/rfc3261#section-20.7
type Authorization string

func (Authorization) Name() string { return "Authorization" }

// CallID uniquely identifies a particular invitation or all registrations of
// a particular client.
//
// See: https://datatracker.ietf.org/doc/html/rfc3261#section-20.8
type CallID string

func (CallID) Name() string { return "Call-ID" }

// CallInfo provides additional information about the caller or callee,
// depending on whether it is found in a request or response. The purpose
// of the URI is described by the "purpose" parameter.  The "icon" parameter
// designates an image suitable as an iconic representation of the caller or
// callee.  The "info" parameter describes the caller or callee in general,
// for example, through a web page.  The "card" parameter provides a business
// card, for example, in vCard [36] or LDIF [37] formats.  Additional tokens
// can be registered using IANA and the procedures in Section 27.
//
// See: https://datatracker.ietf.org/doc/html/rfc3261#section-20.9
type CallInfo string

func (CallInfo) Name() string { return "Call-Info" }

// Contact provides a URI whose meaning depends on the type of request or response
// it is in.
//
// A Contact header field value can contain a display name, a URI with URI
// parameters, and header parameters.
//
// This document defines the Contact parameters "q" and "expires". These
// parameters are only used when the Contact is present in a REGISTER
// request or response, or in a 3xx response.  Additional parameters may be
// defined in other specifications.
//
// When the header field value contains a display name, the URI including all
// URI parameters is enclosed in "<" and ">".  If no "<" and ">" are present, all
// parameters after the URI are header parameters, not URI parameters.  The
// display name can be tokens, or a quoted string, if a larger character set is desired.
//
// Even if the "display-name" is empty, the "name-addr" form MUST be used if the
// "addr-spec" contains a comma, semicolon, or question mark. There may or may
// not be LWS between the display-name and the "<".
//
// These rules for parsing a display name, URI and URI parameters, and header
// parameters also apply for the header fields To and From.
//
// See: https://datatracker.ietf.org/doc/html/rfc3261#section-20.10
type Contact struct {
	Scheme      string
	DisplayName string
	User        string
	Host        string
	Port        string
	Transport   string
	Q           string
	Expires     int
}

func (Contact) Name() string { return "Contact" }

// ContentDisposition describes how the message body or, for multipart
// messages, a message body part is to be interpreted by the UAC or UAS.
// This SIP header field extends the MIME Content-Type (RFC 2183 [18]).
//
// See: https://datatracker.ietf.org/doc/html/rfc3261#section-20.11
type ContentDisposition string

func (ContentDisposition) Name() string { return "Content-Disposition" }

type CSeq struct {
	Sequence uint32
	Method   string
}

func (CSeq) Name() string {
	return "CSeq"
}

type From struct {
	Scheme      string
	DisplayName string
	User        string
	Host        string
	Port        string
	Tag         string
	UserType    string
}

func (From) Name() string { return "From" }

type MaxForwards uint8

func (MaxForwards) Name() string { return "Max-Forwards" }

// Content Length indicates the size of the message-body, in decimal number of
// octets, sent to the recipient.
//
// See: https://datatracker.ietf.org/doc/html/rfc3261#section-20.14
type ContentLength int

func (ContentLength) Name() string { return "Content-Length" }

// ContentType indicates the media type of the message-body sent to the
// recipient.
//
// See: https://datatracker.ietf.org/doc/html/rfc3261#section-20.15
type ContentType string

func (ContentType) Name() string { return "Content-Type" }

// See: https://datatracker.ietf.org/doc/html/rfc3261#section-20.19
type Expires uint32

func (Expires) Name() string { return "Expires" }

type ContentEncoding string

func (ContentEncoding) Name() string { return "Content-Encoding" }

type ContentLanguage string

func (ContentLanguage) Name() string { return "Content-Language" }

type Date string

func (Date) Name() string { return "Date" }

type ErrorInfo string

func (ErrorInfo) Name() string { return "Error-Info" }

type InReplyTo string

func (InReplyTo) Name() string { return "In-Reply-To" }

type MinExpires string

func (MinExpires) Name() string { return "Min-Expires" }

type MimeVersion string

func (MimeVersion) Name() string { return "Mime-Version" }

type Organization string

func (Organization) Name() string { return "Organization" }

type Priority string

func (Priority) Name() string { return "Priority" }

type ProxyAuthenticate string

func (ProxyAuthenticate) Name() string { return "Proxy-Authenticate" }

type ProxyAuthorization string

func (ProxyAuthorization) Name() string { return "Proxy-Authorization" }

type ProxyRequire string

func (ProxyRequire) Name() string { return "Proxy-Require" }

type RecordRoute string

func (RecordRoute) Name() string { return "Record-Route" }

type ReplyTo string

func (ReplyTo) Name() string { return "Reply-To" }

type Require string

func (Require) Name() string { return "Require" }

type RetryAfter string

func (RetryAfter) Name() string { return "Retry-After" }

type Route string

func (Route) Name() string { return "Route" }

type Server string

func (Server) Name() string { return "Server" }

type Subject string

func (Subject) Name() string { return "Subject" }

type Supported string

func (Supported) Name() string { return "Supported" }

type Timestamp string

func (Timestamp) Name() string { return "Timestamp" }

type To struct {
	Scheme      string
	DisplayName string
	User        string
	Host        string
	Port        string
	Tag         string
	UserType    string
}

func (To) Name() string { return "To" }

type Unsupported string

func (Unsupported) Name() string { return "Unsupported" }

type UserAgent string

func (UserAgent) Name() string { return "User-Agent" }

type Via struct {
	Transport string
	Host      string
	Port      string
	Branch    string
	Rport     string
	Maddr     string
	TTL       string
	Received  string
}

func (Via) Name() string { return "Via" }

type Warning string

func (Warning) Name() string { return "Warning" }

type WWWAuthenticate string

func (WWWAuthenticate) Name() string { return "WWW-Authenticate" }
