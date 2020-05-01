package config

import (
    "github.com/joho/godotenv"
    "github.com/jonatascabral/pokedex/backend/pkg/database"
    "github.com/jonatascabral/pokedex/backend/pkg/utils"
    "os"
)

type Settings struct {
    Environment string
    AppName     string
    AppPort     string
    Database    database.DatabaseConfig
}

func LoadSettings(envFiles ...string) *Settings {
    err := godotenv.Load(envFiles...)
    utils.PanicOnError(err)

    return &Settings{
        Environment: os.Getenv("APP_ENV"),
        AppName:     os.Getenv("APP_NAME"),
        AppPort:     os.Getenv("APP_PORT"),

        Database: database.DatabaseConfig{
            Driver:        os.Getenv("DATABASE_DRIVER"),
            Host:          os.Getenv("DATABASE_HOST"),
            Port:          os.Getenv("DATABASE_PORT"),
            Name:          os.Getenv("DATABASE_NAME"),
            Username:      os.Getenv("DATABASE_USERNAME"),
            Password:      os.Getenv("DATABASE_PASSWORD"),
            MigrationPath: os.Getenv("DATABASE_MIGRATION_PATH"),
        },
    }
}
