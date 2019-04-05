package utils_test

import (
	"os"
	"testing"

	"github.com/TinyKitten/ComingServer/utils"
)

func TestMain(m *testing.M) {
	before()
	code := m.Run()
	os.Exit(code)
}

func before() {
	utils.LoadEnv()
}

func TestGetPort(t *testing.T) {
	port := utils.GetPort()
	if port != ":52360" {
		t.Fatal("port number is invalid")
	}
}

func TestGetDSN(t *testing.T) {
	dsn := utils.GetDSN()
	if dsn != "" {
		t.Fatal("default DSN is invalid")
	}
}

func TestGetEnv(t *testing.T) {
	port := utils.GetEnv()
	if port != "development" {
		t.Fatal("env is invalid")
	}
}
