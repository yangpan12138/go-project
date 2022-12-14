package main

import (
	"flag"
	"go-interface/api"
	"go-interface/config"

	"github.com/rs/zerolog/log"
)

func main() {

	var (
		configYaml = flag.String("config", "", "")
	)
	flag.Parse()

	conf := config.NewConfig("FILE_SYSTEM")
	if *configYaml != "" {
		log.Info().Str("config", *configYaml).Msg("read config")
		if err := conf.ReadConfigFile(*configYaml); err != nil {
			log.Fatal().Err(err).Send()
		}
	}

	s, err := api.NewServer(conf)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	if err := s.Run(); err != nil {
		log.Fatal().Err(err).Send()
	}

}
