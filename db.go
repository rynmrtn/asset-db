package assetdb

import (
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

type DBType string

// const (
// 	Postgres DBType = "postgres"
// 	SQLite   DBType = "sqlite"
// )

// func New(dbType DBType, dsn string) (*DB, error) {
// 	switch dbType {
// 	case Postgres:
// 		return postgresDatabase(dsn)
// 	case SQLite:
// 		return sqliteDatabase(dsn)
// 	default:
// 		panic("Unknown db type")
// 	}
// }

// func postgresDatabase(dsn string) (*DB, error) {
// 	gdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &DB{gdb}, nil
// }

// func sqliteDatabase(dsn string) (*DB, error) {
// 	gdb, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &DB{gdb}, nil
// }
