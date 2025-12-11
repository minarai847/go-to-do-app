package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	// .envファイルを読み込む（エラーは無視 - 本番環境では環境変数が設定されているため）
	_ = godotenv.Load()

	// RenderではDATABASE_URLが自動的に提供される
	var url string
	if databaseURL := os.Getenv("DATABASE_URL"); databaseURL != "" {
		url = databaseURL
		fmt.Println("Using DATABASE_URL from environment")
	} else {
		// 個別の環境変数から構築（ローカル開発用）
		postgresUser := os.Getenv("POSTGRES_USER")
		postgresPw := os.Getenv("POSTGRES_PW")
		postgresHost := os.Getenv("POSTGRES_HOST")
		postgresPort := os.Getenv("POSTGRES_PORT")
		postgresDB := os.Getenv("POSTGRES_DB")

		// 環境変数が設定されているかチェック
		if postgresUser == "" || postgresPw == "" || postgresHost == "" || postgresPort == "" || postgresDB == "" {
			log.Fatal("Database connection configuration is missing. Please set DATABASE_URL or individual POSTGRES_* environment variables.")
		}

		url = fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
			postgresUser,
			postgresPw,
			postgresHost,
			postgresPort,
			postgresDB)
		fmt.Println("Using individual POSTGRES_* environment variables")
	}

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	
	// 接続プールの設定
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}
	
	// 接続タイムアウトを設定
	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)
	
	fmt.Println("Connected to database")
	return db
}

func CloseDB(db *gorm.DB) {
	if db == nil {
		return
	}
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatal(err)
	}
}
