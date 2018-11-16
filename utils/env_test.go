package utils_test

import (
	"os"
	"testing"

	"github.com/TeamKitten/API/utils"
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
	if port != ":8080" {
		t.Fatal("port number is invalid")
	}
}

func TestGetEnv(t *testing.T) {
	port := utils.GetEnv()
	if port != "development" {
		t.Fatal("env is invalid")
	}
}
