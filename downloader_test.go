package nhsfinder

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"
)

func TestHTTPDownloader(t *testing.T) {
	tmpDir := tempDir()
	defer func() {
		os.RemoveAll(tmpDir)
	}()

	msg := "hello, world"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(msg))
	}))
	defer server.Close()

	src := server.URL
	dest := path.Join(tmpDir, "hello.txt")

	downloader := HTTPDownloader{}
	downloader.Download(src, dest)

	contents, _ := ioutil.ReadFile(dest)
	got := string(contents)
	want := msg
	if got != want {
		t.Errorf("file (%s) contents wrong! got '%s' want '%s'", dest, got, want)
	}
}

func tempDir() string {
	dir, _ := ioutil.TempDir("", "nhsfinder")
	return dir
}
