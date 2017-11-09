package api

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"strings"
)

//
// GetCreateEntryResource :: with empty data :: returns a valid object
//
func TestCreateResourceWithEmptyData(t *testing.T) {
	resolver := NewResourceResolver()

	resource := resolver.GetCreateEntryResource("", "")

	assert.Equal(t, "/", resource.GetURL())
	assert.Equal(t, http.MethodPost, resource.GetHTTPMethod())
}

//
// GetCreateEntryResource :: with concrete resource :: returns a valid object
//
func TestCreateResourceWithResource(t *testing.T) {
	res := "account"

	resolver := NewResourceResolver()

	resource := resolver.GetCreateEntryResource("", res)

	assert.Equal(t, fmt.Sprintf(`/%s`, res), resource.GetURL())
	assert.Equal(t, http.MethodPost, resource.GetHTTPMethod())
}

//
// GetCreateEntryResource :: with concrete resource and base url :: returns a valid object
//
func TestCreateResourceWithResourceAndBaseURL(t *testing.T) {
	res := "account"
	url := "https://virgilsecurity.com"
	resolver := NewResourceResolver()

	resource := resolver.GetCreateEntryResource(url, res)

	assert.Equal(t, fmt.Sprintf(`%s/%s`, url, res), resource.GetURL())
	assert.Equal(t, http.MethodPost, resource.GetHTTPMethod())
}

//
// GetReadEntryResource :: with concrete resource, base url and id :: returns a valid object
//
func TestReadResourceWithResourceAndBaseURL(t *testing.T) {
	res := "account"
	url := "https://virgilsecurity.com"
	id := "ID"
	resolver := NewResourceResolver()

	resource := resolver.GetReadEntryResource(url, res, id)

	assert.Equal(t, fmt.Sprintf(`%s/%s/%s`, url, res, id), resource.GetURL())
	assert.Equal(t, http.MethodGet, resource.GetHTTPMethod())
}

//
// GetUpdateEntryResource :: with concrete resource, base url and id :: returns a valid object
//
func TestUpdateResourceWithResourceAndBaseURL(t *testing.T) {
	res := "account"
	url := "https://virgilsecurity.com"
	id := "ID"
	resolver := NewResourceResolver()

	resource := resolver.GetUpdateEntryResource(url, res, id)

	assert.Equal(t, fmt.Sprintf(`%s/%s/%s`, url, res, id), resource.GetURL())
	assert.Equal(t, http.MethodPost, resource.GetHTTPMethod())
}

//
// GetDeleteEntryResource :: with concrete resource, base url and id :: returns a valid object
//
func TestDeleteResourceWithResourceAndBaseURL(t *testing.T) {
	res := "account"
	url := "https://virgilsecurity.com"
	id := "ID"
	resolver := NewResourceResolver()

	resource := resolver.GetDeleteEntryResource(url, res, id)

	assert.Equal(t, fmt.Sprintf(`%s/%s/%s`, url, res, id), resource.GetURL())
	assert.Equal(t, http.MethodDelete, resource.GetHTTPMethod())
}

//
// GetReadEntryNestedCollectionResource :: with concrete resource, base url and id :: returns a valid object
//
func TestReadEntryNestedCollectionResourceWithResourceAndBaseURL(t *testing.T) {
	res := "account"
	url := "https://virgilsecurity.com"
	id := "ID"
	collection := "apps"
	resolver := NewResourceResolver()
	expectedURL := fmt.Sprintf(`%s/%s/%s/%s/%s`, url, res, id, CollectionsResource, collection)

	resource := resolver.GetReadEntryNestedCollectionResource(url, res, id, collection)

	assert.Equal(t, expectedURL, resource.GetURL())
	assert.Equal(t, http.MethodGet, resource.GetHTTPMethod())
}

//
// GetReadEntryNestedCollectionWithNestedEntriesResource :: with concrete resource, base url and id :: returns a valid object
//
func TestReadEntryNestedCollectionWithNestedEntriesResourceWithResourceAndBaseURL(t *testing.T) {
	res := "account"
	url := "https://virgilsecurity.com"
	id := "ID"
	collection := "apps"
	nestedEntries := []string{"Key", "Page"}
	resolver := NewResourceResolver()
	expectedURL := fmt.Sprintf(
		`%s/%s/%s/%s/%s?%s=%s`,
		url,
		res,
		id,
		CollectionsResource,
		collection,
		ResourceNestedEntriesURLKey,
		strings.Join(nestedEntries, ","),
	)

	resource := resolver.GetReadEntryNestedCollectionWithNestedEntriesResource(url, res, id, collection, nestedEntries)

	assert.Equal(t, expectedURL, resource.GetURL())
	assert.Equal(t, http.MethodGet, resource.GetHTTPMethod())
}

//
// InvokeResourceAction :: with concrete resource, base url and id :: returns a valid object
//
func TestInvokeResourceActionResourceWithResourceAndBaseURL(t *testing.T) {
	res := "account"
	url := "https://virgilsecurity.com"
	action := "action"
	resolver := NewResourceResolver()
	expectedURL := fmt.Sprintf(`%s/%s/%s/%s`, url, res, ActionsResource, action)

	resource := resolver.InvokeResourceAction(url, res, action)

	assert.Equal(t, expectedURL, resource.GetURL())
	assert.Equal(t, http.MethodPost, resource.GetHTTPMethod())
}
