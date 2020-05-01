package database

import (
    "github.com/jinzhu/gorm"
)

type Migration struct {
    db *gorm.DB
}

func InitMigration(db *gorm.DB) *Migration {
    return &Migration{
        db: db,
    }
}

func (m *Migration) Run() {
    //
}
