package http

import (
	"net/http"

	"github.com/bmizerany/pat"

	"github.com/ameteiko/golang-kit/errors"
	"github.com/ameteiko/golang-kit/log"
)

//
// HTTPHandlerProvider interface provides an HTTP handler.
//
type HTTPHandlerProvider interface {
	GetHTTPHandler() http.Handler
}

//
// Server is an HTTP server interface. Describes the service that handles the HTTP requests.
//
type Server interface {
	//
	// Get registers an HTTP GET handler.
	//
	Get(path string, handler RequestHandler)

	//
	// Post registers an HTTP POST handler.
	//
	Post(path string, handler RequestHandler)

	//
	// Put registers an HTTP PUT handler.
	//
	Put(path string, handler RequestHandler)

	//
	// Delete registers an HTTP DELETE handler.
	//
	Delete(path string, handler RequestHandler)
}

//
// Router represents router class.
// It registers all the HTTP and error handlers used for the requests serving and implements the Server interface.
//
type Router struct {
	httpHandler   *pat.PatternServeMux
	requestReader requestReader
	log           log.Logger
}

//
// NewRouter returns an instance of request multiplexer.
//
func NewRouter(log log.Logger) *Router {
	r := Router{
		httpHandler:   pat.New(),
		requestReader: new(requestRead),
		log:           log,
	}
	r.httpHandler.NotFound = http.HandlerFunc(defaultNotFoundHandler)

	return &r
}

//
// Get registers an HTTP GET httpHandler.
//
func (r *Router) Get(path string, handler RequestHandler) {
	r.httpHandler.Get(path, r.wrapHTTPHandler(handler))
}

//
// Post registers an HTTP POST httpHandler.
//
func (r *Router) Post(path string, handler RequestHandler) {
	r.httpHandler.Post(path, r.wrapHTTPHandler(handler))
}

//
// Put registers an HTTP PUT httpHandler.
//
func (r *Router) Put(path string, handler RequestHandler) {
	r.httpHandler.Post(path, r.wrapHTTPHandler(handler))
}

//
// Delete registers an HTTP DELETE httpHandler.
//
func (r *Router) Delete(path string, handler RequestHandler) {
	r.httpHandler.Post(path, r.wrapHTTPHandler(handler))
}

//
// GetHTTPHandler returns an httpHandler instance.
//
func (r *Router) GetHTTPHandler() http.Handler {

	return r.httpHandler
}

//
// wrapHTTPHandler wraps an HTTP request handler with a universal wrapper.
//
// Helper function to read the request body if any and to pass it to the HTTP httpHandler.
//
func (r *Router) wrapHTTPHandler(handler RequestHandler) http.HandlerFunc {

	return func(response http.ResponseWriter, request *http.Request) {

		var err error

		//requestBody, err := r.requestReader.readBody(request)
		if err != nil {
			r.log.Debug("%+v", err)
			WriteResponseError(response, errors.ErrRequestRead)

			return
		}

		if request, err = r.injectResource(handler, request, response); nil != err {
			r.log.Debug("%s\n", err)
			// ITODO: think on stack traces logging
			//r.log.Debugf("%+v\n", err)
			WriteResponseError(response, errors.ErrNotFound)
			return
		}

		// Inject applicationIDs
		//_, ok := handler.(core_http.AuthorizationInfoReader)
		//if ok {
		//	 TODO: Fix this request reader concurrency issues.
		//request = core_http.WriteApplicationIDsToContext(
		//	request,
		//	r.requestReader.ReadScopeApplicationIDs(request),
		//)
		//}

		// Handle the request and return an process an error if any.
		handlerResponse := NewResponse()
		//if err := handler.Handle(requestBody, handlerResponse, request); nil != err {
		// TODO: log headers and request body
		//r.handleError(response, err)
		//return
		//}

		response.WriteHeader(handlerResponse.GetStatus())
		response.Write(handlerResponse.GetBody())
	}
}

//
// handleError handles the HTTP httpHandler error.
//
func (r *Router) handleError(response http.ResponseWriter, err error) {

	switch e := errors.Cause(err).(type) {
	case errors.HTTPErrorInfoProvider:
		WriteResponseError(response, e)
		r.log.Debug("%+s\n", err)
	default:
		WriteResponseError(response, errors.ErrInternalServerError)
		r.log.Debug("%+v\n", err)
	}
}

//
// injectResource injects a Virgil Card resource if action requires an instance.
//
func (r *Router) injectResource(
	handler RequestHandler,
	request *http.Request,
	response http.ResponseWriter,
) (*http.Request, error) {
	//switch handler.(type) {
	//case core_http.ResourceRetriever:
	//	cardID := r.requestReader.readURLParameter(request, ResourceIDPart)
	//	cardDTO := model.NewCardDTO()
	//	if err := r.cardDAL.GetCardByID(cardID, cardDTO); nil != err {
	//		err := errors.Wrapf(err, "HTTP request to the resource that cannot be found in the database by its ID(%s)", cardID)
	//		return nil, err
	//	}
	//	request = core_http.SetResourceIntoTheRequestContext(request, cardDTO)
	//}

	return request, nil
}

//
// defaultNotFoundHandler handles all HTTP Not Found errors.
//
func defaultNotFoundHandler(responseWriter http.ResponseWriter, _ *http.Request) {
	WriteResponseError(responseWriter, errors.ErrNotFound)
}
