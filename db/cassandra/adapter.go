package cassandra

import (
	"github.com/gocql/gocql"
)

//
// Database constants.
//
const (
	DefaultConsistencyLevel = gocql.LocalQuorum
)

//
// Adapter interface wraps the access to the database.
// It is used to decouple the database layer.
//
type Adapter interface {
	Query(stmt string, values ...interface{}) *gocql.Query
	Close()
}

//
// DBAdapter is the database DBAdapter object.
//
type DBAdapter struct {
	session     *gocql.Session
	consistency gocql.Consistency
}

//
// NewDBAdapter returns an instance of the database DBAdapter.
//
func NewDBAdapter(dbSession *gocql.Session, consistency gocql.Consistency) *DBAdapter {
	return &DBAdapter{dbSession, consistency}
}

//
// Query function wraps the built-in query function to apply a centralized consistency management.
// We use LOCAL QUORUM consistency here in order to provide a strong consistency for the cluster requests.
//
func (db *DBAdapter) Query(stmt string, values ...interface{}) *gocql.Query {
	return db.session.Query(stmt, values...).Consistency(db.consistency)
}

//
// Close closes the session connection.
//
func (db *DBAdapter) Close() {
	db.session.Close()
}
