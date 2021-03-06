package models

import (
	"errors"
	"fmt"
	"github.com/niki4/challenge_quick_wallet_api/types"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	maxConnectRetries = 10
	db                *gorm.DB
)

func InitStorage() error {
	dbUser := "root"
	dbPass := ""
	dbName := "wallet_db"
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	var err error

	// To handle time.Time correctly, we included parseTime as a parameter to DSN.
	// To fully support UTF-8 encoding, we changed charset=utf8 to charset=utf8mb4 in DSN.
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)

	for i := 0; i < maxConnectRetries; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Println("Error connecting to database: ", err)
		} else {
			log.Println("DB connection established")
			break
		}
		if i == maxConnectRetries-1 {
			err = errors.New("couldn't connect to DB")
			log.Println(err)
			return err
		}
		time.Sleep(1 * time.Second)
	}

	// Migrate the schema (automatically create tables, missing columns and missing indexes, if needed)
	if err = db.AutoMigrate(&types.Wallet{}); err != nil {
		err = errors.New("failed DB schema migration for Wallet")
		log.Println(err)
		return err
	}

	return nil
}
