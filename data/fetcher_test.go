package data

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHTTPFetcher_Fetch(t *testing.T) {

	tmpDir := mktempdir()
	defer func() {
		os.Remove(tmpDir)
	}()

	msg := "hello, world"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.WriteHeader(http.StatusOK)
		w.Write([]byte(msg))
	}))
	defer server.Close()

	src := server.URL
	dest := tmpDir + "/hello.txt"

	fetcher := HTTPFetcher{}
	fetcher.Fetch(src, dest)

	contents, _ := ioutil.ReadFile(dest)
	got := string(contents)
	want := msg
	if got != want {
		t.Errorf("file (%s) contents wrong! got '%s' want '%s'", dest, got, want)
	}
}

func mktempdir() string {
	dir, _ := ioutil.TempDir("", "nhsfinder")
	return dir
}
