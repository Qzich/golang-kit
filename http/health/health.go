//
// Package health provides functionality to run registered application dependencies checks and to provide an HTTP
// response.
//
// TODO: Add dependency ignoring option to prevent cyclic resolve.
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
	DependencyCassandra = "cassandra"
	DependencyRedis     = "redis"

	DependencyVirgilCrypto  = "virgil_crypto"
	DependencyVirgilCardsV5 = "virgil_cards_v5"
)
