package sip

const (
	StatusTrying                = 100
	StatusRinging               = 180
	StatusCallIsBeingForwarded  = 181
	StatusQueued                = 182
	StatusSessionProgress       = 183
	StatusEarlyDialogTerminated = 199

	StatusOK             = 200
	StatusAccepted       = 202
	StatusNoNotification = 204

	StatusMultipleChoices    = 300
	StatusMovedPermanently   = 301
	StatusMovedTemporarily   = 302
	StatusUseProxy           = 305
	StatusAlternativeService = 380

	StatusBadRequest                   = 400
	StatusUnauthorized                 = 401
	StatusPaymentRequired              = 402
	StatusForbidden                    = 403
	StatusNotFound                     = 404
	StatusMethodNotAllowed             = 405
	StatusNotAcceptable                = 406
	StatusProxyAuthenticationRequired  = 407
	StatusRequestTimeout               = 408
	StatusGone                         = 410
	StatusConditionalRequestFailed     = 412
	StatusRequestEntityTooLarge        = 413
	StatusRequestURITooLong            = 414
	StatusUnsupportedMediaType         = 415
	StatusUnsupportedURIScheme         = 416
	StatusUnknownResourcePriority      = 417
	StatusBadExtension                 = 420
	StatusExtensionRequired            = 421
	StatusSessionIntervalTooSmall      = 422
	StatusIntervalTooBrief             = 423
	StatusBadLocationInformation       = 424
	StatusBadAlertMessage              = 425
	StatusUseIdentityHeader            = 428
	StatusProvideReferrerIdentity      = 429
	StatusFlowFailed                   = 430
	StatusAnonymityDisallowed          = 433
	StatusBadIdentityInfo              = 436
	StatusUnsupportedCredential        = 437
	StatusInvalidIdentityHeader        = 438
	StatusFirstHopLacksOutboundSupport = 439
	StatusMaxBreadthExceeded           = 440
	StatusBadInfoPackage               = 469
	StatusConsentNeeded                = 470
	StatusTemporarilyUnavailable       = 480
	StatusCallTransactionDoesNotExist  = 481
	StatusLoopDetected                 = 482
	StatusTooManyHops                  = 483
	StatusAddressIncomplete            = 484
	StatusAmbiguous                    = 485
	StatusBusyHere                     = 486
	StatusRequestTerminated            = 487
	StatusNotAcceptableHere            = 488
	StatusBadEvent                     = 489
	StatusRequestPending               = 491
	StatusUndecipherable               = 493
	StatusSecurityAgreementRequired    = 494

	StatusInternalServerError                 = 500
	StatusNotImplemented                      = 501
	StatusBadGateway                          = 502
	StatusServiceUnavailable                  = 503
	StatusServerTimeout                       = 504
	StatusVersionNotSupported                 = 505
	StatusMessageTooLarge                     = 513
	StatusPushNotificationServiceNotSupported = 555
	StatusPreconditionFailure                 = 580

	StatusBusyEverywhere       = 600
	StatusDecline              = 603
	StatusDoesNotExistAnywhere = 604
)

var defaultResponses = map[int]string{
	100: "Trying",
	180: "Ringing",
	181: "Call Is Being Forwarded",
	182: "Queued",
	183: "Session Progress",
	199: "Early Dialog Terminated",
	200: "OK",
	202: "Accepted",
	204: "No Notification",
	300: "Multiple Choices",
	301: "Moved Permanently",
	302: "Moved Temporarily",
	305: "Use Proxy",
	380: "Alternative Service",
	400: "Bad Request",
	401: "Unauthorized",
	402: "Payment Required",
	403: "Forbidden",
	404: "Not Found",
	405: "Method Not Allowed",
	406: "Not Acceptable",
	407: "Proxy Authentication Required",
	408: "Request Timeout",
	410: "Gone",
	412: "Conditional Request Failed",
	413: "Request Entity Too Large",
	414: "Request URI Too Long",
	415: "Unsupported Media Type",
	416: "Unsupported URI Type",
	417: "Unknown Resource Priority",
	420: "Bad Extension",
	421: "Extension Required",
	422: "Session Interval Too Small",
	423: "Interval Too Brief",
	424: "Bad Location Information",
	425: "Bad Alert Message",
	428: "Use Identity Header",
	429: "Provide Referrer Identity",
	430: "Flow Failed",
	433: "Anonymity Disallowed",
	436: "Bad Identity Info",
	437: "Unsupported Credential",
	438: "Invalid Identity Header",
	439: "First Hop Lacks Outbound Support",
	440: "Max Breadth Exceeded",
	469: "Bad Info Package",
	470: "Consent Needed",
	480: "Temporarily Unavailable",
	481: "Call/Transaction Does Not Exist",
	482: "Loop Detected",
	483: "Too Many Hops",
	484: "Address Incomplete",
	485: "Ambiguous",
	486: "Busy Here",
	487: "Request Terminated",
	488: "Not Acceptable Here",
	489: "Bad Event",
	491: "Request Pending",
	493: "Undecipherabe",
	494: "Security Agreement Required",
	500: "Internal Server Error",
	501: "Not Implemented",
	502: "Bad Gateway",
	503: "Service Unavailable",
	504: "Server Time-out",
	505: "Version Not Supported",
	513: "Message Too Large",
	555: "Push Notification Service Not Supported",
	580: "Precondition Failure",
	600: "Busy Everywhere",
	603: "Decline",
	604: "Does Not Exist Anywhere",
}

func StatusText(code int) string {
	if str, exists := defaultResponses[code]; exists {
		return str
	}
	return ""
}
