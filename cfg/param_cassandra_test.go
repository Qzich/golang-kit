package cfg

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ameteiko/golang-kit/test/helper"
)

//
// Testing constants.
//
const (
	cassandraHost1    = "host1.com:1234"
	cassandraHost2    = "host2.com:2345"
	cassandraKeyspace = "cards"
	cassandraDc       = "US"
	cassandraUsername = "user"
	cassandraPassword = "pwd"
)

func TestCassandraParseConnectionString_WithAnEmptyString_ReturnsAnError(t *testing.T) {
	ci := new(CassandraConnectionInfo)

	err := ci.validate()

	assert.Error(t, err)
	helper.AssertError(t, ErrCassandraConnectionStringIsEmpty, err)
}

func TestCassandraParseConnectionString_WithAnIncorrectConnectionString_ReturnsAnError(t *testing.T) {
	incorrectConnectionString := "*:?//"
	ci := new(CassandraConnectionInfo)
	ci.value = incorrectConnectionString

	err := ci.validate()

	assert.Error(t, err)
	helper.AssertError(t, ErrCassandraConnectionStringIsIncorrect, err)
}

func TestCassandraParseConnectionString_WithAnEmptyURLScheme_ReturnsAnError(t *testing.T) {
	ci := new(CassandraConnectionInfo)
	ci.value = "host.com"

	err := ci.validate()

	assert.Error(t, err)
	helper.AssertError(t, ErrCassandraProtocolIsIncorrect, err)
}

func TestCassandraParseConnectionString_WithAnIncorrectURLScheme_ReturnsAnError(t *testing.T) {
	ci := new(CassandraConnectionInfo)
	ci.value = "mongodb://host.com"

	err := ci.validate()

	assert.Error(t, err)
	helper.AssertError(t, ErrCassandraProtocolIsIncorrect, err)
}

func TestCassandraParseConnectionString_WithAnEmptyKeyspace_ReturnsAnError(t *testing.T) {
	ci := new(CassandraConnectionInfo)
	ci.value = "cassandra://host.com/?dv=US"

	err := ci.validate()

	assert.Error(t, err)
	helper.AssertError(t, ErrCassandraKeyspaceIsEmpty, err)
}

func TestCassandraParseConnectionString_WithACorrectConnectionString_Passes(t *testing.T) {
	ci := new(CassandraConnectionInfo)
	ci.value = fmt.Sprintf("cassandra://%s/%s?dc=%s", cassandraHost1, cassandraKeyspace, cassandraDc)

	err := ci.validate()

	assert.Empty(t, err)
	assert.Equal(t, []string{cassandraHost1}, ci.GetHosts())
	assert.Equal(t, cassandraKeyspace, ci.GetKeyspace())
	assert.Equal(t, cassandraDc, ci.GetDataCenter())
	assert.Empty(t, ci.GetPassword())
	assert.Empty(t, ci.GetUser())
}

func TestCassandraParseConnectionString_WithAUserParameterButWithoutPasswordOne_ReturnsAnError(t *testing.T) {
	ci := new(CassandraConnectionInfo)
	ci.value = fmt.Sprintf(
		"cassandra://%s@%s/%s?cassandraDc=%s",
		cassandraUsername,
		cassandraHost1,
		cassandraKeyspace,
		cassandraDc,
	)

	err := ci.validate()

	assert.Error(t, err)
	helper.AssertError(t, ErrCassandraPasswordIsEmpty, err)
}

func TestCassandraParseConnectionString_WithAnAuthParameters_Passes(t *testing.T) {
	ci := new(CassandraConnectionInfo)
	ci.value = fmt.Sprintf(
		"cassandra://%s:%s@%s,%s/%s?cassandraDc=%s",
		cassandraUsername,
		cassandraPassword,
		cassandraHost1,
		cassandraHost2,
		cassandraKeyspace,
		cassandraDc,
	)

	err := ci.validate()

	assert.Empty(t, err)
	assert.Equal(t, cassandraUsername, ci.GetUser())
	assert.Equal(t, cassandraPassword, ci.GetPassword())
	assert.Equal(t, []string{cassandraHost1, cassandraHost2}, ci.GetHosts())
}

func TestCassandraParseConnectionString_WithSeveralHosts_Passes(t *testing.T) {
	ci := new(CassandraConnectionInfo)
	ci.value = fmt.Sprintf(
		"cassandra://%s,%s/%s",
		cassandraHost1,
		cassandraHost2,
		cassandraKeyspace,
	)

	err := ci.validate()

	assert.Empty(t, err)
	assert.Equal(t, ci.GetKeyspace(), cassandraKeyspace)
	assert.Equal(t, ci.GetHosts(), []string{cassandraHost1, cassandraHost2})
}

func TestCassandraRegisterCassandraParser_WithoutParameters_PassesAndRegistersAParser(t *testing.T) {
	config := NewConfig()

	config.Register(Cassandra)
	err := config.Parse()

	assert.Empty(t, err)
}

func TestCassandraRegisterCassandraParser_WithParameters_PassesAndRegistersAParser(t *testing.T) {
	// Cassandra connection string
	cassandraConnectionString := "cassandra://127.0.0.1/virgil_card"
	cassandraConnectionEnvVariableName := string(Cassandra)
	setEnvVariable(cassandraConnectionEnvVariableName, cassandraConnectionString)
	// Config object
	config := NewConfig()

	config.Register(Cassandra)
	errParsing := config.Parse()
	connectionString := config.GetValue(Cassandra)

	assert.Empty(t, errParsing)
	assert.Equal(t, cassandraConnectionString, connectionString)
}
