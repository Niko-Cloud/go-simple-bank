package test

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"simplebank/app/db/sqlc/gen"
	"testing"
)

var testQueries *db.Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = db.New(conn)

	os.Exit(m.Run())
}
