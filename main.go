package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/nasjp/mysql-on-go-for-multitenancy/partitioning"
	"github.com/nasjp/mysql-on-go-for-multitenancy/table"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {
	db, err := sqlx.Open("mysql", "root@tcp(db:3306)/test_db")
	if err != nil {
		return err
	}
	defer db.Close()
	switch command := os.Getenv("COMMAND"); command {
	case "table":
		fmt.Println("Start table")
		if err := table.Run(db); err != nil {
			return err
		}
	case "partitioning":
		fmt.Println("Start partitioning")
		if err := partitioning.Run(db); err != nil {
			return err
		}
	default:
		return fmt.Errorf("Unexpected command: %s", command)
	}

	fmt.Println("Success")
	return nil
}
