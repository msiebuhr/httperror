package httperror

import "net/http"

// HTTP 100 Continue
func NewContinue(message ...string) *HTTPError {
	return New(http.StatusContinue, message...)
}

// HTTP 100 Continue
func Continue(err error) *HTTPError {
	return Wrap(http.StatusContinue, err)
}

// HTTP 101 Switching Protocols
func NewSwitchingProtocols(message ...string) *HTTPError {
	return New(http.StatusSwitchingProtocols, message...)
}

// HTTP 101 Switching Protocols
func SwitchingProtocols(err error) *HTTPError {
	return Wrap(http.StatusSwitchingProtocols, err)
}

// HTTP 102 Processing
func NewProcessing(message ...string) *HTTPError {
	return New(http.StatusProcessing, message...)
}

// HTTP 102 Processing
func Processing(err error) *HTTPError {
	return Wrap(http.StatusProcessing, err)
}

// HTTP 200 OK
func NewOK(message ...string) *HTTPError {
	return New(http.StatusOK, message...)
}

// HTTP 200 OK
func OK(err error) *HTTPError {
	return Wrap(http.StatusOK, err)
}

// HTTP 201 Created
func NewCreated(message ...string) *HTTPError {
	return New(http.StatusCreated, message...)
}

// HTTP 201 Created
func Created(err error) *HTTPError {
	return Wrap(http.StatusCreated, err)
}

// HTTP 202 Accepted
func NewAccepted(message ...string) *HTTPError {
	return New(http.StatusAccepted, message...)
}

// HTTP 202 Accepted
func Accepted(err error) *HTTPError {
	return Wrap(http.StatusAccepted, err)
}

// HTTP 203 NonAuthoritativeInfo
func NewNonAuthoritativeInfo(message ...string) *HTTPError {
	return New(http.StatusNonAuthoritativeInfo, message...)
}

// HTTP 203 NonAuthoritativeInfo
func NonAuthoritativeInfo(err error) *HTTPError {
	return Wrap(http.StatusNonAuthoritativeInfo, err)
}

// HTTP 204 No Content
func NewNoContent(message ...string) *HTTPError {
	return New(http.StatusNoContent, message...)
}

// HTTP 204 No Content
func NoContent(err error) *HTTPError {
	return Wrap(http.StatusNoContent, err)
}

// HTTP 205 Reset Content
func NewResetContent(message ...string) *HTTPError {
	return New(http.StatusResetContent, message...)
}

// HTTP 205 Reset Content
func ResetContent(err error) *HTTPError {
	return Wrap(http.StatusResetContent, err)
}

// HTTP 206 Partial Content
func NewPartialContent(message ...string) *HTTPError {
	return New(http.StatusPartialContent, message...)
}

// HTTP 206 Partial Content
func PartialContent(err error) *HTTPError {
	return Wrap(http.StatusPartialContent, err)
}

// HTTP 207 Multi-Status
func NewMultiStatus(message ...string) *HTTPError {
	return New(http.StatusMultiStatus, message...)
}

// HTTP 207 Multi-Status
func MultiStatus(err error) *HTTPError {
	return Wrap(http.StatusMultiStatus, err)
}

// HTTP 208 Already Reported
func NewAlreadyReported(message ...string) *HTTPError {
	return New(http.StatusAlreadyReported, message...)
}

// HTTP 208 Already Reported
func AlreadyReported(err error) *HTTPError {
	return Wrap(http.StatusAlreadyReported, err)
}

// HTTP 226 IM Used
func NewIMUsed(message ...string) *HTTPError {
	return New(http.StatusIMUsed, message...)
}

// HTTP 226 IM Used
func IMUsed(err error) *HTTPError {
	return Wrap(http.StatusIMUsed, err)
}

// HTTP 300 Multiple Choices
func NewMultipleChoices(message ...string) *HTTPError {
	return New(http.StatusMultipleChoices, message...)
}

// HTTP 300 Multiple Choices
func MultipleChoices(err error) *HTTPError {
	return Wrap(http.StatusMultipleChoices, err)
}

// HTTP 301 Moved Permanently
func NewMovedPermanently(message ...string) *HTTPError {
	return New(http.StatusMovedPermanently, message...)
}

// HTTP 301 Moved Permanently
func MovedPermanently(err error) *HTTPError {
	return Wrap(http.StatusMovedPermanently, err)
}

// HTTP 302 Found
func NewFound(message ...string) *HTTPError {
	return New(http.StatusFound, message...)
}

// HTTP 302 Found
func Found(err error) *HTTPError {
	return Wrap(http.StatusFound, err)
}

// HTTP 303 See Other
func NewSeeOther(message ...string) *HTTPError {
	return New(http.StatusSeeOther, message...)
}

// HTTP 303 See Other
func SeeOther(err error) *HTTPError {
	return Wrap(http.StatusSeeOther, err)
}

// HTTP 304 Not Modified
func NewNotModified(message ...string) *HTTPError {
	return New(http.StatusNotModified, message...)
}

// HTTP 304 Not Modified
func NotModified(err error) *HTTPError {
	return Wrap(http.StatusNotModified, err)
}

// HTTP 305 Use Proxy
func NewUseProxy(message ...string) *HTTPError {
	return New(http.StatusUseProxy, message...)
}

// HTTP 305 Use Proxy
func UseProxy(err error) *HTTPError {
	return Wrap(http.StatusUseProxy, err)
}

// HTTP 307 Temporary Redirect
func NewTemporaryRedirect(message ...string) *HTTPError {
	return New(http.StatusTemporaryRedirect, message...)
}

// HTTP 307 Temporary Redirect
func TemporaryRedirect(err error) *HTTPError {
	return Wrap(http.StatusTemporaryRedirect, err)
}

// HTTP 308 Permanent Redirect
func NewPermanentRedirect(message ...string) *HTTPError {
	return New(http.StatusPermanentRedirect, message...)
}

// HTTP 308 Permanent Redirect
func PermanentRedirect(err error) *HTTPError {
	return Wrap(http.StatusPermanentRedirect, err)
}

// HTTP 400 Bad Request
func NewBadRequest(message ...string) *HTTPError {
	return New(http.StatusBadRequest, message...)
}

// HTTP 400 Bad Request
func BadRequest(err error) *HTTPError {
	return Wrap(http.StatusBadRequest, err)
}

// HTTP 401 Unauthorized
func NewUnauthorized(message ...string) *HTTPError {
	return New(http.StatusUnauthorized, message...)
}

// HTTP 401 Unauthorized
func Unauthorized(err error) *HTTPError {
	return Wrap(http.StatusUnauthorized, err)
}

// HTTP 402 Payment Required
func NewPaymentRequired(message ...string) *HTTPError {
	return New(http.StatusPaymentRequired, message...)
}

// HTTP 402 Payment Required
func PaymentRequired(err error) *HTTPError {
	return Wrap(http.StatusPaymentRequired, err)
}

// HTTP 403 Forbidden
func NewForbidden(message ...string) *HTTPError {
	return New(http.StatusForbidden, message...)
}

// HTTP 403 Forbidden
func Forbidden(err error) *HTTPError {
	return Wrap(http.StatusForbidden, err)
}

// HTTP 404 Not Found
func NewNotFound(message ...string) *HTTPError {
	return New(http.StatusNotFound, message...)
}

// HTTP 404 Not Found
func NotFound(err error) *HTTPError {
	return Wrap(http.StatusNotFound, err)
}

// HTTP 405 Method Not Allowed
func NewMethodNotAllowed(message ...string) *HTTPError {
	return New(http.StatusMethodNotAllowed, message...)
}

// HTTP 405 Method Not Allowed
func MethodNotAllowed(err error) *HTTPError {
	return Wrap(http.StatusMethodNotAllowed, err)
}

// HTTP 406 Not Acceptable
func NewNotAcceptable(message ...string) *HTTPError {
	return New(http.StatusNotAcceptable, message...)
}

// HTTP 406 Not Acceptable
func NotAcceptable(err error) *HTTPError {
	return Wrap(http.StatusNotAcceptable, err)
}

// HTTP 407 ProxyAuthRequired
func NewProxyAuthRequired(message ...string) *HTTPError {
	return New(http.StatusProxyAuthRequired, message...)
}

// HTTP 407 ProxyAuthRequired
func ProxyAuthRequired(err error) *HTTPError {
	return Wrap(http.StatusProxyAuthRequired, err)
}

// HTTP 408 Request Timeout
func NewRequestTimeout(message ...string) *HTTPError {
	return New(http.StatusRequestTimeout, message...)
}

// HTTP 408 Request Timeout
func RequestTimeout(err error) *HTTPError {
	return Wrap(http.StatusRequestTimeout, err)
}

// HTTP 409 Conflict
func NewConflict(message ...string) *HTTPError {
	return New(http.StatusConflict, message...)
}

// HTTP 409 Conflict
func Conflict(err error) *HTTPError {
	return Wrap(http.StatusConflict, err)
}

// HTTP 410 Gone
func NewGone(message ...string) *HTTPError {
	return New(http.StatusGone, message...)
}

// HTTP 410 Gone
func Gone(err error) *HTTPError {
	return Wrap(http.StatusGone, err)
}

// HTTP 411 Length Required
func NewLengthRequired(message ...string) *HTTPError {
	return New(http.StatusLengthRequired, message...)
}

// HTTP 411 Length Required
func LengthRequired(err error) *HTTPError {
	return Wrap(http.StatusLengthRequired, err)
}

// HTTP 412 Precondition Failed
func NewPreconditionFailed(message ...string) *HTTPError {
	return New(http.StatusPreconditionFailed, message...)
}

// HTTP 412 Precondition Failed
func PreconditionFailed(err error) *HTTPError {
	return Wrap(http.StatusPreconditionFailed, err)
}

// HTTP 413 Request Entity Too Large
func NewRequestEntityTooLarge(message ...string) *HTTPError {
	return New(http.StatusRequestEntityTooLarge, message...)
}

// HTTP 413 Request Entity Too Large
func RequestEntityTooLarge(err error) *HTTPError {
	return Wrap(http.StatusRequestEntityTooLarge, err)
}

// HTTP 414 Request URI Too Long
func NewRequestURITooLong(message ...string) *HTTPError {
	return New(http.StatusRequestURITooLong, message...)
}

// HTTP 414 Request URI Too Long
func RequestURITooLong(err error) *HTTPError {
	return Wrap(http.StatusRequestURITooLong, err)
}

// HTTP 415 Unsupported Media Type
func NewUnsupportedMediaType(message ...string) *HTTPError {
	return New(http.StatusUnsupportedMediaType, message...)
}

// HTTP 415 Unsupported Media Type
func UnsupportedMediaType(err error) *HTTPError {
	return Wrap(http.StatusUnsupportedMediaType, err)
}

// HTTP 416 Requested Range Not Satisfiable
func NewRequestedRangeNotSatisfiable(message ...string) *HTTPError {
	return New(http.StatusRequestedRangeNotSatisfiable, message...)
}

// HTTP 416 Requested Range Not Satisfiable
func RequestedRangeNotSatisfiable(err error) *HTTPError {
	return Wrap(http.StatusRequestedRangeNotSatisfiable, err)
}

// HTTP 417 Expectation Failed
func NewExpectationFailed(message ...string) *HTTPError {
	return New(http.StatusExpectationFailed, message...)
}

// HTTP 417 Expectation Failed
func ExpectationFailed(err error) *HTTPError {
	return Wrap(http.StatusExpectationFailed, err)
}

// HTTP 418 Teapot
func NewTeapot(message ...string) *HTTPError {
	return New(http.StatusTeapot, message...)
}

// HTTP 418 Teapot
func Teapot(err error) *HTTPError {
	return Wrap(http.StatusTeapot, err)
}

// HTTP 422 Unprocessable Entity
func NewUnprocessableEntity(message ...string) *HTTPError {
	return New(http.StatusUnprocessableEntity, message...)
}

// HTTP 422 Unprocessable Entity
func UnprocessableEntity(err error) *HTTPError {
	return Wrap(http.StatusUnprocessableEntity, err)
}

// HTTP 423 Locked
func NewLocked(message ...string) *HTTPError {
	return New(http.StatusLocked, message...)
}

// HTTP 423 Locked
func Locked(err error) *HTTPError {
	return Wrap(http.StatusLocked, err)
}

// HTTP 424 Failed Dependency
func NewFailedDependency(message ...string) *HTTPError {
	return New(http.StatusFailedDependency, message...)
}

// HTTP 424 Failed Dependency
func FailedDependency(err error) *HTTPError {
	return Wrap(http.StatusFailedDependency, err)
}

// HTTP 426 Upgrade Required
func NewUpgradeRequired(message ...string) *HTTPError {
	return New(http.StatusUpgradeRequired, message...)
}

// HTTP 426 Upgrade Required
func UpgradeRequired(err error) *HTTPError {
	return Wrap(http.StatusUpgradeRequired, err)
}

// HTTP 428 Precondition Required
func NewPreconditionRequired(message ...string) *HTTPError {
	return New(http.StatusPreconditionRequired, message...)
}

// HTTP 428 Precondition Required
func PreconditionRequired(err error) *HTTPError {
	return Wrap(http.StatusPreconditionRequired, err)
}

// HTTP 429 Too Many Requests
func NewTooManyRequests(message ...string) *HTTPError {
	return New(http.StatusTooManyRequests, message...)
}

// HTTP 429 Too Many Requests
func TooManyRequests(err error) *HTTPError {
	return Wrap(http.StatusTooManyRequests, err)
}

// HTTP 431 Request Header Fields Too Large
func NewRequestHeaderFieldsTooLarge(message ...string) *HTTPError {
	return New(http.StatusRequestHeaderFieldsTooLarge, message...)
}

// HTTP 431 Request Header Fields Too Large
func RequestHeaderFieldsTooLarge(err error) *HTTPError {
	return Wrap(http.StatusRequestHeaderFieldsTooLarge, err)
}

// HTTP 451 Unavailable For Legal Reasons
func NewUnavailableForLegalReasons(message ...string) *HTTPError {
	return New(http.StatusUnavailableForLegalReasons, message...)
}

// HTTP 451 Unavailable For Legal Reasons
func UnavailableForLegalReasons(err error) *HTTPError {
	return Wrap(http.StatusUnavailableForLegalReasons, err)
}

// HTTP 500 Internal Server Error
func NewInternalServerError(message ...string) *HTTPError {
	return New(http.StatusInternalServerError, message...)
}

// HTTP 500 Internal Server Error
func InternalServerError(err error) *HTTPError {
	return Wrap(http.StatusInternalServerError, err)
}

// HTTP 501 Not Implemented
func NewNotImplemented(message ...string) *HTTPError {
	return New(http.StatusNotImplemented, message...)
}

// HTTP 501 Not Implemented
func NotImplemented(err error) *HTTPError {
	return Wrap(http.StatusNotImplemented, err)
}

// HTTP 502 Bad Gateway
func NewBadGateway(message ...string) *HTTPError {
	return New(http.StatusBadGateway, message...)
}

// HTTP 502 Bad Gateway
func BadGateway(err error) *HTTPError {
	return Wrap(http.StatusBadGateway, err)
}

// HTTP 503 Service Unavailable
func NewServiceUnavailable(message ...string) *HTTPError {
	return New(http.StatusServiceUnavailable, message...)
}

// HTTP 503 Service Unavailable
func ServiceUnavailable(err error) *HTTPError {
	return Wrap(http.StatusServiceUnavailable, err)
}

// HTTP 504 Gateway Timeout
func NewGatewayTimeout(message ...string) *HTTPError {
	return New(http.StatusGatewayTimeout, message...)
}

// HTTP 504 Gateway Timeout
func GatewayTimeout(err error) *HTTPError {
	return Wrap(http.StatusGatewayTimeout, err)
}

// HTTP 505 HTTP Version Not Supported
func NewHTTPVersionNotSupported(message ...string) *HTTPError {
	return New(http.StatusHTTPVersionNotSupported, message...)
}

// HTTP 505 HTTP Version Not Supported
func HTTPVersionNotSupported(err error) *HTTPError {
	return Wrap(http.StatusHTTPVersionNotSupported, err)
}

// HTTP 506 Variant Also Negotiates
func NewVariantAlsoNegotiates(message ...string) *HTTPError {
	return New(http.StatusVariantAlsoNegotiates, message...)
}

// HTTP 506 Variant Also Negotiates
func VariantAlsoNegotiates(err error) *HTTPError {
	return Wrap(http.StatusVariantAlsoNegotiates, err)
}

// HTTP 507 Insufficient Storage
func NewInsufficientStorage(message ...string) *HTTPError {
	return New(http.StatusInsufficientStorage, message...)
}

// HTTP 507 Insufficient Storage
func InsufficientStorage(err error) *HTTPError {
	return Wrap(http.StatusInsufficientStorage, err)
}

// HTTP 508 Loop Detected
func NewLoopDetected(message ...string) *HTTPError {
	return New(http.StatusLoopDetected, message...)
}

// HTTP 508 Loop Detected
func LoopDetected(err error) *HTTPError {
	return Wrap(http.StatusLoopDetected, err)
}

// HTTP 510 Not Extended
func NewNotExtended(message ...string) *HTTPError {
	return New(http.StatusNotExtended, message...)
}

// HTTP 510 Not Extended
func NotExtended(err error) *HTTPError {
	return Wrap(http.StatusNotExtended, err)
}

// HTTP 511 Network Authentication Required
func NewNetworkAuthenticationRequired(message ...string) *HTTPError {
	return New(http.StatusNetworkAuthenticationRequired, message...)
}

// HTTP 511 Network Authentication Required
func NetworkAuthenticationRequired(err error) *HTTPError {
	return Wrap(http.StatusNetworkAuthenticationRequired, err)
}
