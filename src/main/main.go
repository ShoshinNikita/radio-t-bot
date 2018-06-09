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
	var port string
	var debug bool

	flag.StringVar(&port, "port", ":505", "port for bot")
	flag.BoolVar(&debug, "debug", false, "debug mode")
	flag.Parse()
	if port[0] != ':' {
		port = ":" + port
	}

	router := mux.NewRouter()
	router.Path("/").Methods("POST").HandlerFunc(bot.Index)
	router.Path("/").Methods("GET").HandlerFunc(index)

	if err := logging.Init("logs"); err != nil {
		log.Fatalf("[ERR] %s", err.Error())
	}

	if debug {
		if err := http.ListenAndServe(port, router); err != nil {
			log.Fatalf("[ERR] %s", err.Error())
		}
	} else {
		if err := http.ListenAndServeTLS(port, "public.pem", "private.key", router); err != nil {
			log.Fatalf("[ERR] %s", err.Error())
		}
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Bot is working")
}
