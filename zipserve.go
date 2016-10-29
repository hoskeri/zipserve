package zipserve

import "net/http"
import "archive/zip"
import "fmt"

type entryMap map[string]*zip.File

type ZipServer struct {
	http.Handler
	http.FileSystem
	zip     *zip.ReadCloser
	entries entryMap
}

func New(f string) (z *ZipServer, err error) {
	z = new(ZipServer)
	z.entries = make(entryMap)

	z.zip, err = zip.OpenReader(f)
	if err != nil {
		return nil, err
	}

	for _, f := range z.zip.File {
		z.entries[f.Name] = f
	}
	return
}

func (z *ZipServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("zipfile"))
}
