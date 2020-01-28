package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

type mockDownloader struct {
	fromFile string
}

func (m mockDownloader) Download(src, destFile string) error {
	df, err := os.Create(destFile)
	if err != nil {
		fmt.Println("err", err)
	}
	defer df.Close()

	sf, err := os.Open(m.fromFile)
	if err != nil {
		fmt.Println("err", err)
	}
	defer sf.Close()

	io.Copy(df, sf)
	return nil
}

func TestDownloadPharmacy(t *testing.T) {
	destDir := tempDir()
	defer func() { os.RemoveAll(destDir) }()

	downloader := mockDownloader{fromFile: "../../../testdata/Pharmacy.csv"}
	downloadPharmacy(downloader, destDir)

	want := "pharmacies.csv"
	info, err := os.Stat(path.Join(destDir, want))
	if err != nil || info.Name() != want {
		t.Errorf("wanted file not created: %v!", err)
	}
}

func tempDir() string {
	dir, _ := ioutil.TempDir("", "findertest-")
	return dir
}
