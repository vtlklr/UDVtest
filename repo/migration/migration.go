package migration

import (
	"database/sql"
	"fmt"
	"github.com/pressly/goose"
)

func Run(db *sql.DB) {
	runGooseMigration(db)
}

func runGooseMigration(db *sql.DB) {
	goose.SetTableName("goose_migration")
	goose.SetVerbose(false)
	if err := goose.Up(db, "."); err != nil {
		fmt.Println(err)
	}
}
