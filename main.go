package main

import (
	"log"
	"strings"
	"time"

	"github.com/pattack/fcjanstun/team"
	"github.com/pattack/rcss"
)

var (
	Application string = "fcjanstun"

	cfgDir      string = "/etc/" + Application
	cfgFilename string = cfgDir + "/conf.toml"
)

func init() {
	log.SetFlags(log.Lshortfile)
	log.SetPrefix("FCJanstun :: " + strings.ToUpper(Application) + " :: ")
}

func main() {
	cfg := LoadConfigs(defaultConfig(Application, cfgFilename))

	log.Printf("server address: %s\n", cfg.serverAddress)
	log.Printf("initalizing `%s` team\n", cfg.teamName)

	if srv, err := rcss.NewServer(cfg.serverAddress); err != nil {
		panic(err)
	} else if err := srv.Join(team.NewTeam(cfg.teamName)); err != nil {
		panic(err)
	}

	// TODO: Remove this. Wait for stop signal and do graceful shutdown instead
	<-time.After(time.Hour)
}
