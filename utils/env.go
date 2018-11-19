package utils

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

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

// GetApproachedThreshold 接近閾値を返す
func GetApproachedThreshold() (float64, error) {
	thresholdStr := os.Getenv("APPROACHED_THRESHOLD")
	if thresholdStr == "" {
		return 1000, nil
	}
	threshold, err := strconv.ParseFloat(thresholdStr, 64)
	if err != nil {
		return 0, err
	}
	return threshold, nil
}

// GetLeaveThreshold 出発閾値を返す
func GetLeaveThreshold() (float64, error) {
	thresholdStr := os.Getenv("LEAVE_THRESHOLD")
	if thresholdStr == "" {
		return 1500, nil
	}
	threshold, err := strconv.ParseFloat(thresholdStr, 64)
	if err != nil {
		return 0, err
	}
	return threshold, nil
}
