package cfg

import (
	"net/url"
	"strings"

	"fmt"
	"github.com/ameteiko/golang-kit/errors"
)

//
// Cassandra connection URL constants.
//
const (
	CassandraConnectionProtocol = "cassandra"
)

//
// CassandraConnectionInfo is a cassandra database connection info parameter.
//
type CassandraConnectionInfo struct {
	hosts      []string
	keyspace   string
	dataCenter string
	user       string
	password   string

	*StringParameter
}

//
// CassandraConnectionInfoProvider declares all the connection info getters.
//
type CassandraConnectionInfoProvider interface {
	GetHosts() []string
	GetKeyspace() string
	GetDataCenter() string
	GetUser() string
	GetPassword() string

	IsAuthorizationRequired() bool
	IsDCAware() bool

	ParameterInfoProvider
}

//
// newCassandraConnectionInfo returns a new CassandraConnectionInfo object instance.
//
func newCassandraConnectionInfo() *CassandraConnectionInfo {
	return &CassandraConnectionInfo{
		StringParameter: &StringParameter{},
	}
}

//
// GetHosts returns the list of hosts.
//
func (c *CassandraConnectionInfo) GetHosts() []string {

	return c.hosts
}

//
// GetKeyspace returns cassandraKeyspace value.
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
// GetPassword returns cassandraPassword value.
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
	errMsg := fmt.Sprintf(`kit-cfg@CassandraConnectionInfo.validate [connection string (%s)]`, connectionString)

	if "" == connectionString {
		return ErrCassandraConnectionStringIsEmpty
	}

	if connectionInfo, err = url.Parse(connectionString); nil != err {
		return errors.WrapError(ErrCassandraConnectionStringIsIncorrect, errors.WithMessage(err, errMsg))
	}

	if _, err := validateConnectionProtocolClause(connectionInfo, CassandraConnectionProtocol); nil != err {
		return errors.WrapError(ErrCassandraProtocolIsIncorrect, errors.WithMessage(err, errMsg))
	}

	if c.hosts, err = validateCassandraHostsClause(connectionInfo); nil != err {
		return errors.WithMessage(err, errMsg)
	}

	if c.keyspace, err = validateCassandraKeyspaceClause(connectionInfo); nil != err {
		return errors.WithMessage(err, errMsg)
	}

	if c.user, c.password, err = validateCassandraAuthorizationClause(connectionInfo); nil != err {
		return errors.WithMessage(err, errMsg)
	}

	c.dataCenter = validateCassandraDCClause(connectionInfo)

	return nil
}

//
// validateCassandraHostsClause validates the hosts clause.
//
func validateCassandraHostsClause(url *url.URL) ([]string, error) {
	hosts := strings.Split(url.Host, ",")
	if 0 == len(hosts) {
		return nil, ErrCassandraHostsAreEmpty
	}

	return hosts, nil
}

//
// validateCassandraKeyspaceClause validates the cassandraKeyspace clause.
//
func validateCassandraKeyspaceClause(url *url.URL) (string, error) {
	keyspace := strings.Trim(url.Path, "/")
	if "" == keyspace {
		return "", ErrCassandraKeyspaceIsEmpty
	}

	return keyspace, nil
}

//
// validateCassandraAuthorizationClause validates the authorization clause.
//
func validateCassandraAuthorizationClause(url *url.URL) (string, string, error) {
	if url.User != nil {
		pwd, exists := url.User.Password()
		if !exists {
			return "", "", ErrCassandraPasswordIsEmpty
		}

		return url.User.Username(), pwd, nil
	}

	return "", "", nil
}

//
// validateCassandraDCClause validates the data-center clause.
//
func validateCassandraDCClause(url *url.URL) string {
	if "" != url.Query().Get("dc") {
		return url.Query().Get("dc")
	}

	return ""
}
