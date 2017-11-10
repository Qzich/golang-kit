package health

import (
	"net/http"

	"encoding/json"
)

//
// InfoHandler is an HTTP health handler for the info endpoint.
//
type InfoHandler struct {
	baseHandler
}

//
// NewInfoHandler returns a new info handler instance.
//
func newInfoHandler(buildInfo BuildVersionProvider, dependencies []Dependency) *InfoHandler {

	return &InfoHandler{
		baseHandler{buildInfo, dependencies},
	}
}

//
// Handle returns the application status.
//
func (h *InfoHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	areAllDependenciesUp := true
	responseObject := newResponse(h.getBuildInfo())
	for _, dep := range h.deps {
		dependencyInfo := newDependencyInfo()
		dependencyName := dep.GetName()
		latency, err := dep.Check()
		if nil != err {
			dependencyInfo.markFailed()
			areAllDependenciesUp = false
		}
		dependencyInfo.setLatency(latency)

		responseObject.AddDependencyInfo(dependencyName, dependencyInfo)
	}
	resp, err := json.Marshal(responseObject)
	if nil != err {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if !areAllDependenciesUp {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write(resp)
}
