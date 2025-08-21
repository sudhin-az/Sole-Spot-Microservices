package db

import (
	"fmt"
	"log"

	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Product-Service/pkg/config"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Product-Service/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("port=%s host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Port, cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if dbErr != nil {
		return nil, dbErr
	}
	err := db.AutoMigrate(&domain.Product{})
	if err != nil {
		return nil, err
	}
	log.Println("Database migrated successfully")
	return db, nil
}
