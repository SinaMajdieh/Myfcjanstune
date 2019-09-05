package main

import (
	"fmt"

	"gopkg.in/fzerorubigd/onion.v3"
	_ "gopkg.in/fzerorubigd/onion.v3/tomlloader"
)

type Config struct {
	serverAddress string
	teamName      string
}

const (
	DEFAULT_SERVER_HOST string = "127.0.0.1"
	DEFAULT_SERVER_PORT int    = 6000
	DEFAULT_TEAM_NAME   string = "Janstun"
)

const (
	ConfServerHost string = "server.host"
	ConfServerPort string = "server.port"

	ConfTeamName string = "team.name"
)

func LoadConfigs(o *onion.Onion) Config {
	cfg := Config{
		teamName:      o.GetString(ConfTeamName),
		serverAddress: fmt.Sprintf("%s:%d", o.GetString(ConfServerHost), o.GetInt(ConfServerPort)),
	}

	return cfg
}

func defaultConfig(name, filename string) *onion.Onion {
	o := onion.New()

	dl := onion.NewDefaultLayer()
	dl.SetDefault(ConfServerHost, DEFAULT_SERVER_HOST)
	dl.SetDefault(ConfServerPort, DEFAULT_SERVER_PORT)
	dl.SetDefault(ConfTeamName, DEFAULT_TEAM_NAME)

	o.AddLayer(dl)
	o.AddLayer(onion.NewFileLayer(filename))

	return o
}
