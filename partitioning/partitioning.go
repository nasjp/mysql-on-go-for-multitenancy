package partitioning

import (
	"github.com/jmoiron/sqlx"
	"github.com/nasjp/mysql-on-go-for-multitenancy/partitioning/sqldrivers"
)

func Run(db *sqlx.DB) error {
	if err := sqldrivers.DropUserTable(db); err != nil {
		return err
	}

	if err := sqldrivers.CreateUserTable(db); err != nil {
		return err
	}

	if err := sqldrivers.AlterUserTable(db); err != nil {
		return err
	}

	if err := sqldrivers.InsertUsers(db); err != nil {
		return err
	}
	return nil
}
