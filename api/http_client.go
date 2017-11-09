package api

import (
	"io/ioutil"
	"net/http"

	"github.com/ameteiko/golang-kit/errors"
)

//
// HTTPClient interface.
//
type HTTPClient interface {
	//
	// RawCall calls the resource and returns success response bytes.
	//
	RawCall(resource HTTPResourceInfoProvider) ([]byte, error)
}

//
// HTTP performs HTTP calls for external API-services.
//
type HTTP struct{}

//
// NewHTTPClient creates a new HTTP client object.
//
func NewHTTPClient() HTTP {

	return HTTP{}
}

//
// RawCall calls the resource and returns success response bytes.
//
func (c HTTP) RawCall(httpResource HTTPResourceInfoProvider) ([]byte, error) {
	var emptyResponse []byte

	httpRequest, err := http.NewRequest(httpResource.GetHTTPMethod(), httpResource.GetURL(), nil)
	if nil != err {
		return emptyResponse, errors.WrapError(err, ErrRequestCreationError)
	}

	for h, hValues := range httpResource.GetHeaders() {
		for _, hValue := range hValues {
			httpRequest.Header.Add(h, hValue)
		}
	}

	httpClient := http.Client{}
	httpResponse, err := httpClient.Do(httpRequest)
	if nil != err {
		return emptyResponse, errors.WrapError(
			errors.WithMessage(
				err,
				`kit.api@HTTP.Call [url (%s), method (%s), headers (%s)]`,
				httpResource.GetURL(), httpResource.GetHTTPMethod(), httpResource.GetHeaders(),
			),
			ErrRequestToExternalAPI,
		)
	}

	if http.StatusOK != httpResponse.StatusCode {
		responseBody, _ := ioutil.ReadAll(httpResponse.Body)
		return emptyResponse, errors.WrapError(
			errors.Errorf(`kit.api@HTTP.Call [external server responded with not HTTP 200 but (%s) and response body (%s)]`, httpResponse.StatusCode, responseBody),
			ErrRequestToExternalAPI,
		)
	}

	responseBody, err := ioutil.ReadAll(httpResponse.Body)
	if nil != err {
		return emptyResponse, errors.WrapError(
			errors.Errorf(`kit.api@HTTP.Call [external server response read error]`),
			ErrExternalAPIResponseReadError,
		)
	}

	return responseBody, nil
}
