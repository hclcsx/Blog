package config

import (
	"github.com/go-clog/clog"
	"github.com/go-ini/ini"
	"os"
	"path/filepath"
)

type Scheme string

const (
	DEFAULT_INI               = "config.ini"
	SCHEME_HTTP        Scheme = "http"
	SCHEME_HTTPS       Scheme = "https"
	SCHEME_UNIX_SOCKET Scheme = "unix"
)

var (
	// Protocol setting
	ProtPort string
	ProtAddr string

	PageSize int
)

func init() {
	clog.New(clog.CONSOLE, clog.ConsoleConfig{})
}

// NewContext should be called by init of main.go to load all config information
func NewContext() {
	appRoot, _ := os.Getwd()
	appIni := filepath.Join(appRoot, "config", DEFAULT_INI)

	cfg, err := ini.Load(appIni)
	if err != nil {
		clog.Fatal(2, "Failed to %s : %v.", appIni, err)
	}

	sec := cfg.Section("server")
	ProtPort = sec.Key("HTTP_PORT").MustString("8080")
	ProtAddr = sec.Key("HTTP_ADDR").MustString("")

	sec = cfg.Section("app")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
