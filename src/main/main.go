package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"bot"
	"logging"
)

func main() {
	const port = ":80"

	var debug bool
	flag.BoolVar(&debug, "debug", false, "debug mode")
	flag.Parse()

	router := mux.NewRouter()
	router.Path("/").Methods("POST").HandlerFunc(bot.Index)
	router.Path("/").Methods("GET").HandlerFunc(index)

	if err := logging.Init("logs/logs.log"); err != nil {
		log.Fatalf("[ERR] %s", err.Error())
	}

	if err := http.ListenAndServeTLS(port, "ssl/public.pem", "ssl/private.key", router); err != nil {
		log.Fatalf("[ERR] %s", err.Error())
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Bot is working")
}