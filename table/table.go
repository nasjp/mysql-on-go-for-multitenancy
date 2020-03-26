package table

import (
	"github.com/jmoiron/sqlx"
	"github.com/nasjp/mysql-on-go-for-multitenancy/table/sqldrivers"
)

func Run(db *sqlx.DB) error {
	for i := 1; i <= 5; i++ {
		if err := sqldrivers.DropUserTable(db, i); err != nil {
			return err
		}
		if err := sqldrivers.CreateUserTable(db, i); err != nil {
			return err
		}
		if err := sqldrivers.InsertUsers(db, i); err != nil {
			return err
		}
	}
	return nil
}
