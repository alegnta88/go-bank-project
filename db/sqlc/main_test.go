package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const dbDriver = "postgres"
const dbSource = "postgresql://admin:secret@147.182.241.168:5432/go_test?sslmode=disable"

var testQueries *Queries

func TestMain(m *testing.M) {

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("can not connect to database", err)
	}
	testQueries = New(conn)
	os.Exit(m.Run())
}
