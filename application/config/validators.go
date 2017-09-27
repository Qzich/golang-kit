package config

import (
	"net/url"

	"github.com/ameteiko/errors"
)

//
// validateCassandraConnectionProtocolClause validates the connection protocol clause.
//
func validateConnectionProtocolClause(url *url.URL, protocol string) (string, error) {
	scheme := url.Scheme
	if protocol != scheme {
		return "", errors.Wrapf(
			ErrCassandraProtocolIsIncorrect,
			`database connection protocol validation failed for (%s)`,
			scheme,
		)
	}

	return scheme, nil
}
