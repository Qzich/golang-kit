package health

import (
	"encoding/json"
	"net/http"
)

//
// StatusHandler is an HTTP health handler for the status endpoint.
//
type StatusHandler struct {
	baseHandler
}

//
// NewStatusHandler returns a new status handler instance.
//
func newStatusHandler(buildInfo BuildVersionProvider, dependencies []Dependency) *StatusHandler {

	return &StatusHandler{
		baseHandler{buildInfo, dependencies},
	}
}

//
// Handle returns the application status.
//
func (h *StatusHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	areAllDependenciesUp := true
	response := newResponse(h.getBuildInfo())
	for _, dep := range h.getDependencies() {
		_, err := dep.Check()
		if nil != err {
			areAllDependenciesUp = false
		}
	}

	resp, err := json.Marshal(response)
	if nil != err {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if !areAllDependenciesUp {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write(resp)
}
