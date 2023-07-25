package app

import (
	"log"
	"netWorkService/internal/config"
	"netWorkService/internal/transport"
)

func Run() {

	cfg, err := config.NewConfig("./configs/rest.yml")
	if err != nil {
		log.Fatal(err)
	}

	transport.Run(cfg)

}
