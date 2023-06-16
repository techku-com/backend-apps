package Config

import (
	"bytes"
	"database/sql"
	"log"
	"techku/Constant"
	"time"
)

type connectionPool interface {
	PostgresConnectionPool(Connection interface{})
}

func (env *database) PostgresConnectionPool(Connection interface{}) {
	Con := Connection.(*sql.DB)
	var buffer bytes.Buffer
	buffer.WriteString("postgres://")
	buffer.WriteString(env.Username + ":" + env.Password)
	buffer.WriteString("@")
	buffer.WriteString(env.Host + ":" + env.Port + "/")
	buffer.WriteString(env.Name)
	buffer.WriteString("?sslmode=disable")
	connection_string := buffer.String()
	Connection, err := sql.Open(Constant.POSTGRES, connection_string)
	if err != nil {
		//panic err
		panic(err.Error())
	}
	Connection.(*sql.DB).SetMaxOpenConns(env.Maximum_connection)
	Connection.(*sql.DB).SetConnMaxIdleTime(env.MaximumIdleTime * time.Second)
	*Con = *Connection.(*sql.DB)
	err = Con.Ping()
	if err != nil {
		log.Print(err.Error())
		panic(err.Error())
	}
	return
}
