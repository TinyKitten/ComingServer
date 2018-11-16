package utils

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv .envを読み込み
func LoadEnv() {
	err := godotenv.Load()
	if err != nil && flag.Lookup("test.v") == nil {
		log.Print("Error loading .env file")
	}
}

// GetPort ポート番号を返す
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return ":8080"
	}
	return fmt.Sprintf(":%s", port)
}

// GetEnv 動作環境フラグを返す
func GetEnv() string {
	env := os.Getenv("ENV")
	if env == "" {
		return "development"
	}
	return env
}
