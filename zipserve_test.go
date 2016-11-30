package zipserve

import "testing"
import "net/http/httptest"
import "net/http"
import "bytes"

func TestZipRead(t *testing.T) {
	z, err := New("testdata/test.jar")
	if err != nil {
		t.Fatalf("failed to initialize %s", err)
	}

	server := http.FileServer(z)

	req := httptest.NewRequest("GET", "/a.txt", nil)
	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Log(w.Code)
		t.Fail()
	}

	if !bytes.Equal(w.Body.Bytes(), []byte("aaaa\n")) {
		t.Fail()
	}
}
