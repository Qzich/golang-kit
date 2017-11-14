package api

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

//
// GetHTTPMethod :: with an empty method :: returns empty value
//
func TestGetHTTPMethodForEmptyValue(t *testing.T) {
	r := NewHTTPResource("", "")

	httpMethod := r.GetHTTPMethod()
	httpURL := r.GetURL()

	assert.Empty(t, httpMethod)
	assert.Empty(t, httpURL)
}

//
// GetHTTPMethod :: with a valid method :: returns passed value
//
func TestGetHTTPMethodForAValidValue(t *testing.T) {
	expectedHTTPMethod := http.MethodGet
	r := NewHTTPResource(expectedHTTPMethod, "")

	httpMethod := r.GetHTTPMethod()
	httpURL := r.GetURL()

	assert.Equal(t, expectedHTTPMethod, httpMethod)
	assert.Empty(t, httpURL)
}

//
// GetURL :: with a valid URL :: returns passed value
//
func TestGetURLForAValidValue(t *testing.T) {
	expectedHTTPMethod := http.MethodPost
	expectedURL := "https://api.virgilsecutity.com/cards/v4/card"
	r := NewHTTPResource(expectedHTTPMethod, expectedURL)

	httpMethod := r.GetHTTPMethod()
	httpURL := r.GetURL()

	assert.Equal(t, expectedHTTPMethod, httpMethod)
	assert.Equal(t, expectedURL, httpURL)
}

//
// SetHeader :: without header value set :: returns empty headers list
//
func TestGetHeadersForNotSetHeaders(t *testing.T) {
	r := NewHTTPResource("", "")

	httpHeaders := r.GetHeaders()

	assert.Equal(t, 0, len(httpHeaders))
}

//
// SetHeader :: with an empty value :: returns empty headers list
//
func TestGetHeadersForAnEmptyHeader(t *testing.T) {
	r := NewHTTPResource("", "")
	r.AddHeader("", "")

	httpHeaders := r.GetHeaders()

	assert.Equal(t, 1, len(httpHeaders))
	assert.Empty(t, httpHeaders[""][0])
}

//
// SetHeader :: with a value :: returns passed value
//
func TestGetHeadersForAHeader(t *testing.T) {
	headerName := "name"
	headerValue := "value"
	r := NewHTTPResource("", "")
	r.AddHeader(headerName, headerValue)

	httpHeaders := r.GetHeaders()

	assert.Equal(t, 1, len(httpHeaders))
	assert.Equal(t, headerValue, httpHeaders[headerName][0])
}
