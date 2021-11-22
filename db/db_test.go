package db

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDBConnection(t *testing.T) {
	Convey("Testing DB connection", t, func() {
		params := ConnectionParams{
			Driver:   "pgx",
			Host:     "localhost",
			Port:     "5432",
			User:     "postgres",
			Password: "kbr199sd5shi",
			DBName:   "postgres",
			SSLMode:  "disable",
		}
		DBConnect(params)
		DBClose()
	})
}

func TestGetDBTables(t *testing.T) {
	Convey("Testing get DB tables", t, func() {
		params := ConnectionParams{
			Driver:   "pgx",
			Host:     "localhost",
			Port:     "5432",
			User:     "postgres",
			Password: "kbr199sd5shi",
			DBName:   "postgres",
			SSLMode:  "disable",
		}
		DBConnect(params)

		postgres, _ := GetDialect("pgx")
		res := postgres.tables()
		t.Log(res)
		DBClose()
	})
}

func TestGetDBTableColumns(t *testing.T) {
	Convey("Testing get DB tables", t, func() {
		params := ConnectionParams{
			Driver:   "pgx",
			Host:     "localhost",
			Port:     "5432",
			User:     "postgres",
			Password: "kbr199sd5shi",
			DBName:   "postgres",
			SSLMode:  "disable",
		}
		DBConnect(params)

		postgres, _ := GetDialect("pgx")
		res := postgres.columns("products")
		t.Log(res)
		DBClose()
	})
}

func TestGetDBTableIndex(t *testing.T) {
	Convey("Testing get DB tables", t, func() {
		params := ConnectionParams{
			Driver:   "pgx",
			Host:     "localhost",
			Port:     "5432",
			User:     "postgres",
			Password: "kbr199sd5shi",
			DBName:   "postgres",
			SSLMode:  "disable",
		}
		DBConnect(params)

		postgres, _ := GetDialect("pgx")
		res := postgres.indexExists("products", "id")
		t.Log(11111111)
		t.Log(res)
		DBClose()
	})
}
