package models

import (
	"errors"
	"fmt"
	"github.com/niki4/challenge_quick_wallet_api/types"
	"github.com/shopspring/decimal"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	maxConnectRetries = 10
)

func InitStorage() (db *gorm.DB, err error) {
	dbUser := "root"
	dbPass := ""
	dbName := "wallet_db"
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	dbDebug := os.Getenv("DEBUG")

	// To handle time.Time correctly, we included parseTime as a parameter to DSN.
	// To fully support UTF-8 encoding, we changed charset=utf8 to charset=utf8mb4 in DSN.
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)

	for i := 1; i < maxConnectRetries+1; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Println("Error connecting to database: ", err)
		} else {
			log.Println("DB connection established")
			break
		}
		if i == maxConnectRetries {
			err = errors.New("couldn't connect to DB")
			log.Println(err)
			return nil, err
		}
		waitSec := 3 * i
		log.Println("Wait before retry DB connect:", waitSec, "seconds")
		time.Sleep(time.Duration(waitSec) * time.Second)
	}

	// Migrate the schema (automatically create tables, missing columns and missing indexes, if needed)
	if err = db.AutoMigrate(&types.Wallet{}); err != nil {
		err = errors.New("failed DB schema migration for Wallet")
		log.Println(err)
		return nil, err
	}

	// Create some wallets for testing in DB if debug mode is ON
	if dbDebug != "" {
		walletsData := []types.Wallet{
			{Balance: decimal.RequireFromString("100.00")},
			{Balance: decimal.RequireFromString("50.00")},
			{Balance: decimal.RequireFromString("500.00")},
			{Balance: decimal.RequireFromString("1.00")},
			{Balance: decimal.RequireFromString("0.05")},
		}
		db.CreateInBatches(walletsData, 5)
		log.Printf("Test wallets has been created: \n%#v\n", walletsData)
	}
	return db, nil
}

func CreateRepository(db *gorm.DB) Repository {
	return Repository{
		db,
	}
}
