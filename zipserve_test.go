package zipserve

import "testing"
import "net/http/httptest"
import "bytes"

func TestZipRead(t *testing.T) {
	z, err := New("testdata/test.jar")
	if err != nil {
		t.Fatalf("failed to initialize %s", err)
	}

	req := httptest.NewRequest("GET", "/a.txt", nil)
	w := httptest.NewRecorder()

	z.ServeHTTP(w, req)
	if w.Body != bytes.NewBuffer([]byte("aaaa")) {
		t.Log(w.Body)
		t.Fail()
	}
}
