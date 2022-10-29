package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "8081"

type Config struct {
}

func main() {
	app := Config{}

	log.Printf("starting broker service on port %s\n", webPort)

	//define http serve

	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	//start serve
	err := srv.ListenAndServe()

	if err != nil {

	}
}
