package httperror

import "net/http"

// HTTP 100 Continue
func NewContinue(message ...string) *HTTPError {
	return New(http.StatusContinue, message...)
}

// HTTP 100 Continue
func WrapContinue(err error) *HTTPError {
	return Wrap(http.StatusContinue, err)
}

// HTTP 101 Switching Protocols
func NewSwitchingProtocols(message ...string) *HTTPError {
	return New(http.StatusSwitchingProtocols, message...)
}

// HTTP 101 Switching Protocols
func WrapSwitchingProtocols(err error) *HTTPError {
	return Wrap(http.StatusSwitchingProtocols, err)
}

// HTTP 102 Processing
func NewProcessing(message ...string) *HTTPError {
	return New(http.StatusProcessing, message...)
}

// HTTP 102 Processing
func WrapProcessing(err error) *HTTPError {
	return Wrap(http.StatusProcessing, err)
}

// HTTP 200 OK
func NewOK(message ...string) *HTTPError {
	return New(http.StatusOK, message...)
}

// HTTP 200 OK
func WrapOK(err error) *HTTPError {
	return Wrap(http.StatusOK, err)
}

// HTTP 201 Created
func NewCreated(message ...string) *HTTPError {
	return New(http.StatusCreated, message...)
}

// HTTP 201 Created
func WrapCreated(err error) *HTTPError {
	return Wrap(http.StatusCreated, err)
}

// HTTP 202 Accepted
func NewAccepted(message ...string) *HTTPError {
	return New(http.StatusAccepted, message...)
}

// HTTP 202 Accepted
func WrapAccepted(err error) *HTTPError {
	return Wrap(http.StatusAccepted, err)
}

// HTTP 203 NonAuthoritativeInfo
func NewNonAuthoritativeInfo(message ...string) *HTTPError {
	return New(http.StatusNonAuthoritativeInfo, message...)
}

// HTTP 203 NonAuthoritativeInfo
func WrapNonAuthoritativeInfo(err error) *HTTPError {
	return Wrap(http.StatusNonAuthoritativeInfo, err)
}

// HTTP 204 No Content
func NewNoContent(message ...string) *HTTPError {
	return New(http.StatusNoContent, message...)
}

// HTTP 204 No Content
func WrapNoContent(err error) *HTTPError {
	return Wrap(http.StatusNoContent, err)
}

// HTTP 205 Reset Content
func NewResetContent(message ...string) *HTTPError {
	return New(http.StatusResetContent, message...)
}

// HTTP 205 Reset Content
func WrapResetContent(err error) *HTTPError {
	return Wrap(http.StatusResetContent, err)
}

// HTTP 206 Partial Content
func NewPartialContent(message ...string) *HTTPError {
	return New(http.StatusPartialContent, message...)
}

// HTTP 206 Partial Content
func WrapPartialContent(err error) *HTTPError {
	return Wrap(http.StatusPartialContent, err)
}

// HTTP 207 Multi-Status
func NewMultiStatus(message ...string) *HTTPError {
	return New(http.StatusMultiStatus, message...)
}

// HTTP 207 Multi-Status
func WrapMultiStatus(err error) *HTTPError {
	return Wrap(http.StatusMultiStatus, err)
}

// HTTP 208 Already Reported
func NewAlreadyReported(message ...string) *HTTPError {
	return New(http.StatusAlreadyReported, message...)
}

// HTTP 208 Already Reported
func WrapAlreadyReported(err error) *HTTPError {
	return Wrap(http.StatusAlreadyReported, err)
}

// HTTP 226 IM Used
func NewIMUsed(message ...string) *HTTPError {
	return New(http.StatusIMUsed, message...)
}

// HTTP 226 IM Used
func WrapIMUsed(err error) *HTTPError {
	return Wrap(http.StatusIMUsed, err)
}

// HTTP 300 Multiple Choices
func NewMultipleChoices(message ...string) *HTTPError {
	return New(http.StatusMultipleChoices, message...)
}

// HTTP 300 Multiple Choices
func WrapMultipleChoices(err error) *HTTPError {
	return Wrap(http.StatusMultipleChoices, err)
}

// HTTP 301 Moved Permanently
func NewMovedPermanently(message ...string) *HTTPError {
	return New(http.StatusMovedPermanently, message...)
}

// HTTP 301 Moved Permanently
func WrapMovedPermanently(err error) *HTTPError {
	return Wrap(http.StatusMovedPermanently, err)
}

// HTTP 302 Found
func NewFound(message ...string) *HTTPError {
	return New(http.StatusFound, message...)
}

// HTTP 302 Found
func WrapFound(err error) *HTTPError {
	return Wrap(http.StatusFound, err)
}

// HTTP 303 See Other
func NewSeeOther(message ...string) *HTTPError {
	return New(http.StatusSeeOther, message...)
}

// HTTP 303 See Other
func WrapSeeOther(err error) *HTTPError {
	return Wrap(http.StatusSeeOther, err)
}

// HTTP 304 Not Modified
func NewNotModified(message ...string) *HTTPError {
	return New(http.StatusNotModified, message...)
}

// HTTP 304 Not Modified
func WrapNotModified(err error) *HTTPError {
	return Wrap(http.StatusNotModified, err)
}

// HTTP 305 Use Proxy
func NewUseProxy(message ...string) *HTTPError {
	return New(http.StatusUseProxy, message...)
}

// HTTP 305 Use Proxy
func WrapUseProxy(err error) *HTTPError {
	return Wrap(http.StatusUseProxy, err)
}

// HTTP 307 Temporary Redirect
func NewTemporaryRedirect(message ...string) *HTTPError {
	return New(http.StatusTemporaryRedirect, message...)
}

// HTTP 307 Temporary Redirect
func WrapTemporaryRedirect(err error) *HTTPError {
	return Wrap(http.StatusTemporaryRedirect, err)
}

// HTTP 308 Permanent Redirect
func NewPermanentRedirect(message ...string) *HTTPError {
	return New(http.StatusPermanentRedirect, message...)
}

// HTTP 308 Permanent Redirect
func WrapPermanentRedirect(err error) *HTTPError {
	return Wrap(http.StatusPermanentRedirect, err)
}

// HTTP 400 Bad Request
func NewBadRequest(message ...string) *HTTPError {
	return New(http.StatusBadRequest, message...)
}

// HTTP 400 Bad Request
func WrapBadRequest(err error) *HTTPError {
	return Wrap(http.StatusBadRequest, err)
}

// HTTP 401 Unauthorized
func NewUnauthorized(message ...string) *HTTPError {
	return New(http.StatusUnauthorized, message...)
}

// HTTP 401 Unauthorized
func WrapUnauthorized(err error) *HTTPError {
	return Wrap(http.StatusUnauthorized, err)
}

// HTTP 402 Payment Required
func NewPaymentRequired(message ...string) *HTTPError {
	return New(http.StatusPaymentRequired, message...)
}

// HTTP 402 Payment Required
func WrapPaymentRequired(err error) *HTTPError {
	return Wrap(http.StatusPaymentRequired, err)
}

// HTTP 403 Forbidden
func NewForbidden(message ...string) *HTTPError {
	return New(http.StatusForbidden, message...)
}

// HTTP 403 Forbidden
func WrapForbidden(err error) *HTTPError {
	return Wrap(http.StatusForbidden, err)
}

// HTTP 404 Not Found
func NewNotFound(message ...string) *HTTPError {
	return New(http.StatusNotFound, message...)
}

// HTTP 404 Not Found
func WrapNotFound(err error) *HTTPError {
	return Wrap(http.StatusNotFound, err)
}

// HTTP 405 Method Not Allowed
func NewMethodNotAllowed(message ...string) *HTTPError {
	return New(http.StatusMethodNotAllowed, message...)
}

// HTTP 405 Method Not Allowed
func WrapMethodNotAllowed(err error) *HTTPError {
	return Wrap(http.StatusMethodNotAllowed, err)
}

// HTTP 406 Not Acceptable
func NewNotAcceptable(message ...string) *HTTPError {
	return New(http.StatusNotAcceptable, message...)
}

// HTTP 406 Not Acceptable
func WrapNotAcceptable(err error) *HTTPError {
	return Wrap(http.StatusNotAcceptable, err)
}

// HTTP 407 ProxyAuthRequired
func NewProxyAuthRequired(message ...string) *HTTPError {
	return New(http.StatusProxyAuthRequired, message...)
}

// HTTP 407 ProxyAuthRequired
func WrapProxyAuthRequired(err error) *HTTPError {
	return Wrap(http.StatusProxyAuthRequired, err)
}

// HTTP 408 Request Timeout
func NewRequestTimeout(message ...string) *HTTPError {
	return New(http.StatusRequestTimeout, message...)
}

// HTTP 408 Request Timeout
func WrapRequestTimeout(err error) *HTTPError {
	return Wrap(http.StatusRequestTimeout, err)
}

// HTTP 409 Conflict
func NewConflict(message ...string) *HTTPError {
	return New(http.StatusConflict, message...)
}

// HTTP 409 Conflict
func WrapConflict(err error) *HTTPError {
	return Wrap(http.StatusConflict, err)
}

// HTTP 410 Gone
func NewGone(message ...string) *HTTPError {
	return New(http.StatusGone, message...)
}

// HTTP 410 Gone
func WrapGone(err error) *HTTPError {
	return Wrap(http.StatusGone, err)
}

// HTTP 411 Length Required
func NewLengthRequired(message ...string) *HTTPError {
	return New(http.StatusLengthRequired, message...)
}

// HTTP 411 Length Required
func WrapLengthRequired(err error) *HTTPError {
	return Wrap(http.StatusLengthRequired, err)
}

// HTTP 412 Precondition Failed
func NewPreconditionFailed(message ...string) *HTTPError {
	return New(http.StatusPreconditionFailed, message...)
}

// HTTP 412 Precondition Failed
func WrapPreconditionFailed(err error) *HTTPError {
	return Wrap(http.StatusPreconditionFailed, err)
}

// HTTP 413 Request Entity Too Large
func NewRequestEntityTooLarge(message ...string) *HTTPError {
	return New(http.StatusRequestEntityTooLarge, message...)
}

// HTTP 413 Request Entity Too Large
func WrapRequestEntityTooLarge(err error) *HTTPError {
	return Wrap(http.StatusRequestEntityTooLarge, err)
}

// HTTP 414 Request URI Too Long
func NewRequestURITooLong(message ...string) *HTTPError {
	return New(http.StatusRequestURITooLong, message...)
}

// HTTP 414 Request URI Too Long
func WrapRequestURITooLong(err error) *HTTPError {
	return Wrap(http.StatusRequestURITooLong, err)
}

// HTTP 415 Unsupported Media Type
func NewUnsupportedMediaType(message ...string) *HTTPError {
	return New(http.StatusUnsupportedMediaType, message...)
}

// HTTP 415 Unsupported Media Type
func WrapUnsupportedMediaType(err error) *HTTPError {
	return Wrap(http.StatusUnsupportedMediaType, err)
}

// HTTP 416 Requested Range Not Satisfiable
func NewRequestedRangeNotSatisfiable(message ...string) *HTTPError {
	return New(http.StatusRequestedRangeNotSatisfiable, message...)
}

// HTTP 416 Requested Range Not Satisfiable
func WrapRequestedRangeNotSatisfiable(err error) *HTTPError {
	return Wrap(http.StatusRequestedRangeNotSatisfiable, err)
}

// HTTP 417 Expectation Failed
func NewExpectationFailed(message ...string) *HTTPError {
	return New(http.StatusExpectationFailed, message...)
}

// HTTP 417 Expectation Failed
func WrapExpectationFailed(err error) *HTTPError {
	return Wrap(http.StatusExpectationFailed, err)
}

// HTTP 418 Teapot
func NewTeapot(message ...string) *HTTPError {
	return New(http.StatusTeapot, message...)
}

// HTTP 418 Teapot
func WrapTeapot(err error) *HTTPError {
	return Wrap(http.StatusTeapot, err)
}

// HTTP 422 Unprocessable Entity
func NewUnprocessableEntity(message ...string) *HTTPError {
	return New(http.StatusUnprocessableEntity, message...)
}

// HTTP 422 Unprocessable Entity
func WrapUnprocessableEntity(err error) *HTTPError {
	return Wrap(http.StatusUnprocessableEntity, err)
}

// HTTP 423 Locked
func NewLocked(message ...string) *HTTPError {
	return New(http.StatusLocked, message...)
}

// HTTP 423 Locked
func WrapLocked(err error) *HTTPError {
	return Wrap(http.StatusLocked, err)
}

// HTTP 424 Failed Dependency
func NewFailedDependency(message ...string) *HTTPError {
	return New(http.StatusFailedDependency, message...)
}

// HTTP 424 Failed Dependency
func WrapFailedDependency(err error) *HTTPError {
	return Wrap(http.StatusFailedDependency, err)
}

// HTTP 426 Upgrade Required
func NewUpgradeRequired(message ...string) *HTTPError {
	return New(http.StatusUpgradeRequired, message...)
}

// HTTP 426 Upgrade Required
func WrapUpgradeRequired(err error) *HTTPError {
	return Wrap(http.StatusUpgradeRequired, err)
}

// HTTP 428 Precondition Required
func NewPreconditionRequired(message ...string) *HTTPError {
	return New(http.StatusPreconditionRequired, message...)
}

// HTTP 428 Precondition Required
func WrapPreconditionRequired(err error) *HTTPError {
	return Wrap(http.StatusPreconditionRequired, err)
}

// HTTP 429 Too Many Requests
func NewTooManyRequests(message ...string) *HTTPError {
	return New(http.StatusTooManyRequests, message...)
}

// HTTP 429 Too Many Requests
func WrapTooManyRequests(err error) *HTTPError {
	return Wrap(http.StatusTooManyRequests, err)
}

// HTTP 431 Request Header Fields Too Large
func NewRequestHeaderFieldsTooLarge(message ...string) *HTTPError {
	return New(http.StatusRequestHeaderFieldsTooLarge, message...)
}

// HTTP 431 Request Header Fields Too Large
func WrapRequestHeaderFieldsTooLarge(err error) *HTTPError {
	return Wrap(http.StatusRequestHeaderFieldsTooLarge, err)
}

// HTTP 451 Unavailable For Legal Reasons
func NewUnavailableForLegalReasons(message ...string) *HTTPError {
	return New(http.StatusUnavailableForLegalReasons, message...)
}

// HTTP 451 Unavailable For Legal Reasons
func WrapUnavailableForLegalReasons(err error) *HTTPError {
	return Wrap(http.StatusUnavailableForLegalReasons, err)
}

// HTTP 500 Internal Server Error
func NewInternalServerError(message ...string) *HTTPError {
	return New(http.StatusInternalServerError, message...)
}

// HTTP 500 Internal Server Error
func WrapInternalServerError(err error) *HTTPError {
	return Wrap(http.StatusInternalServerError, err)
}

// HTTP 501 Not Implemented
func NewNotImplemented(message ...string) *HTTPError {
	return New(http.StatusNotImplemented, message...)
}

// HTTP 501 Not Implemented
func WrapNotImplemented(err error) *HTTPError {
	return Wrap(http.StatusNotImplemented, err)
}

// HTTP 502 Bad Gateway
func NewBadGateway(message ...string) *HTTPError {
	return New(http.StatusBadGateway, message...)
}

// HTTP 502 Bad Gateway
func WrapBadGateway(err error) *HTTPError {
	return Wrap(http.StatusBadGateway, err)
}

// HTTP 503 Service Unavailable
func NewServiceUnavailable(message ...string) *HTTPError {
	return New(http.StatusServiceUnavailable, message...)
}

// HTTP 503 Service Unavailable
func WrapServiceUnavailable(err error) *HTTPError {
	return Wrap(http.StatusServiceUnavailable, err)
}

// HTTP 504 Gateway Timeout
func NewGatewayTimeout(message ...string) *HTTPError {
	return New(http.StatusGatewayTimeout, message...)
}

// HTTP 504 Gateway Timeout
func WrapGatewayTimeout(err error) *HTTPError {
	return Wrap(http.StatusGatewayTimeout, err)
}

// HTTP 505 HTTP Version Not Supported
func NewHTTPVersionNotSupported(message ...string) *HTTPError {
	return New(http.StatusHTTPVersionNotSupported, message...)
}

// HTTP 505 HTTP Version Not Supported
func WrapHTTPVersionNotSupported(err error) *HTTPError {
	return Wrap(http.StatusHTTPVersionNotSupported, err)
}

// HTTP 506 Variant Also Negotiates
func NewVariantAlsoNegotiates(message ...string) *HTTPError {
	return New(http.StatusVariantAlsoNegotiates, message...)
}

// HTTP 506 Variant Also Negotiates
func WrapVariantAlsoNegotiates(err error) *HTTPError {
	return Wrap(http.StatusVariantAlsoNegotiates, err)
}

// HTTP 507 Insufficient Storage
func NewInsufficientStorage(message ...string) *HTTPError {
	return New(http.StatusInsufficientStorage, message...)
}

// HTTP 507 Insufficient Storage
func WrapInsufficientStorage(err error) *HTTPError {
	return Wrap(http.StatusInsufficientStorage, err)
}

// HTTP 508 Loop Detected
func NewLoopDetected(message ...string) *HTTPError {
	return New(http.StatusLoopDetected, message...)
}

// HTTP 508 Loop Detected
func WrapLoopDetected(err error) *HTTPError {
	return Wrap(http.StatusLoopDetected, err)
}

// HTTP 510 Not Extended
func NewNotExtended(message ...string) *HTTPError {
	return New(http.StatusNotExtended, message...)
}

// HTTP 510 Not Extended
func WrapNotExtended(err error) *HTTPError {
	return Wrap(http.StatusNotExtended, err)
}

// HTTP 511 Network Authentication Required
func NewNetworkAuthenticationRequired(message ...string) *HTTPError {
	return New(http.StatusNetworkAuthenticationRequired, message...)
}

// HTTP 511 Network Authentication Required
func WrapNetworkAuthenticationRequired(err error) *HTTPError {
	return Wrap(http.StatusNetworkAuthenticationRequired, err)
}
