package sqldrivers

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/nasjp/mysql-on-go-for-multitenancy/domains"
)

func DropUserTable(db *sqlx.DB, companyID int) error {
	q := fmt.Sprintf(`
DROP TABLE IF EXISTS %d_users;
 `, companyID)

	if _, err := db.Exec(q); err != nil {
		return err
	}
	return nil
}

func CreateUserTable(db *sqlx.DB, companyID int) error {
	q := fmt.Sprintf(`
CREATE TABLE %d_users (
  id int(11) NOT NULL AUTO_INCREMENT,
  company_id int(11) NOT NULL,
  name varchar(255) COLLATE utf8mb4_bin NOT NULL,
  age int(11) NOT NULL,
  PRIMARY KEY (id, company_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
`, companyID)

	if _, err := db.Exec(q); err != nil {
		return err
	}
	return nil
}

func InsertUsers(db *sqlx.DB, companyID int) error {
	us := domains.NewUserService()
	for j := 0; j < 100; j++ {
		name := us.RandomName()
		age := us.RandomAge()

		q := fmt.Sprintf(`
INSERT INTO %[1]d_users (company_id, name, age)
VALUES (%[1]d, "%s", %d);
`, companyID, name, age)

		if _, err := db.Exec(q); err != nil {
			return err
		}
	}

	return nil
}
