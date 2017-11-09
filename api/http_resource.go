package api

//
// HTTPResourceInfoProvider is an HTTP resource info interface.
//
type HTTPResourceInfoProvider interface {
	//
	// GetHTTPMethod returns HTTP method info.
	//
	GetHTTPMethod() string

	//
	// GetURL returns resource HTTP URL.
	//
	GetURL() string

	//
	// GetHeaders returns the map of HTTP headers.
	//
	GetHeaders() map[string][]string
}

//
// HTTPResource is an object storing HTTP resource information.
//
type HTTPResource struct {
	method  string
	url     string
	headers map[string][]string
	body    []byte
}

//
// NewHTTPResource is an HTTPResource constructor.
//
func NewHTTPResource(method string, url string) *HTTPResource {

	return &HTTPResource{
		method:  method,
		url:     url,
		headers: make(map[string][]string),
	}
}

//
// GetHTTPMethod returns HTTP method info.
//
func (r *HTTPResource) GetHTTPMethod() string {

	return r.method
}

//
// GetURL returns resource HTTP URL.
//
func (r *HTTPResource) GetURL() string {

	return r.url
}

//
// SetBody sets resource HTTP body.
//
func (r *HTTPResource) SetBody(body []byte) {

	r.body = body
}

//
// SetHeader sets the request HTTP header.
//
func (r *HTTPResource) AddHeader(name, value string) {

	r.headers[name] = append(r.headers[name], value)
}

//
// GetHeaders returns set resource HTTP headers.
//
func (r *HTTPResource) GetHeaders() map[string][]string {

	return r.headers
}
