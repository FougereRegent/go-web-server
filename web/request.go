package web

type Verb string
type HttpResponse int
type Header map[string]string

type Request struct {
	HttpVersion string
	Path        string
	Verb        Verb
	Header      Header
	Body        string
}

type Response struct {
}

const (
	GET       Verb = "GET"
	PUT       Verb = "PUT"
	DELETE    Verb = "DELETE"
	POST      Verb = "POST"
	HEAD      Verb = "HEAD"
	UNDEFINED Verb = "UNDEFINED"
)

const (
	HTTP_1_1 string = "1.1"
	HTTP_1_0 string = "1.0"
	HTTP_2_0 string = "2.0"
	HTTP_3_0 string = "3.0"
)

const (
	LEN_GET    int = len(GET)
	LEN_PUT    int = len(PUT)
	LEN_DELETE int = len(DELETE)
	LEN_POST   int = len(POST)
	LEN_HEAD   int = len(HEAD)
)

const (
	CONTINUE                             HttpResponse = 100
	SWITCHING_PROTOCOLS                  HttpResponse = 101
	PROCESSING                           HttpResponse = 102
	EARLY_HINTS                          HttpResponse = 103
	OK                                   HttpResponse = 200
	CREATED                              HttpResponse = 201
	ACCEPTED                             HttpResponse = 202
	NON_AUTHORITATIVE_INFORMATION        HttpResponse = 203
	NO_CONTENT                           HttpResponse = 204
	RESET_CONTENT                        HttpResponse = 205
	PARTIAL_CONTENT                      HttpResponse = 206
	MULTI_STATUS                         HttpResponse = 207
	ALREADY_REPORTED                     HttpResponse = 208
	CONTENT_DIFFERENT                    HttpResponse = 210
	IM_USED                              HttpResponse = 226
	MULTIPLE_CHOICES                     HttpResponse = 300
	MOVED_PERMANENTLY                    HttpResponse = 301
	FOUND                                HttpResponse = 302
	SEE_OTHER                            HttpResponse = 303
	NOT_MODIFIED                         HttpResponse = 304
	USE_PROXY                            HttpResponse = 305
	TEMPORARY_REDIRECT                   HttpResponse = 307
	PERMANENT_REDIRECT                   HttpResponse = 308
	TOO_MANY_REDIRECTS                   HttpResponse = 310
	BAD_REQUEST                          HttpResponse = 400
	UNAUTHORIZED                         HttpResponse = 401
	PAYMENT_REQUIRED                     HttpResponse = 402
	FORBIDDEN                            HttpResponse = 403
	NOT_FOUND                            HttpResponse = 404
	METHOD_NOT_ALLOWED                   HttpResponse = 405
	NOT_ACCEPTABLE                       HttpResponse = 406
	PROXY_AUTHENTICATION_REQUIRED        HttpResponse = 407
	REQUEST_TIME_OUT                     HttpResponse = 408
	CONFLICT                             HttpResponse = 409
	GONE                                 HttpResponse = 410
	LENGTH_REQUIRED                      HttpResponse = 411
	PRECONDITION_FAILED                  HttpResponse = 412
	REQUEST_ENTITY_TOO_LARGE             HttpResponse = 413
	REQUEST_URI_TOO_LONG                 HttpResponse = 414
	UNSUPPORTED_MEDIA_TYPE               HttpResponse = 415
	REQUESTED_RANGE_UNSATISFIABLE        HttpResponse = 416
	EXPECTATION_FAILED                   HttpResponse = 417
	I_M_A_TEAPOT                         HttpResponse = 418
	PAGE_EXPIRED                         HttpResponse = 419
	BAD_MAPPING                          HttpResponse = 421
	UNPROCESSABLE_ENTITY                 HttpResponse = 422
	LOCKED                               HttpResponse = 423
	METHOD_FAILURE                       HttpResponse = 424
	TOO_EARLY                            HttpResponse = 425
	UPGRADE_REQUIRED                     HttpResponse = 426
	INVALID_DIGITAL_SIGNATURE            HttpResponse = 427
	PRECONDITION_REQUIRED                HttpResponse = 428
	TOO_MANY_REQUESTS                    HttpResponse = 429
	REQUEST_HEADER_FIELDS_TOO_LARGE      HttpResponse = 431
	RETRY_WITH                           HttpResponse = 449
	BLOCKED_BY_WINDOWS_PARENTAL_CONTROLS HttpResponse = 450
	UNAVAILABLE_FOR_LEGAL_REASONS        HttpResponse = 451
	UNRECOVERABLE_ERROR                  HttpResponse = 456
	INTERNAL_SERVER_ERROR                HttpResponse = 500
	NOT_IMPLEMENTED                      HttpResponse = 501
	BAD_GATEWAY                          HttpResponse = 502
	SERVICE_UNAVAILABLE                  HttpResponse = 503
	GATEWAY_TIME_OUT                     HttpResponse = 504
	HTTP_VERSION_NOT_SUPPORTED           HttpResponse = 505
	VARIANT_ALSO_NEGOTIATES              HttpResponse = 506
	INSUFFICIENT_STORAGE                 HttpResponse = 507
	LOOP_DETECTED                        HttpResponse = 508
	BANDWIDTH_LIMIT_EXCEEDED             HttpResponse = 509
	NOT_EXTENDED                         HttpResponse = 510
	NETWORK_AUTHENTICATION_REQUIRED      HttpResponse = 511
)
