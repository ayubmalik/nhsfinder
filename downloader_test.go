package pharmacyfinder

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"
)

type mockDownloader struct {
	contents string
}

func (m mockDownloader) Download(src, destDir string) error {
	s := path.Base((src))
	d, _ := os.Create(path.Join(destDir, s))
	defer d.Close()
	io.WriteString(d, m.contents)
	return nil
}

func TestHTTPFetcher_Fetch(t *testing.T) {
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
	dir, _ := ioutil.TempDir("", "pharmacyfinder")
	return dir
}
