package health

//
// URLLocator is a health checker URLs locator interface.
// It is used for a seamless integrations the the HTTP router.
//
type URLLocator interface {
	//
	// GetInfoURL returns an info URL.
	//
	GetInfoURL() string

	//
	// GetStatusURL returns a status URL.
	//
	GetStatusURL() string
}

//
// URL is a resource url locator object.
//
type URL struct {
}

//
// NewURLLocator returns a new URL locator instance.
//
func NewURLLocator() *URL {

	return &URL{}
}

//
// GetInfoURL returns info URL.
//
func (h *URL) GetInfoURL() string {

	return "/health/info"
}

//
// GetStatusURL returns status URL.
//
func (h *URL) GetStatusURL() string {

	return "/health/status"
}
