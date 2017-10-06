package cassandra

//
// ConnectionInfoProvider declares all the connection info getters.
//
type ConnectionInfoProvider interface {
	GetHosts() []string
	GetKeyspace() string
	GetDataCenter() string
	GetUser() string
	GetPassword() string

	IsAuthorizationRequired() bool
	IsDCAware() bool
}
