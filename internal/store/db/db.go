package db

import (
	"goth/internal/config"
	"goth/internal/store"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func open(cfg *config.Config) (*gorm.DB, error) {
	dsn := "host=" + cfg.DatabaseHost + " user=" + cfg.DatabaseUser + " password=" + cfg.DatabasePassword + " dbname=" + cfg.DatabaseName + " port=" + cfg.DatabasePort + " sslmode=" + cfg.DatabaseSSLMODE + " TimeZone=UTC"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func MustOpen(cfg *config.Config) *gorm.DB {
	db, err := open(cfg)
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&store.User{}, &store.Session{})

	if err != nil {
		panic(err)
	}

	return db
}
