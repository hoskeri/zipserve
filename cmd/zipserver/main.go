package main

import "log"
import "flag"
import "os"
import "net/http"
import "github.com/hoskeri/zipserve"

func main() {
	var zipfile string
	var listenAddr string

	flag.StringVar(&zipfile, "zip", "", "single zip file")
	flag.StringVar(&listenAddr, "addr", ":8080", "http listen address")
	flag.Parse()

	if zipfile == "" {
		flag.Usage()
		os.Exit(1)
	}

	z, err := zipserve.New(zipfile)
	if err != nil {
		log.Fatalf("failed to open zip file %s", zipfile)
	}

	err = http.ListenAndServe(listenAddr, http.FileServer(z))
	if err != nil {
		log.Fatalf("failed to start http server %s", err)
	}
}
