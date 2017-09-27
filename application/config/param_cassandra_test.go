package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//
// Testing constants.
//
const (
	host1    = "host1.com:1234"
	host2    = "host2.com:2345"
	keyspace = "cards"
	dc       = "US"
	username = "user"
	password = "pwd"
)

func TestParseConnectionString_WithAnEmptyString_ReturnsAnError(t *testing.T) {
	ci := new(CassandraConnectionInfo)

	err := ci.validate()

	assert.Error(t, err)
	assert.Equal(t, ErrConnectionStringIsEmpty, err)
}

func TestParseConnectionString_WithAnIncorrectConnectionString_ReturnsAnError(t *testing.T) {
	ci := new(CassandraConnectionInfo)
	ci.value = "Incorrect connection string"

	err := ci.validate()

	assert.Error(t, err)
}

func TestParseConnectionString_WithAnEmptyURLScheme_ReturnsAnError(t *testing.T) {
	ci := new(CassandraConnectionInfo)
	ci.value = "host.com"

	err := ci.validate()

	assert.Error(t, err)
	//helper.AssertError(t, ErrConnectionStringSchemeIsIncorrect, err)
}

func TestParseConnectionString_WithAnIncorrectURLScheme_ReturnsAnError(t *testing.T) {
	ci := new(CassandraConnectionInfo)
	ci.value = "mongodb://host.com"

	err := ci.validate()

	assert.Error(t, err)
	//helper.AssertError(t, ErrConnectionStringSchemeIsIncorrect, err)
}

func TestParseConnectionString_WithAnEmptyKeyspace_ReturnsAnError(t *testing.T) {
	ci := new(CassandraConnectionInfo)
	ci.value = "cassandra://host.com/?dv=US"

	err := ci.validate()

	assert.Error(t, err)
	//helper.AssertError(t, ErrConnectionStringKeyspaceIsEmpty, err)
}

func TestParseConnectionString_WithACorrectConnectionString_Passes(t *testing.T) {
	ci := new(CassandraConnectionInfo)
	ci.value = fmt.Sprintf("cassandra://%s/%s?dc=%s", host1, keyspace, dc)

	err := ci.validate()

	assert.Empty(t, err)
	assert.Equal(t, []string{host1}, ci.GetHosts())
	assert.Equal(t, keyspace, ci.GetKeyspace())
	assert.Equal(t, dc, ci.GetDataCenter())
	assert.Empty(t, ci.GetPassword())
	assert.Empty(t, ci.GetUser())
}

func TestParseConnectionString_WithAUserParameterButWithoutPasswordOne_ReturnsAnError(t *testing.T) {
	ci := new(CassandraConnectionInfo)
	ci.value = fmt.Sprintf(
		"cassandra://%s@%s/%s?dc=%s",
		username,
		host1,
		keyspace,
		dc,
	)

	err := ci.validate()

	assert.Error(t, err)
	//helper.AssertError(t, ErrConnectionStringPasswordIsEmpty, err)
}

func TestParseConnectionString_WithAnAuthParameters_Passes(t *testing.T) {
	ci := new(CassandraConnectionInfo)
	ci.value = fmt.Sprintf(
		"cassandra://%s:%s@%s,%s/%s?dc=%s",
		username,
		password,
		host1,
		host2,
		keyspace,
		dc,
	)

	err := ci.validate()

	assert.Empty(t, err)
	assert.Equal(t, username, ci.GetUser())
	assert.Equal(t, password, ci.GetPassword())
	assert.Equal(t, []string{host1, host2}, ci.GetHosts())
}

func TestParseConnectionString_WithSeveralHosts_Passes(t *testing.T) {
	ci := new(CassandraConnectionInfo)
	ci.value = fmt.Sprintf(
		"cassandra://%s,%s/%s",
		host1,
		host2,
		keyspace,
	)

	err := ci.validate()

	assert.Empty(t, err)
	assert.Equal(t, ci.GetKeyspace(), keyspace)
	assert.Equal(t, ci.GetHosts(), []string{host1, host2})
}

func TestRegisterCassandraParser_WithoutParameters_PassesAndRegistersAParser(t *testing.T) {
	config := NewConfig()

	config.RegisterCassandraParser()
	err := config.Parse()

	assert.Empty(t, err)
}

func TestRegisterCassandraParser_WithParameters_PassesAndRegistersAParser(t *testing.T) {
	// Cassandra connection string
	cassandraConnectionString := "cassandra://127.0.0.1/virgil_card"
	cassandraConnectionEnvVariableName := EnvParameters[ConfigCassandra]
	setEnvVariable(cassandraConnectionEnvVariableName, cassandraConnectionString)
	// Config object
	config := NewConfig()

	config.RegisterCassandraParser()
	errParsing := config.Parse()
	connectionString := config.GetParameterValue(ConfigCassandra)

	assert.Empty(t, errParsing)
	assert.Equal(t, cassandraConnectionString, connectionString)
}
