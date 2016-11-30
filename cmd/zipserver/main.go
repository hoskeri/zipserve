package main

import "log"
import "flag"
import "os"
import "github.com/hoskeri/zipserve"

type zipserver struct {
	spec     string
	zipfiles []string
	entries  []zipserve.ZipFileSystem
}

func main() {
	zs := zipserver{}
	flag.StringVar(&zs.spec, "spec", "", "list of zip file paths")
	flag.Parse()

	if zs.spec == "" {
		flag.Usage()
		os.Exit(1)
	}
	log.Println("starting zipserver")
}
