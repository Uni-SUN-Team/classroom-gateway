package main

import (
	"log"
	"unisun/api/classroom-gateway/src"
	"unisun/api/classroom-gateway/src/config/environment"
	"unisun/api/classroom-gateway/src/constants"
)

func main() {
	err := environment.LoadEnvironment()
	if err != nil {
		log.Fatal(err)
	}
	r := src.App()
	port := environment.ENV.App.Port
	if port == "" {
		r.Run(":" + constants.PORT)
	} else {
		r.Run(":" + port)
	}
}
