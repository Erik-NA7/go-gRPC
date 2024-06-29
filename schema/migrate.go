package schema

import (
	"database/sql"

	"github.com/GuiaBolso/darwin"
)

var migrations = []darwin.Migration{
	{
		Version: 1,
		Description: "Create table cities",
		Script: `
			CREATE TABLE IF NOT EXISTS cities (
				id SERIAL PRIMARY KEY,
				name VARCHAR(15) NOT NULL
			);
		`,
	},
}

func Migrate(db *sql.DB) error {
	driver := darwin.NewGenericDriver(db, darwin.PostgresDialect{})
	d := darwin.New(driver, migrations, nil)
	return d.Migrate()
}