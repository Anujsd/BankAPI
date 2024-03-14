package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	conn, err := pgxpool.New(context.Background(), dbSource)

	if err != nil {
		log.Fatal("connection to db failed", err)
	}

	testQueries = New(conn)
	os.Exit(m.Run())
}
