package sip

import (
	"bytes"
	"strconv"
	"strings"
)

var (
	crlf           = []byte("\r\n")
	defaultParsers = map[string]func(b []byte) ([]Header, error){
		"allow":               parseAllow,
		"authentication-info": parseAuthenticationInfo,
		"authorization":       parseAuthorization,
		"to":                  parseTo,
		"t":                   parseTo,
		"from":                parseFrom,
		"f":                   parseFrom,
		"contact":             parseContact,
		"m":                   parseContact,
		"call-id":             parseCallID,
		"i":                   parseCallID,
		"cseq":                parseCSeq,
		"via":                 parseVia,
		"v":                   parseVia,
		"max-forwards":        parseMaxForwards,
		"content-length":      parseContentLength,
		"l":                   parseContentLength,
		"expires":             parseExpires,
		"user-agent":          parseUserAgent,
		"server":              parseServer,
		"content-type":        parseContentType,
		"c":                   parseContentType,
		"accept":              parseAccept,
		"require":             parseRequire,
		"supported":           parseSupported,
		"k":                   parseSupported,
		"route":               parseRoute,
		"record-route":        parseRecordRoute,
	}
)

func Parse(b []byte) (Message, error) {
	lines := bytes.Split(b, crlf)

	r, err := parseRequestLine(lines[0])
	if err != nil {
		return nil, err
	}

	msg := newMessage(r)

	for i := 1; i < len(lines); i++ {
		line := lines[i]
		spos, stype := indexSep(line)

		if spos > 0 && stype == ':' {
			hdr := strings.ToLower(string(line[0:spos]))
			val := bytes.TrimSpace(line[spos+1:])

			if parser, exists := defaultParsers[hdr]; exists {
				hdrs, err := parser(val)
				if err != nil {
					return nil, err
				}
				for _, v := range hdrs {
					msg.AppendHeader(v)
				}
			}
		}

	}
	return msg, nil
}

type Field int

const (
	FieldNil        Field = 0
	FieldBase       Field = 1
	FieldValue      Field = 2
	FieldName       Field = 3
	FieldNameQ      Field = 4
	FieldUser       Field = 5
	FieldHost       Field = 6
	FieldPort       Field = 7
	FieldTag        Field = 8
	FieldID         Field = 9
	FieldMethod     Field = 10
	FieldTransport  Field = 11
	FieldBranch     Field = 12
	FieldRport      Field = 13
	FieldMaddr      Field = 14
	FieldTTL        Field = 15
	FieldReceived   Field = 16
	FieldExpires    Field = 17
	FieldQ          Field = 18
	FieldUserType   Field = 19
	FieldStatus     Field = 20
	FieldStatusDesc Field = 21
	FieldAddrType   Field = 40
	FieldConnAddr   Field = 41
	FieldMedia      Field = 42
	FieldProto      Field = 43
	FieldFmt        Field = 44
	FieldCat        Field = 45
	FieldIgnore     Field = 255
)

func parseRequestLine(b []byte) (*RequestLine, error) {
	var (
		pos               = 0
		state             = FieldNil
		method            = []byte{}
		scheme            string
		statusCode        = []byte{}
		statusDescription = []byte{}
		user              = []byte{}
		host              = []byte{}
		port              = []byte{}
		userType          = []byte{}
	)

	// Loop through the bytes making up the line
	for pos < len(b) {
		// FSM
		switch state {
		case FieldNil:
			if b[pos] >= 'A' && b[pos] <= 'S' && pos == 0 {
				state = FieldMethod
				continue
			}

		case FieldMethod:
			if b[pos] == ' ' || pos > 9 {
				if string(method) == "SIP/2.0" {
					state = FieldStatus
					method = []byte{}
				} else {
					state = FieldBase
				}
				pos++
				continue
			}
			method = append(method, b[pos])

		case FieldBase:
			if b[pos] != ' ' {
				// Not a space so check for uri types
				if getString(b, pos, pos+4) == "sip:" {
					state = FieldUser
					pos = pos + 4
					scheme = "sip"
					continue
				}
				if getString(b, pos, pos+5) == "sips:" {
					state = FieldUser
					pos = pos + 5
					scheme = "sips"
					continue
				}
				if getString(b, pos, pos+4) == "tel:" {
					state = FieldUser
					pos = pos + 4
					scheme = "tel"
					continue
				}
				if getString(b, pos, pos+5) == "user=" {
					state = FieldUserType
					pos = pos + 5
					continue
				}
				if b[pos] == '@' {
					state = FieldHost
					user = host // Move host to user
					host = nil  // Clear the host
					pos++
					continue
				}
			}
		case FieldUser:
			if b[pos] == ':' {
				state = FieldPort
				pos++
				continue
			}
			if b[pos] == ';' || b[pos] == '>' {
				state = FieldBase
				pos++
				continue
			}
			if b[pos] == '@' {
				state = FieldHost
				user = host // Move host to user
				host = nil  // Clear the host
				pos++
				continue
			}
			host = append(host, b[pos]) // Append to host for now

		case FieldHost:
			if b[pos] == ':' {
				state = FieldPort
				pos++
				continue
			}
			if b[pos] == ';' || b[pos] == '>' || b[pos] == ' ' {
				state = FieldBase
				pos++
				continue
			}
			host = append(host, b[pos])

		case FieldPort:
			if b[pos] == ';' || b[pos] == '>' || b[pos] == ' ' {
				state = FieldBase
				pos++
				continue
			}
			port = append(port, b[pos])

		case FieldUserType:
			if b[pos] == ';' || b[pos] == '>' || b[pos] == ' ' {
				state = FieldBase
				pos++
				continue
			}
			userType = append(userType, b[pos])

		case FieldStatus:
			if b[pos] == ';' || b[pos] == '>' {
				state = FieldBase
				pos++
				continue
			}
			if b[pos] == ' ' {
				state = FieldStatusDesc
				pos++
				continue
			}
			statusCode = append(statusCode, b[pos])

		case FieldStatusDesc:
			if b[pos] == ';' || b[pos] == '>' {
				state = FieldBase
				pos++
				continue
			}
			statusDescription = append(statusDescription, b[pos])

		}
		pos++
	}

	var result RequestLine
	result.Method = string(method)
	result.Scheme = scheme

	if len(statusCode) > 0 {
		code, err := strconv.Atoi(string(statusCode))
		if err != nil {
			return nil, err
		}
		result.StatusCode = code
	}

	result.StatusDescription = string(statusDescription)
	result.User = string(user)
	result.Host = string(host)
	result.Port = string(port)
	result.UserType = string(userType)
	return &result, nil
}

func parseAllow(b []byte) ([]Header, error) {
	return []Header{Allow(strings.Split(string(b), ","))}, nil
}

func parseAuthenticationInfo(b []byte) ([]Header, error) {
	return []Header{AuthenticationInfo(string(b))}, nil
}

func parseAuthorization(b []byte) ([]Header, error) {
	return []Header{Authorization(string(b))}, nil
}

func parseCSeq(b []byte) ([]Header, error) {
	var (
		pos      = 0
		state    = FieldID
		sequence = []byte{}
		method   = []byte{}
	)

	for pos < len(b) {
		switch state {
		case FieldID:
			if b[pos] == ' ' {
				state = FieldMethod
				pos++
				continue
			}
			sequence = append(sequence, b[pos])
		case FieldMethod:
			method = append(method, b[pos])
		}
		pos++
	}

	var result CSeq
	val, err := strconv.Atoi(string(sequence))
	if err != nil {
		return nil, err
	}

	result.Sequence = uint32(val)
	result.Method = strings.TrimSpace(string(method))
	return []Header{&result}, nil
}

func parseFrom(b []byte) ([]Header, error) {
	var (
		pos      = 0
		state    = FieldBase
		scheme   string
		name     = []byte{}
		user     = []byte{}
		host     = []byte{}
		port     = []byte{}
		userType = []byte{}
		tag      = []byte{}
	)

	for pos < len(b) {
		switch state {
		case FieldBase:
			if b[pos] == '"' && scheme == "" {
				state = FieldNameQ
				pos++
				continue
			}
			if b[pos] != ' ' {
				// Not a space so check for uri types
				if getString(b, pos, pos+4) == "sip:" {
					state = FieldUser
					pos = pos + 4
					scheme = "sip"
					continue
				}
				if getString(b, pos, pos+5) == "sips:" {
					state = FieldUser
					pos = pos + 5
					scheme = "sips"
					continue
				}
				if getString(b, pos, pos+4) == "tel:" {
					state = FieldUser
					pos = pos + 4
					scheme = "tel"
					continue
				}
				// Look for a Tag identifier
				if getString(b, pos, pos+4) == "tag=" {
					state = FieldTag
					pos = pos + 4
					continue
				}
				// Look for other identifiers and ignore
				if b[pos] == '=' {
					state = FieldIgnore
					pos = pos + 1
					continue
				}
				// Look for a User Type identifier
				if getString(b, pos, pos+5) == "user=" {
					state = FieldUserType
					pos = pos + 5
					continue
				}
				// Check for other chrs
				if b[pos] != '<' && b[pos] != '>' && b[pos] != ';' && scheme == "" {
					state = FieldName
					continue
				}
			}

		case FieldNameQ:
			if b[pos] == '"' {
				state = FieldBase
				pos++
				continue
			}
			name = append(name, b[pos])

		case FieldName:
			if b[pos] == '<' || b[pos] == ' ' {
				state = FieldBase
				pos++
				continue
			}
			name = append(name, b[pos])

		case FieldUser:
			if b[pos] == '@' {
				state = FieldHost
				pos++
				continue
			}
			user = append(user, b[pos])

		case FieldHost:
			if b[pos] == ':' {
				state = FieldPort
				pos++
				continue
			}
			if b[pos] == ';' || b[pos] == '>' {
				state = FieldBase
				pos++
				continue
			}
			host = append(host, b[pos])

		case FieldPort:
			if b[pos] == ';' || b[pos] == '>' || b[pos] == ' ' {
				state = FieldBase
				pos++
				continue
			}
			port = append(port, b[pos])

		case FieldUserType:
			if b[pos] == ';' || b[pos] == '>' || b[pos] == ' ' {
				state = FieldBase
				pos++
				continue
			}
			userType = append(userType, b[pos])

		case FieldTag:
			if b[pos] == ';' || b[pos] == '>' || b[pos] == ' ' {
				state = FieldBase
				pos++
				continue
			}
			tag = append(tag, b[pos])

		case FieldIgnore:
			if b[pos] == ';' || b[pos] == '>' {
				state = FieldBase
				pos++
				continue
			}
		}
		pos++
	}

	var result From
	result.Scheme = scheme
	result.DisplayName = string(name)
	result.User = string(user)
	result.Host = string(host)
	result.Port = string(port)
	result.Tag = string(tag)
	result.UserType = string(userType)

	return []Header{&result}, nil
}

func parseTo(b []byte) ([]Header, error) {
	var (
		pos         = 0
		state       = FieldBase
		scheme      string
		displayName = []byte{}
		user        = []byte{}
		host        = []byte{}
		port        = []byte{}
		userType    = []byte{}
		tag         = []byte{}
	)

	for pos < len(b) {
		switch state {
		case FieldBase:
			if b[pos] == '"' && scheme == "" {
				state = FieldNameQ
				pos++
				continue
			}
			if b[pos] != ' ' {
				// Not a space so check for uri types
				if getString(b, pos, pos+4) == "sip:" {
					state = FieldUser
					pos = pos + 4
					scheme = "sip"
					continue
				}
				if getString(b, pos, pos+5) == "sips:" {
					state = FieldUser
					pos = pos + 5
					scheme = "sips"
					continue
				}
				if getString(b, pos, pos+4) == "tel:" {
					state = FieldUser
					pos = pos + 4
					scheme = "tel"
					continue
				}
				// Look for a Tag identifier
				if getString(b, pos, pos+4) == "tag=" {
					state = FieldTag
					pos = pos + 4
					continue
				}
				// Look for other identifiers and ignore
				if b[pos] == '=' {
					state = FieldIgnore
					pos = pos + 1
					continue
				}
				// Look for a User Type identifier
				if getString(b, pos, pos+5) == "user=" {
					state = FieldUserType
					pos = pos + 5
					continue
				}
				// Check for other chrs
				if b[pos] != '<' && b[pos] != '>' && b[pos] != ';' && scheme == "" {
					state = FieldName
					continue
				}
			}

		case FieldNameQ:
			if b[pos] == '"' {
				state = FieldBase
				pos++
				continue
			}
			displayName = append(displayName, b[pos])

		case FieldName:
			if b[pos] == '<' || b[pos] == ' ' {
				state = FieldBase
				pos++
				continue
			}
			displayName = append(displayName, b[pos])

		case FieldUser:
			if b[pos] == '@' {
				state = FieldHost
				pos++
				continue
			}
			user = append(user, b[pos])

		case FieldHost:
			if b[pos] == ':' {
				state = FieldPort
				pos++
				continue
			}
			if b[pos] == ';' || b[pos] == '>' {
				state = FieldBase
				pos++
				continue
			}
			host = append(host, b[pos])

		case FieldPort:
			if b[pos] == ';' || b[pos] == '>' || b[pos] == ' ' {
				state = FieldBase
				pos++
				continue
			}
			port = append(port, b[pos])

		case FieldUserType:
			if b[pos] == ';' || b[pos] == '>' || b[pos] == ' ' {
				state = FieldBase
				pos++
				continue
			}
			userType = append(userType, b[pos])

		case FieldTag:
			if b[pos] == ';' || b[pos] == '>' || b[pos] == ' ' {
				state = FieldBase
				pos++
				continue
			}
			tag = append(tag, b[pos])

		case FieldIgnore:
			if b[pos] == ';' || b[pos] == '>' {
				state = FieldBase
				pos++
				continue
			}
		}
		pos++
	}

	var result To
	result.Scheme = scheme
	result.DisplayName = string(displayName)
	result.User = string(user)
	result.Host = string(host)
	result.Port = string(port)
	result.Tag = string(tag)
	result.UserType = string(userType)

	return []Header{&result}, nil
}

func parseContact(b []byte) ([]Header, error) {
	var (
		pos         = 0
		state       = FieldBase
		scheme      string
		displayName = []byte{}
		user        = []byte{}
		host        = []byte{}
		port        = []byte{}
		transport   = []byte{}
		q           = []byte{}
		expires     = []byte{}
	)

	for pos < len(b) {
		switch state {
		case FieldBase:
			if b[pos] == '"' && scheme == "" {
				state = FieldNameQ
				pos++
				continue
			}
			if b[pos] != ' ' {
				// Not a space so check for uri types
				if getString(b, pos, pos+4) == "sip:" {
					state = FieldUser
					pos = pos + 4
					scheme = "sip"
					continue
				}
				if getString(b, pos, pos+5) == "sips:" {
					state = FieldUser
					pos = pos + 5
					scheme = "sips"
					continue
				}
				if getString(b, pos, pos+4) == "tel:" {
					state = FieldUser
					pos = pos + 4
					scheme = "tel"
					continue
				}
				// Look for a Q identifier
				if getString(b, pos, pos+2) == "q=" {
					state = FieldQ
					pos = pos + 2
					continue
				}
				// Look for a Expires identifier
				if getString(b, pos, pos+8) == "expires=" {
					state = FieldExpires
					pos = pos + 8
					continue
				}
				// Look for a transport identifier
				if getString(b, pos, pos+10) == "transport=" {
					state = FieldTransport
					pos = pos + 10
					continue
				}
				if getString(b, pos, pos+4) == "tsp=" {
					state = FieldTransport
					pos = pos + 4
					continue
				}
				// Look for other identifiers and ignore
				if b[pos] == '=' {
					state = FieldIgnore
					pos = pos + 1
					continue
				}
				// Check for other chrs
				if b[pos] != '<' && b[pos] != '>' && b[pos] != ';' && scheme == "" {
					state = FieldName
					continue
				}
			}

		case FieldNameQ:
			if b[pos] == '"' {
				state = FieldBase
				pos++
				continue
			}
			displayName = append(displayName, b[pos])

		case FieldName:
			if b[pos] == '<' || b[pos] == ' ' {
				state = FieldBase
				pos++
				continue
			}
			displayName = append(displayName, b[pos])

		case FieldUser:
			if b[pos] == '@' {
				state = FieldHost
				pos++
				continue
			}
			user = append(user, b[pos])

		case FieldHost:
			if b[pos] == ':' {
				state = FieldPort
				pos++
				continue
			}
			if b[pos] == ';' || b[pos] == '>' {
				state = FieldBase
				pos++
				continue
			}
			host = append(host, b[pos])

		case FieldPort:
			if b[pos] == ';' || b[pos] == '>' || b[pos] == ' ' {
				state = FieldBase
				pos++
				continue
			}
			port = append(port, b[pos])

		case FieldTransport:
			if b[pos] == ';' || b[pos] == '>' || b[pos] == ' ' {
				state = FieldBase
				pos++
				continue
			}
			transport = append(transport, b[pos])

		case FieldQ:
			if b[pos] == ';' || b[pos] == '>' || b[pos] == ' ' {
				state = FieldBase
				pos++
				continue
			}
			q = append(q, b[pos])

		case FieldExpires:
			if b[pos] == ';' || b[pos] == '>' || b[pos] == ' ' {
				state = FieldBase
				pos++
				continue
			}
			expires = append(expires, b[pos])

		case FieldIgnore:
			if b[pos] == ';' || b[pos] == '>' {
				state = FieldBase
				pos++
				continue
			}
		}
		pos++
	}

	var result Contact
	result.Scheme = string(scheme)
	result.DisplayName = string(displayName)
	result.User = string(user)
	result.Host = string(host)
	result.Port = string(port)
	result.Transport = string(transport)
	result.Q = string(q)

	if len(expires) > 0 {
		exp, err := strconv.Atoi(string(expires))
		if err != nil {
			return nil, err
		}
		result.Expires = exp
	}

	return []Header{&result}, nil
}

func parseVia(b []byte) ([]Header, error) {
	var (
		pos       = 0
		state     = FieldBase
		transport string
		host      = []byte{}
		port      = []byte{}
		branch    = []byte{}
		rport     = []byte{}
		maddr     = []byte{}
		ttl       = []byte{}
		received  = []byte{}
	)

	for pos < len(b) {
		switch state {
		case FieldBase:
			if b[pos] != ' ' {
				// Not a space
				if getString(b, pos, pos+8) == "SIP/2.0/" {
					// Transport type
					state = FieldHost
					pos = pos + 8
					if getString(b, pos, pos+3) == "UDP" {
						transport = "udp"
						pos = pos + 3
						continue
					}
					if getString(b, pos, pos+3) == "TCP" {
						transport = "tcp"
						pos = pos + 3
						continue
					}
					if getString(b, pos, pos+3) == "TLS" {
						transport = "tls"
						pos = pos + 3
						continue
					}
					if getString(b, pos, pos+4) == "SCTP" {
						transport = "sctp"
						pos = pos + 4
						continue
					}
				}
				// Look for a Branch identifier
				if getString(b, pos, pos+7) == "branch=" {
					state = FieldBranch
					pos = pos + 7
					continue
				}
				// Look for a Rport identifier
				if getString(b, pos, pos+6) == "rport=" {
					state = FieldRport
					pos = pos + 6
					continue
				}
				// Look for a maddr identifier
				if getString(b, pos, pos+6) == "maddr=" {
					state = FieldMaddr
					pos = pos + 6
					continue
				}
				// Look for a ttl identifier
				if getString(b, pos, pos+4) == "ttl=" {
					state = FieldTTL
					pos = pos + 4
					continue
				}
				// Look for a recevived identifier
				if getString(b, pos, pos+9) == "received=" {
					state = FieldReceived
					pos = pos + 9
					continue
				}
			}
		case FieldHost:
			if b[pos] == ':' {
				state = FieldPort
				pos++
				continue
			}
			if b[pos] == ';' {
				state = FieldBase
				pos++
				continue
			}
			if b[pos] == ' ' {
				pos++
				continue
			}
			host = append(host, b[pos])

		case FieldPort:
			if b[pos] == ';' {
				state = FieldBase
				pos++
				continue
			}
			port = append(port, b[pos])

		case FieldBranch:
			if b[pos] == ';' {
				state = FieldBase
				pos++
				continue
			}
			branch = append(branch, b[pos])

		case FieldRport:
			if b[pos] == ';' {
				state = FieldBase
				pos++
				continue
			}
			rport = append(rport, b[pos])

		case FieldMaddr:
			if b[pos] == ';' {
				state = FieldBase
				pos++
				continue
			}
			maddr = append(maddr, b[pos])

		case FieldTTL:
			if b[pos] == ';' {
				state = FieldBase
				pos++
				continue
			}
			ttl = append(ttl, b[pos])

		case FieldReceived:
			if b[pos] == ';' {
				state = FieldBase
				pos++
				continue
			}
			received = append(received, b[pos])
		}
		pos++
	}

	var result Via
	result.Transport = transport
	result.Host = string(host)
	result.Port = string(port)
	result.Branch = string(branch)
	result.Rport = string(rport)
	result.Maddr = string(maddr)
	result.TTL = string(ttl)
	result.Received = string(received)
	return []Header{&result}, nil
}

func parseMaxForwards(b []byte) ([]Header, error) {
	val, err := strconv.Atoi(string(b))
	if err != nil {
		return nil, err
	}

	return []Header{MaxForwards(val)}, nil
}

func parseContentType(b []byte) ([]Header, error) {
	return []Header{ContentType(b)}, nil
}

func parseCallID(b []byte) ([]Header, error) {
	return []Header{CallID(b)}, nil
}

func parseContentLength(b []byte) ([]Header, error) {
	v, err := strconv.Atoi(string(b))
	if err != nil {
		return nil, err
	}
	return []Header{ContentLength(v)}, nil
}

func parseExpires(b []byte) ([]Header, error) {
	val, err := strconv.Atoi(string(b))
	if err != nil {
		return nil, err
	}
	return []Header{Expires(val)}, nil
}

func parseRecordRoute(b []byte) ([]Header, error) {
	return []Header{RecordRoute(string(b))}, nil
}

func parseUserAgent(b []byte) ([]Header, error) {
	return []Header{UserAgent(string(b))}, nil
}

func parseRoute(b []byte) ([]Header, error) {
	return []Header{Route(b)}, nil
}

func parseSupported(b []byte) ([]Header, error) {
	return []Header{Supported(string(b))}, nil
}

func parseRequire(b []byte) ([]Header, error) {
	return []Header{Require(string(b))}, nil
}

func parseServer(b []byte) ([]Header, error) {
	return []Header{Server(string(b))}, nil
}

func parseAccept(b []byte) ([]Header, error) {
	return []Header{Accept(string(b))}, nil
}

// Get a string from a slice of bytes
// Checks the bounds to avoid any range errors
func getString(sl []byte, from, to int) string {
	// Remove negative values
	if from < 0 {
		from = 0
	}
	if to < 0 {
		to = 0
	}
	// Limit if over len
	if from > len(sl) || from > to {
		return ""
	}
	if to > len(sl) {
		return string(sl[from:])
	}
	return string(sl[from:to])
}

// Finds the first valid Seperate or notes its type
func indexSep(s []byte) (int, byte) {
	for i := 0; i < len(s); i++ {
		if s[i] == ':' {
			return i, ':'
		}
		if s[i] == '=' {
			return i, '='
		}
	}
	return -1, ' '
}
