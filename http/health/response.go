package health

import (
	"net/http"
)

//
// response is the most outer wrapper for the health endpoint response.
//
type response struct {
	Status       int                        `json:"status,omitempty"`
	Build        BuildVersion               `json:"build,omitempty"`
	Dependencies map[string]*dependencyInfo `json:"dependencies,omitempty"`
	LatencyUnit  string                     `json:"latency_measure_unit"`
}

//
// dependencyInfo is an entry with general information about the dependency.
//
type dependencyInfo struct {
	Status  int     `json:"status,omitempty"`
	Latency float64 `json:"latency,omitempty"`
}

//
// newResponse returns a new health endpoint response object.
//
func newResponse(build BuildVersionProvider) *response {

	return &response{
		Status: http.StatusOK,
		Build: BuildVersion{
			Date:   build.GetDate(),
			Branch: build.GetBranch(),
			Commit: build.GetCommit(),
			Tag:    build.GetTag(),
		},
		Dependencies: make(map[string]*dependencyInfo),
		LatencyUnit:  "seconds",
	}
}

//
// AddDependencyInfo adds a dependency entry to the response.
//
func (r *response) AddDependencyInfo(name string, info *dependencyInfo) {
	r.Dependencies[name] = info
	if !info.isStatusOk() {
		r.MarkFailed()
	}
}

//
// MarkFailed marks all the request failed.
//
func (r *response) MarkFailed() {
	r.Status = http.StatusBadRequest
}

//
// newDependencyInfo returns an instance of the dependency info.
//
func newDependencyInfo() *dependencyInfo {

	return &dependencyInfo{
		Status: http.StatusOK,
	}
}

//
// markFailed marks dependency as failed.
//
func (i *dependencyInfo) markFailed() {
	i.Status = http.StatusBadRequest
}

//
//
// isStatusOk returns true if dependency status is HTTP 200.
//
func (i *dependencyInfo) isStatusOk() bool {
	return http.StatusOK == i.Status
}

//
// setLatency sets the dependency latency.
//
func (i *dependencyInfo) setLatency(latency float64) {
	i.Latency = latency
}
