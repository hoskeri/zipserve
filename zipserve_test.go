package zipserve

import "testing"
import "net/http/httptest"
import "bytes"
import "net/http"

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

	if w.Body != bytes.NewBuffer([]byte("aaaa")) {
		t.Log(w.Body)
		t.Fail()
	}
}
