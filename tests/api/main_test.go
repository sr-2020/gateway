package api

import (
	"github.com/sr-2020/gateway/config"
	"os"
	"testing"
)

var (
	cfg config.Config
)


func TestMain(m *testing.M) {
	cfg = config.LoadConfig()

	exitVal := m.Run()

	os.Exit(exitVal)
}
