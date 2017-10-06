//
// Package version provides all meaningful information about the application build and version.
//
package version

//
// Data parsing constants
//
const (
	TravisBuildDateFormat = "200601021504"
)

//
// BuildVersionProvider provides with build info.
//
type BuildVersionProvider interface {
	//
	// GetVersion returns tag value.
	//
	GetTag() string

	//
	// GetBranch returns branch value.
	//
	GetBranch() string

	//
	// GetCommit returns commit value.
	//
	GetCommit() string

	//
	// GetDate returns time value.
	//
	GetDate() string
}

//
// BuildVersion provides all necessary information about the build version.
//
type BuildVersion struct {
	tag    string
	branch string
	commit string
	date   string
}

//
// NewBuildVersion returns a new BuildVersion instance.
// TODO: Add date parsing and its options here
//
func NewBuildVersion(tag, branch, commit, date string) BuildVersion {
	return BuildVersion{
		tag:    tag,
		branch: branch,
		commit: commit,
		date:   date,
	}
}

//
// GetTag returns tag value.
//
func (v BuildVersion) GetTag() string {

	return v.tag
}

//
// GetBranch returns branch value.
//
func (v BuildVersion) GetBranch() string {

	return v.branch
}

//
// GetCommit returns commit value.
//
func (v BuildVersion) GetCommit() string {

	return v.commit
}

//
// GetDate returns time value.
//
func (v BuildVersion) GetDate() string {

	return v.date
}
