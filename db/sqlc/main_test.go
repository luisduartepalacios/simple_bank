package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/luisduartepalacios/simple_bank/util"
)

var testQueris *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueris = New(testDB)

	os.Exit(m.Run())
}
