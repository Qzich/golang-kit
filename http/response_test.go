package http

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse_HttpHandlerResponse(t *testing.T) {
	response := NewResponse()

	body := []byte("Some body to be set")
	base64Body, _ := json.Marshal(body)
	response.SetBody(body)
	assert.Equal(t, base64Body, response.GetBody())

	status := 400
	response.SetStatus(status)
	assert.Equal(t, status, response.GetStatus())
	assert.Equal(t, base64Body, response.GetBody())
}
