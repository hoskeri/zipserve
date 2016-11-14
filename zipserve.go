package zipserve

import "net/http"
import "archive/zip"
import "os"
import "log"

type ZipHttpFile struct {
	z  *zip.File
	fi os.FileInfo
	http.File
}

func NewZipHttpFile(z *zip.File) ZipHttpFile {
	return ZipHttpFile{z: z, fi.FileInfo.}
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

func (z ZipFileSystem) Open(path string) (f http.File, err error) {
	entry, ok := z.entries[path]
	if !ok {
		return nil, os.ErrNotExist
	}
	return entry, nil
}

type entryMap map[string]ZipHttpFile

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
		z.entries[f.Name] = NewZipHttpFile(f)
	}

	return
}
