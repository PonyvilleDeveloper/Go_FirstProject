package main

import (
	"app/config"
	"app/storageSQL"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

var cfg *config.Config

func main() {
	args := os.Args
	fmt.Println(args)

	//port := flag.String("port", "8080", "service port")

	flag.Parse()

	cfg := config.Get()

	m := http.NewServeMux()

	m.Handle("/", http.HandlerFunc(mapping))

	server := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      m,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	storageSQL.Migrate()
	server.ListenAndServe()
}
