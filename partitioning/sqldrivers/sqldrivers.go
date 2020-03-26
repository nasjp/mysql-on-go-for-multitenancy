package sqldrivers

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/nasjp/mysql-on-go-for-multitenancy/domains"
)

func DropUserTable(db *sqlx.DB) error {
	q := `
DROP TABLE IF EXISTS users;
 `

	if _, err := db.Exec(q); err != nil {
		return err
	}
	return nil
}

func CreateUserTable(db *sqlx.DB) error {
	q := `
CREATE TABLE users (
  id int(11) NOT NULL AUTO_INCREMENT,
  company_id int(11) NOT NULL,
  name varchar(255) COLLATE utf8mb4_bin NOT NULL,
  age int(11) NOT NULL,
  PRIMARY KEY (id, company_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin
PARTITION BY LIST COLUMNS(company_id) (
  PARTITION default_company VALUES IN(0)
);
 `

	if _, err := db.Exec(q); err != nil {
		return err
	}
	return nil
}

func AlterUserTable(db *sqlx.DB) error {
	q1 := `
ALTER TABLE users ADD PARTITION (
  PARTITION demo_bank VALUES IN(1)
);
`
	if _, err := db.Exec(q1); err != nil {
		return err
	}

	q2 := `
ALTER TABLE users ADD PARTITION (
  PARTITION demo_shop VALUES IN(2)
);
`
	if _, err := db.Exec(q2); err != nil {
		return err
	}

	q3 := `

ALTER TABLE users ADD PARTITION (
  PARTITION demo_restaurant VALUES IN(3)
);
`
	if _, err := db.Exec(q3); err != nil {
		return err
	}

	q4 := `
ALTER TABLE users ADD PARTITION (
  PARTITION demo_hotel VALUES IN(4)
);
`
	if _, err := db.Exec(q4); err != nil {
		return err
	}

	q5 := `
ALTER TABLE users ADD PARTITION (
  PARTITION demo_system VALUES IN(5)
);
`
	if _, err := db.Exec(q5); err != nil {
		return err
	}

	return nil
}

func InsertUsers(db *sqlx.DB) error {
	us := domains.NewUserService()
	for i := 1; i <= 5; i++ {
		companyID := us.NextID()

		for j := 0; j < 100; j++ {
			name := us.RandomName()
			age := us.RandomAge()

			q := fmt.Sprintf(`
INSERT INTO users (company_id, name, age)
VALUES (%d, "%s", %d);
`, companyID, name, age)

			if _, err := db.Exec(q); err != nil {
				return err
			}
		}
	}

	return nil
}
