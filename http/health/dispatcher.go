package health

import (
	"net/http"
)

//
// Dispatcher is a health manager interface.
// It is responsible for dependencies registering and health data providing.
//
type Dispatcher interface {
	//
	// GetInfoHandler returns health info handler.
	//
	GetInfoHandler() http.Handler

	//
	// GetInfoURL returns health info URL.
	//
	GetInfoURL() string

	//
	// GetStatusHandler returns health status handler.
	//
	GetStatusHandler() http.Handler

	//
	// GetStatusURL returns health status URL.
	//
	GetStatusURL() string

	//
	// RegisterDependency registers dependency.
	//
	RegisterDependency(name string, checker DependencyChecker)
}

//
// DispatchManager is an object that dispatches web-service health functionality.
//
type DispatchManager struct {
	deps       []Dependency
	buildInfo  BuildVersionProvider
	urlLocator URLLocator
}

//
// NewDispatcher returns an instance of the DispatchManager.
//
func NewDispatcher(buildInfo BuildVersionProvider) *DispatchManager {
	return &DispatchManager{
		buildInfo:  buildInfo,
		urlLocator: NewURLLocator(),
	}
}

//
// GetInfoHandler returns an instance of the info handler.
//
func (d *DispatchManager) GetInfoHandler() http.Handler {

	return newInfoHandler(d.buildInfo, d.deps)
}

//
// GetStatusHandler returns an instance of the status handler.
//
func (d *DispatchManager) GetStatusHandler() http.Handler {

	return newStatusHandler(d.buildInfo, d.deps)
}

//
// GetInfoURL returns info URL.
//
func (d *DispatchManager) GetInfoURL() string {

	return d.urlLocator.GetInfoURL()
}

//
// GetStatusURL returns status URL.
//
func (d *DispatchManager) GetStatusURL() string {

	return d.urlLocator.GetStatusURL()
}

//
// RegisterDependency registers a dependency to track.
//
func (d *DispatchManager) RegisterDependency(name string, checker DependencyChecker) {
	dep := Dependency{name: name, checker: checker}
	d.deps = append(d.deps, dep)
}
