//
// Package health provides functionality to run registered application dependencies checks and to provide an HTTP
// response.
//
package health

//
// DependencyChecker is a function type to check the dependency status.
//
type DependencyChecker func() error

//
// Dependency name constants.
//
const (
	DependencyCassandra    = "cassandra"
	DependencyVirgilCrypto = "virgil_crypto"
	DependencyRedis        = "redis"
)
