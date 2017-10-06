package cassandra

import (
	"github.com/gocql/gocql"

	"github.com/ameteiko/golang-kit/errors"
)

//
// Connector is a database connector interface.
//
type Connector interface {
	Connect(connectionInfo ConnectionInfoProvider) (*DBAdapter, error)
}

//
// Connection is a cassandra database connection class.
//
type Connection struct {
	consistency gocql.Consistency
}

//
// NewConnection returns a new cassandra connection instance.
//
func NewConnection() *Connection {
	return &Connection{DefaultConsistencyLevel}
}

//
// Connect connects to the cassandra database cluster.
//
func (c *Connection) Connect(connectionInfo ConnectionInfoProvider) (*DBAdapter, error) {
	cluster := gocql.NewCluster(connectionInfo.GetHosts()...)
	cluster.Keyspace = connectionInfo.GetKeyspace()

	if connectionInfo.IsAuthorizationRequired() {
		cluster.Authenticator = gocql.PasswordAuthenticator{
			Username: connectionInfo.GetUser(),
			Password: connectionInfo.GetPassword(),
		}
	}

	if connectionInfo.IsDCAware() {
		cluster.PoolConfig = gocql.PoolConfig{
			HostSelectionPolicy: gocql.DCAwareRoundRobinPolicy(connectionInfo.GetDataCenter()),
		}
	}

	dbSession, err := cluster.CreateSession()
	if err != nil {
		return nil, errors.Wrap(err, "was unable to establish a connection to the database cluster")

	}

	return NewDBAdapter(dbSession, c.consistency), nil
}
