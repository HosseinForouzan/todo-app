package psqltask

import "graph/repository/psql"

type DB struct {
	conn *psql.PsqlDB
}

func New(conn *psql.PsqlDB) *DB {
	return &DB{conn: conn}
}