package config

import (
	"net/url"
	"strings"

	"github.com/ameteiko/golang-kit/errors"
)

//
// RegisterCassandraParser registers a Cassandra config parser.
//
func (c *Config) RegisterCassandraParser() {

	c.RegisterConfigParameter(ConfigCassandra)
}

//
// Connection URL constants.
//
const (
	ConnectionProtocol = "cassandra"
)

//
// CassandraConnectionInfo is a cassandra database connection info parameter.
//
type CassandraConnectionInfo struct {
	StringParameter

	hosts      []string
	keyspace   string
	dataCenter string
	user       string
	password   string
}

//
// CassandraConnectionInfoProvider declares all the connection info getters.
//
type CassandraConnectionInfoProvider interface {
	ParseConnectionString(string) error

	GetHosts() []string
	GetKeyspace() string
	GetDataCenter() string
	GetUser() string
	GetPassword() string

	IsAuthorizationRequired() bool
	IsDCAware() bool
}

//
// GetHosts returns the list of hosts.
//
func (c *CassandraConnectionInfo) GetHosts() []string {
	return c.hosts
}

//
// GetKeyspace returns keyspace value.
//
func (c *CassandraConnectionInfo) GetKeyspace() string {
	return c.keyspace
}

//
// GetDataCenter returns data center value.
//
func (c *CassandraConnectionInfo) GetDataCenter() string {
	return c.dataCenter
}

//
// GetUser returns user value.
//
func (c *CassandraConnectionInfo) GetUser() string {
	return c.user
}

//
// GetPassword returns password value.
//
func (c *CassandraConnectionInfo) GetPassword() string {
	return c.password
}

//
// IsDCAware returns true if data center info was set.
//
func (c *CassandraConnectionInfo) IsDCAware() bool {
	return "" != c.dataCenter
}

//
// IsAuthorizationRequired returns true if auth info is set in the connection string.
//
func (c *CassandraConnectionInfo) IsAuthorizationRequired() bool {
	return "" != c.user && "" != c.password
}

//
// validate validates the cassandra connection string parameter.
//
func (c *CassandraConnectionInfo) validate() error {
	var err error
	var connectionInfo *url.URL
	connectionString := c.GetValue()

	if "" == connectionString {
		return ErrConnectionStringIsEmpty
	}

	if connectionInfo, err = url.Parse(connectionString); nil != err {
		return errors.Wrapf(
			err,
			`incorrect database connection string (%s)`,
			connectionString,
		)
	}

	if _, err := validateConnectionProtocolClause(connectionInfo); nil != err {
		return errors.Wrapf(
			err,
			`incorrect protocol value for connection string (%s)`,
			connectionString,
		)
	}

	if c.hosts, err = validateHostsClause(connectionInfo); nil != err {
		return errors.Wrapf(
			err,
			`incorrect hosts value for connection string (%s)`,
			connectionString,
		)
	}

	if c.keyspace, err = validateKeyspaceClause(connectionInfo); nil != err {
		return errors.Wrapf(
			err,
			`incorrect keyspace value for connection string (%s)`,
			connectionString,
		)
	}

	if c.user, c.password, err = validateAuthorizationClause(connectionInfo); nil != err {
		return errors.Wrapf(
			err,
			`incorrect authorization info values for connection string (%s)`,
			connectionString,
		)
	}

	c.dataCenter = validateDCClause(connectionInfo)

	return nil
}

//
// validateConnectionProtocolClause validates the connection protocol clause.
//
func validateConnectionProtocolClause(url *url.URL) (string, error) {
	scheme := url.Scheme
	if ConnectionProtocol != scheme {
		return "", errors.Wrapf(
			ErrConnectionStringSchemeIsIncorrect,
			`database connection protocol validation failed for (%s)`,
			scheme,
		)
	}

	return scheme, nil
}

//
// validateHostsClause validates the hosts clause.
//
func validateHostsClause(url *url.URL) ([]string, error) {
	hosts := strings.Split(url.Host, ",")
	if 0 == len(hosts) {
		return nil, ErrConnectionStringHostsAreEmpty
	}

	return hosts, nil
}

//
// validateKeyspaceClause validates the keyspace clause.
//
func validateKeyspaceClause(url *url.URL) (string, error) {
	keyspace := strings.Trim(url.Path, "/")
	if "" == keyspace {
		return "", ErrConnectionStringKeyspaceIsEmpty
	}

	return keyspace, nil
}

//
// validateAuthorizationClause validates the authorization clause.
//
func validateAuthorizationClause(url *url.URL) (string, string, error) {
	if url.User != nil {
		pwd, exists := url.User.Password()
		if !exists {
			return "", "", ErrConnectionStringPasswordIsEmpty
		}

		return url.User.Username(), pwd, nil
	}

	return "", "", nil
}

//
// validateDCClause validates the data-center clause.
//
func validateDCClause(url *url.URL) string {
	if "" != url.Query().Get("dc") {
		return url.Query().Get("dc")
	}

	return ""
}
