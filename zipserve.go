package zipserve

import "net/http"
import "archive/zip"
import "os"

type ZipHttpFile struct {
	z *zip.File
	http.File
}

func NewZipHttpFile(z *zip.File) ZipHttpFile {
	return ZipHttpFile{z: z}
}

func (z ZipHttpFile) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (z ZipHttpFile) Close() error {
	return nil
}

func (z ZipHttpFile) Seek(offset int64, whence int) (ret int64, err error) {
	return
}

func (z ZipHttpFile) ReadDir(count int) (entries []os.FileInfo, err error) {
	return
}

func (z ZipHttpFile) Stat() (fi os.FileInfo, err error) {
	return
}

type entryMap map[string]ZipHttpFile

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
		z.entries[f.Name] = NewZipHttpFile(f)
	}
	return
}

func (z *ZipServer) Open(path string) (f http.File, err error) {
	return z.entries[path], nil
}

func (z *ZipServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	return
}
