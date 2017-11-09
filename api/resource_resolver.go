package api

import (
	"fmt"
	"net/http"
	"strings"
)

// Resource name constants.
const (
	CollectionsResource         = "collections"
	ActionsResource             = "actions"
	ResourceNestedEntriesURLKey = "include"
)

//
// ResourceResolver is a REST API resource resolver interface.
//
type ResourceResolver interface {
	//
	// GetCreateEntryResource returns an HTTPResource for the create resource request.
	//
	GetCreateEntryResource(baseURL, resource string) *HTTPResource

	//
	// GetReadEntryResource returns an HTTPResource for the get resource request.
	//
	GetReadEntryResource(baseURL, resource, id string) *HTTPResource

	//
	// GetReadEntryResourceWithNestedEntries returns an HTTPResource for the get resource request with nested entries.
	//
	GetReadEntryResourceWithNestedEntries(baseURL, resource, id string, nestedEntries ...string) *HTTPResource

	//
	// GetUpdateEntryResource returns an HTTPResource for the update resource request.
	//
	GetUpdateEntryResource(baseURL, resource, id string) *HTTPResource

	//
	// GetDeleteEntryResource returns an HTTPResource for the delete resource request.
	//
	GetDeleteEntryResource(baseURL, resource, id string) *HTTPResource

	//
	// GetReadEntryNestedCollectionResource returns an HTTPResource for get resource collection entries request.
	//
	GetReadEntryNestedCollectionResource(baseURL, resource, id, collection string) *HTTPResource

	//
	// GetReadEntryNestedCollectionWithNestedEntriesResource returns an HTTPResource for get resource collection entries
	// with nested entries request.
	//
	GetReadEntryNestedCollectionWithNestedEntriesResource(
		baseURL, resource, id, collection string,
		nestedEntries []string,
	) *HTTPResource

	//
	// InvokeResourceAction returns an HTTPResource for invoke resource action request.
	//
	InvokeResourceAction(baseURL, resource, action string) *HTTPResource
}

//
// ResourceResolve resolves resource URLs.
//
type ResourceResolve struct{}

//
// NewResourceResolver returns a new instance of the resource resolver.
//
func NewResourceResolver() ResourceResolve {

	return ResourceResolve{}
}

//
// GetCreateEntryResource returns an HTTPResource for the create resource request.
//
func (r ResourceResolve) GetCreateEntryResource(baseURL, resource string) *HTTPResource {

	return NewHTTPResource(http.MethodPost, r.getResourceURL(baseURL, resource))
}

//
// GetReadEntryResource returns an HTTPResource for the get resource request.
//
func (r ResourceResolve) GetReadEntryResource(baseURL, resource, id string) *HTTPResource {

	return NewHTTPResource(http.MethodGet, r.getResourceEntryURL(baseURL, resource, id))
}

//
// GetReadEntryResourceWithNestedEntries returns an HTTPResource for the get resource request with nested entries.
//
func (r ResourceResolve) GetReadEntryResourceWithNestedEntries(
	baseURL, resource, id string,
	nestedEntries ...string,
) *HTTPResource {
	entries := strings.Join(nestedEntries, ",")
	url := fmt.Sprintf(
		"%s?%s=%s",
		r.getResourceEntryURL(baseURL, resource, id),
		ResourceNestedEntriesURLKey,
		entries,
	)

	return NewHTTPResource(http.MethodGet, url)
}

//
// GetUpdateEntryResource returns an HTTPResource for the update resource request.
//
func (r ResourceResolve) GetUpdateEntryResource(baseURL, resource, id string) *HTTPResource {

	return NewHTTPResource(http.MethodPost, r.getResourceEntryURL(baseURL, resource, id))
}

//
// GetDeleteEntryResource returns an HTTPResource for the delete resource request.
//
func (r ResourceResolve) GetDeleteEntryResource(baseURL, resource, id string) *HTTPResource {

	return NewHTTPResource(http.MethodDelete, r.getResourceEntryURL(baseURL, resource, id))
}

//
// GetReadEntryNestedCollectionResource returns an HTTPResource for get resource collection entries request.
//
func (r ResourceResolve) GetReadEntryNestedCollectionResource(baseURL, resource, id, collection string) *HTTPResource {

	return NewHTTPResource(http.MethodGet, r.getResourceEntryCollectionsURL(baseURL, resource, id, collection))
}

//
// GetReadEntryNestedCollectionWithNestedEntriesResource returns an HTTPResource for get resource collection entries  with nested
// entries request.
//
func (r ResourceResolve) GetReadEntryNestedCollectionWithNestedEntriesResource(
	baseURL, resource, id, collection string,
	nestedEntries []string,
) *HTTPResource {
	entries := strings.Join(nestedEntries, ",")
	url := fmt.Sprintf(
		"%s?%s=%s",
		r.getResourceEntryCollectionsURL(baseURL, resource, id, collection),
		ResourceNestedEntriesURLKey,
		entries,
	)

	return NewHTTPResource(http.MethodGet, url)
}

//
// InvokeResourceAction returns an HTTPResource for invoke resource action request.
//
func (r ResourceResolve) InvokeResourceAction(baseURL, resource, action string) *HTTPResource {
	url := fmt.Sprintf("%s/%s/%s", r.getResourceURL(baseURL, resource), ActionsResource, action)

	return NewHTTPResource(http.MethodPost, url)
}

//
// getResourceEntryURL returns a resource entry URL.
//
func (r ResourceResolve) getResourceEntryURL(baseURL, resource, id string) string {

	return fmt.Sprintf("%s/%s", r.getResourceURL(baseURL, resource), id)
}

//
// getResourceURL returns a resource URL.
//
func (r ResourceResolve) getResourceURL(baseURL, resource string) string {

	return fmt.Sprintf("%s/%s", baseURL, resource)
}

//
// getResourceEntryCollectionsURL returns a resource entry collections URL.
//
func (r ResourceResolve) getResourceEntryCollectionsURL(baseURL, resource, id, collection string) string {

	return fmt.Sprintf("%s/%s/%s", r.getResourceEntryURL(baseURL, resource, id), CollectionsResource, collection)
}
