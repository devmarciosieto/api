package database

import (
	"github.com/devmarciosieto/api/internal/domain/campaign"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func NewDB() *gorm.DB {
	dsn := os.Getenv("DATABASE")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&campaign.Campaign{}, &campaign.Contact{})

	return db

}
