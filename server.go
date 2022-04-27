package main

import (
	"attractionplace/config"
	"attractionplace/service"
	"log"
	"os"

	servCommon "github.com/aunjaidev/aunjai-common"
	configlib "github.com/aunjaidev/aunjaiconfiglib"
	"github.com/mitchellh/mapstructure"
)

func main() {
	scf := config.Config{}
	if cf, err := configlib.LoadConfig(config.Config{}, "", "resource", os.Getenv("ENV"), "application"); err != nil {
		log.Panic(err)
	} else {
		err = mapstructure.Decode(cf, &scf)
		if err != nil {
			log.Panic("Service Close. -> ", err)
		}
	}

	app := servCommon.CreateServer(scf.Server.ServiceName, scf.Server.Port)

	s := service.NewCreateAttractionPlaceService()
	app.App.POST("/attraction-place/create", s.CreatePlace)
	app.StartServer()
}
