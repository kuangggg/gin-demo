package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File
	HttpPort int
	RunMode string
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	PageSize int
)

func init() {

	var err error

	Cfg, err = ini.Load("conf/app.ini")

	if err != nil {
		log.Fatalf("%v", err)
	}
	loadBase()
	loadServer()
	loadApp()
}

func loadBase()  {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func loadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("%v", err)
	}

	HttpPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func loadApp()  {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("%v", err)
	}

	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}