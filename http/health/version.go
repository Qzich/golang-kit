package health

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
	Tag    string `json:"tag,omitempty"`
	Branch string `json:"branch"`
	Commit string `json:"commit"`
	Date   string `json:"date"`
}

//
// NewBuildVersion returns a new BuildVersion instance.
// TODO: Add date parsing and its options here
//
func NewBuildVersion(tag, branch, commit, date string) BuildVersion {
	return BuildVersion{
		Tag:    tag,
		Branch: branch,
		Commit: commit,
		Date:   date,
	}
}

//
// GetTag returns tag value.
//
func (v BuildVersion) GetTag() string {

	return v.Tag
}

//
// GetBranch returns branch value.
//
func (v BuildVersion) GetBranch() string {

	return v.Branch
}

//
// GetCommit returns commit value.
//
func (v BuildVersion) GetCommit() string {

	return v.Commit
}

//
// GetDate returns time value.
//
func (v BuildVersion) GetDate() string {

	return v.Date
}
