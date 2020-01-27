package pharmacyfinder

import (
	"io"
	"net/http"
	"os"
)

// Downloader downloads file from src and saves to destFile.
type Downloader interface {
	Download(src string, destFile string) error
}

// HTTPDownloader downloads a file from a URL.
type HTTPDownloader struct {
}

// Download file from src and save to destFile.
func (h *HTTPDownloader) Download(src string, destFile string) error {
	resp, err := http.Get(src)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(destFile)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
