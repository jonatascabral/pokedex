package database

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type DatabaseConfig struct {
    Driver        string
    Host          string
    Port          string
    Name          string
    Username      string
    Password      string
    MigrationPath string
}

func ConnectDatabase(config DatabaseConfig) (*gorm.DB, error) {
    connectionString := fmt.Sprintf(
        "%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
        config.Username,
        config.Password,
        config.Host,
        config.Name,
    )
    return gorm.Open(config.Driver, connectionString)
}
