package Config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type DbSqlConfigName string

const (
	// Database Connection Constant name
	databaseMain DbSqlConfigName = "mainDB"
)

type DbConInterface interface {
	PostgreMainCon() *sql.DB
	setSqlConnection(conName DbSqlConfigName, con *sql.DB)
}

type DbCon struct {
	sql map[DbSqlConfigName]*sql.DB
}

func (d DbCon) PostgreMainCon() *sql.DB {
	return d.sql[databaseMain]
}

func (d *DbCon) setSqlConnection(conName DbSqlConfigName, con *sql.DB) {
	if d.sql == nil {
		d.sql = make(map[DbSqlConfigName]*sql.DB)
	}
	d.sql[conName] = con
}
