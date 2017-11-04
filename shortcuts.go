package httperror

import "net/http"

// HTTP 100 Continue
func Continue(message ...string) *HTTPError {
	return New(http.StatusContinue, message...)
}

// HTTP 101 Switching Protocols
func SwitchingProtocols(message ...string) *HTTPError {
	return New(http.StatusSwitchingProtocols, message...)
}

// HTTP 102 Processing
func Processing(message ...string) *HTTPError {
	return New(http.StatusProcessing, message...)
}

// HTTP 200 OK
func OK(message ...string) *HTTPError {
	return New(http.StatusOK, message...)
}

// HTTP 201 Created
func Created(message ...string) *HTTPError {
	return New(http.StatusCreated, message...)
}

// HTTP 202 Accepted
func Accepted(message ...string) *HTTPError {
	return New(http.StatusAccepted, message...)
}

// HTTP 203 NonAuthoritativeInfo
func NonAuthoritativeInfo(message ...string) *HTTPError {
	return New(http.StatusNonAuthoritativeInfo, message...)
}

// HTTP 204 No Content
func NoContent(message ...string) *HTTPError {
	return New(http.StatusNoContent, message...)
}

// HTTP 205 Reset Content
func ResetContent(message ...string) *HTTPError {
	return New(http.StatusResetContent, message...)
}

// HTTP 206 Partial Content
func PartialContent(message ...string) *HTTPError {
	return New(http.StatusPartialContent, message...)
}

// HTTP 207 Multi-Status
func MultiStatus(message ...string) *HTTPError {
	return New(http.StatusMultiStatus, message...)
}

// HTTP 208 Already Reported
func AlreadyReported(message ...string) *HTTPError {
	return New(http.StatusAlreadyReported, message...)
}

// HTTP 226 IM Used
func IMUsed(message ...string) *HTTPError {
	return New(http.StatusIMUsed, message...)
}

// HTTP 300 Multiple Choices
func MultipleChoices(message ...string) *HTTPError {
	return New(http.StatusMultipleChoices, message...)
}

// HTTP 301 Moved Permanently
func MovedPermanently(message ...string) *HTTPError {
	return New(http.StatusMovedPermanently, message...)
}

// HTTP 302 Found
func Found(message ...string) *HTTPError {
	return New(http.StatusFound, message...)
}

// HTTP 303 See Other
func SeeOther(message ...string) *HTTPError {
	return New(http.StatusSeeOther, message...)
}

// HTTP 304 Not Modified
func NotModified(message ...string) *HTTPError {
	return New(http.StatusNotModified, message...)
}

// HTTP 305 Use Proxy
func UseProxy(message ...string) *HTTPError {
	return New(http.StatusUseProxy, message...)
}

// HTTP 307 Temporary Redirect
func TemporaryRedirect(message ...string) *HTTPError {
	return New(http.StatusTemporaryRedirect, message...)
}

// HTTP 308 Permanent Redirect
func PermanentRedirect(message ...string) *HTTPError {
	return New(http.StatusPermanentRedirect, message...)
}

// HTTP 400 Bad Request
func BadRequest(message ...string) *HTTPError {
	return New(http.StatusBadRequest, message...)
}

// HTTP 401 Unauthorized
func Unauthorized(message ...string) *HTTPError {
	return New(http.StatusUnauthorized, message...)
}

// HTTP 402 Payment Required
func PaymentRequired(message ...string) *HTTPError {
	return New(http.StatusPaymentRequired, message...)
}

// HTTP 403 Forbidden
func Forbidden(message ...string) *HTTPError {
	return New(http.StatusForbidden, message...)
}

// HTTP 404 Not Found
func NotFound(message ...string) *HTTPError {
	return New(http.StatusNotFound, message...)
}

// HTTP 405 Method Not Allowed
func MethodNotAllowed(message ...string) *HTTPError {
	return New(http.StatusMethodNotAllowed, message...)
}

// HTTP 406 Not Acceptable
func NotAcceptable(message ...string) *HTTPError {
	return New(http.StatusNotAcceptable, message...)
}

// HTTP 407 ProxyAuthRequired
func ProxyAuthRequired(message ...string) *HTTPError {
	return New(http.StatusProxyAuthRequired, message...)
}

// HTTP 408 Request Timeout
func RequestTimeout(message ...string) *HTTPError {
	return New(http.StatusRequestTimeout, message...)
}

// HTTP 409 Conflict
func Conflict(message ...string) *HTTPError {
	return New(http.StatusConflict, message...)
}

// HTTP 410 Gone
func Gone(message ...string) *HTTPError {
	return New(http.StatusGone, message...)
}

// HTTP 411 Length Required
func LengthRequired(message ...string) *HTTPError {
	return New(http.StatusLengthRequired, message...)
}

// HTTP 412 Precondition Failed
func PreconditionFailed(message ...string) *HTTPError {
	return New(http.StatusPreconditionFailed, message...)
}

// HTTP 413 Request Entity Too Large
func RequestEntityTooLarge(message ...string) *HTTPError {
	return New(http.StatusRequestEntityTooLarge, message...)
}

// HTTP 414 Request URI Too Long
func RequestURITooLong(message ...string) *HTTPError {
	return New(http.StatusRequestURITooLong, message...)
}

// HTTP 415 Unsupported Media Type
func UnsupportedMediaType(message ...string) *HTTPError {
	return New(http.StatusUnsupportedMediaType, message...)
}

// HTTP 416 Requested Range Not Satisfiable
func RequestedRangeNotSatisfiable(message ...string) *HTTPError {
	return New(http.StatusRequestedRangeNotSatisfiable, message...)
}

// HTTP 417 Expectation Failed
func ExpectationFailed(message ...string) *HTTPError {
	return New(http.StatusExpectationFailed, message...)
}

// HTTP 418 Teapot
func Teapot(message ...string) *HTTPError {
	return New(http.StatusTeapot, message...)
}

// HTTP 422 Unprocessable Entity
func UnprocessableEntity(message ...string) *HTTPError {
	return New(http.StatusUnprocessableEntity, message...)
}

// HTTP 423 Locked
func Locked(message ...string) *HTTPError {
	return New(http.StatusLocked, message...)
}

// HTTP 424 Failed Dependency
func FailedDependency(message ...string) *HTTPError {
	return New(http.StatusFailedDependency, message...)
}

// HTTP 426 Upgrade Required
func UpgradeRequired(message ...string) *HTTPError {
	return New(http.StatusUpgradeRequired, message...)
}

// HTTP 428 Precondition Required
func PreconditionRequired(message ...string) *HTTPError {
	return New(http.StatusPreconditionRequired, message...)
}

// HTTP 429 Too Many Requests
func TooManyRequests(message ...string) *HTTPError {
	return New(http.StatusTooManyRequests, message...)
}

// HTTP 431 Request Header Fields Too Large
func RequestHeaderFieldsTooLarge(message ...string) *HTTPError {
	return New(http.StatusRequestHeaderFieldsTooLarge, message...)
}

// HTTP 451 Unavailable For Legal Reasons
func UnavailableForLegalReasons(message ...string) *HTTPError {
	return New(http.StatusUnavailableForLegalReasons, message...)
}

// HTTP 500 Internal Server Error
func InternalServerError(message ...string) *HTTPError {
	return New(http.StatusInternalServerError, message...)
}

// HTTP 501 Not Implemented
func NotImplemented(message ...string) *HTTPError {
	return New(http.StatusNotImplemented, message...)
}

// HTTP 502 Bad Gateway
func BadGateway(message ...string) *HTTPError {
	return New(http.StatusBadGateway, message...)
}

// HTTP 503 Service Unavailable
func ServiceUnavailable(message ...string) *HTTPError {
	return New(http.StatusServiceUnavailable, message...)
}

// HTTP 504 Gateway Timeout
func GatewayTimeout(message ...string) *HTTPError {
	return New(http.StatusGatewayTimeout, message...)
}

// HTTP 505 HTTP Version Not Supported
func HTTPVersionNotSupported(message ...string) *HTTPError {
	return New(http.StatusHTTPVersionNotSupported, message...)
}

// HTTP 506 Variant Also Negotiates
func VariantAlsoNegotiates(message ...string) *HTTPError {
	return New(http.StatusVariantAlsoNegotiates, message...)
}

// HTTP 507 Insufficient Storage
func InsufficientStorage(message ...string) *HTTPError {
	return New(http.StatusInsufficientStorage, message...)
}

// HTTP 508 Loop Detected
func LoopDetected(message ...string) *HTTPError {
	return New(http.StatusLoopDetected, message...)
}

// HTTP 510 Not Extended
func NotExtended(message ...string) *HTTPError {
	return New(http.StatusNotExtended, message...)
}

// HTTP 511 Network Authentication Required
func NetworkAuthenticationRequired(message ...string) *HTTPError {
	return New(http.StatusNetworkAuthenticationRequired, message...)
}
