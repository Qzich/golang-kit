package config

import (
	"net/url"

	"github.com/ameteiko/errors"
)

//
// validateCassandraConnectionProtocolClause validates the connection protocol clause.
//
func validateConnectionProtocolClause(url *url.URL, expectedProtocol string) (string, error) {
	scheme := url.Scheme
	if expectedProtocol != scheme {

		return "", errors.Errorf(
			`connection protocol (%s) value is not (%s)`,
			scheme,
			expectedProtocol,
		)
	}

	return scheme, nil
}
