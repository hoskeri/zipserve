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
	var single string

	flag.StringVar(&zs.spec, "spec", "", "list of zip file paths")
	flag.StringVar(&single, "zip", "", "single zip file")
	flag.Parse()

	if zs.spec == "" && single == "" {
		flag.Usage()
		os.Exit(1)
	}

	if single != "" {
		zs.zipfiles = append(zs.zipfiles, single)
	}

	if zs.spec != "" {
		log.Printf("serving from spec: %s\n", zs.spec)
	}

	log.Printf("entries: %s\n", zs.zipfiles)
}
