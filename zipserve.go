package zipserve

import "net/http"
import "archive/zip"
import "os"
import "io"

type zipHttpFile struct {
	z *zip.File
	http.File
	r io.ReadCloser
}

func newZipHttpFile(z *zip.File) zipHttpFile {
	return zipHttpFile{z: z}
}

func (z zipHttpFile) Read(p []byte) (n int, err error) {
	return z.r.Read(p)
}

func (z zipHttpFile) Close() error {
	z.r = nil
	return nil
}

func (z zipHttpFile) Seek(offset int64, whence int) (ret int64, err error) {
	return
}

func (z zipHttpFile) ReadDir(count int) (entries []os.FileInfo, err error) {
	return
}

func (z zipHttpFile) Stat() (fi os.FileInfo, err error) {
	return z.z.FileInfo(), nil
}

func (z ZipFileSystem) Open(path string) (f http.File, err error) {
	entry, ok := z.entries[path]
	if !ok {
		return nil, os.ErrNotExist
	}
	entry.r, err = entry.z.Open()
	return entry, nil
}

type entryMap map[string]zipHttpFile

type ZipFileSystem struct {
	http.FileSystem
	zip     *zip.ReadCloser
	entries entryMap
}

func New(f string) (z *ZipFileSystem, err error) {
	z = new(ZipFileSystem)
	z.entries = make(entryMap)

	z.zip, err = zip.OpenReader(f)
	if err != nil {
		return nil, err
	}

	for _, f := range z.zip.File {
		z.entries["/"+f.Name] = newZipHttpFile(f)
	}

	return
}
