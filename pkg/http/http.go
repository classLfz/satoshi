package http

import (
	"log"
	"net/http"

	"github.com/classlfz/satoshi/cmd/config"
)

func Start(configPath string) {
	const CurrentAPI = "Http Start"
	cfg, _, loadErr := config.Load(configPath)
	if loadErr != nil {
		log.Fatal(CurrentAPI, " config.load \"", configPath, "\" err:\n", loadErr)
		return
	}

	log.Printf("Listen and serve at port %v.", cfg.Satoshi.Http.Port)
	serveErr := http.ListenAndServe(":"+cfg.Satoshi.Http.Port, NewAPIMux())
	if serveErr != nil {
		log.Fatal("ListenAndServe: ", serveErr)
	}
}
