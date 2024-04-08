package sqlite

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migration(DB *sql.DB, migration Migrations) {

	// instance, err := sqlite3.WithInstance(DB, &sqlite3.Config{})
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fSrc, err := (&file.File{}).Open(migrationDir)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	currentDir, err := os.Getwd()
	migrationDir := currentDir + "/pkg/db/migrations/sqlite/"
	databasePath := currentDir + "/pkg/db/sqlite/social-network.db"
	m, err := migrate.New("file://"+migrationDir, "sqlite://"+databasePath)
	if err != nil {
		// fmt.Println(databasePath, m, err)
		log.Fatal(err)
	}

	// m, err := migrate.NewWithInstance("file", fSrc, "sqlite3", instance)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	defer m.Close()

	switch strings.ToLower(migration.Action) {
	case "-up":
		// Apply one migration (1 Up)
		// if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		// 	fmt.Println(err)
		// }
		if err := m.Steps(1); err != nil && err != migrate.ErrNoChange {
			fmt.Println("Migration Error: ", err)
		}
	case "-down":
		// Rollback one migration (1 Down)
		if err := m.Steps(-1); err != nil && err != migrate.ErrNoChange {
			fmt.Println("Migration Error: ", err)
		}
	case "-up--all":
		// Apply migrations (Up)
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			fmt.Println("Migration Error: ", err)
		}
	case "-down--all":
		// Rollback all migration (Down)
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			fmt.Println("Migration Error: ", err)
		}

		case "-to":
			// Migrate directly to the target version
			if err := m.Migrate(uint(migration.Version)); err != nil {
				fmt.Println("Migration Error: ", err)
			}
	}
	currentVersion, dirty, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		log.Fatal(err)
	}
	fmt.Println("Current database version:", currentVersion)
	fmt.Println("Dirty state of Database: ", dirty)
}
