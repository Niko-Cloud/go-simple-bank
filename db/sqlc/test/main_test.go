package test

import (
	"database/sql"
	_ "github.com/lib/pq"
	_ "github.com/yuki/simplebank/db/sqlc/gen"
	db "github.com/yuki/simplebank/db/sqlc/gen"
	"github.com/yuki/simplebank/util"
	"log"
	"os"
	"testing"
)

var testQueries *db.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../..")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = db.New(testDB)

	os.Exit(m.Run())
}
