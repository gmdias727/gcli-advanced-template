// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"basic/pkg/logger"
	"basic/source/migration"
	"basic/source/repository"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func newApp(dbType repository.DBType, conf *viper.Viper, db *gorm.DB, logger2 *logger.Logger) (*migration.Migrate, func(), error) {
	migrate := migration.NewMigrate(db, logger2)
	return migrate, func() {
	}, nil
}

// wire.go:

var MigrateSet = wire.NewSet(migration.NewMigrate)
