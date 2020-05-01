package main

import (
    "database/sql"
    "flag"
    "fmt"
    "github.com/jinzhu/gorm"
    "github.com/jonatascabral/pokedex/backend/api/config"
    "github.com/jonatascabral/pokedex/backend/pkg/database"
    "github.com/jonatascabral/pokedex/backend/pkg/utils"
    "github.com/rubenv/sql-migrate"
    "os"
)

var direction string
func init() {
    flag.StringVar(&direction, "direction", "", "set migration direction")
    flag.Parse()
}

func main() {
    if direction == "" {
        utils.Log.Fatal("Missing -direction flag")
    }
    settings := config.LoadSettings()
    db, err := database.ConnectDatabase(settings.Database)
    utils.PanicOnError(err)

    utils.Log.Info("Running migrations")
    migration := initMigration(db, settings.Database)

    switch direction {
    case "up":
        migration.Run(migrate.Up)
    break
    case "down":
        migration.Run(migrate.Down)
    break
    default:
        utils.Log.Fatal(fmt.Sprintf("Invalid direction %s must be up or down", direction))
    }
}


type Migration interface {
    Run(direction migrate.MigrationDirection)
}

type MySQLMigration struct {
    db     *sql.DB
    path   string
    dbName string
    driver string
}

func initMigration(db *gorm.DB, databaseConfig database.DatabaseConfig) Migration {
    pwd, _ := os.Getwd()
    return &MySQLMigration{
        db:     db.DB(),
        path:   pwd + "/migrations/scripts",
        dbName: databaseConfig.Name,
        driver: "mysql",
    }
}

func (m *MySQLMigration) Run(direction migrate.MigrationDirection) {
    n, err := migrate.Exec(m.db, m.driver, &migrate.FileMigrationSource{
        Dir: m.path,
    }, direction)
    utils.PanicOnError(err)

    utils.Log.Info(fmt.Sprintf("Executed %d migrations", n))
}
