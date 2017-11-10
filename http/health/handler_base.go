package health

//
// baseHandler is a base HTTP health handler.
//
type baseHandler struct {
	buildInfo BuildVersionProvider
	deps      []Dependency
}

//
// getBuildInfo returns a build info.
//
func (h *baseHandler) getBuildInfo() BuildVersionProvider {

	return h.buildInfo
}

//
// getDependencies returns a dependencies list.
//
func (h *baseHandler) getDependencies() []Dependency {

	return h.deps
}
