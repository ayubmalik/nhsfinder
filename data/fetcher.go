package data

import (
	"io"
	"net/http"
	"os"
)

// Fetcher fetches a file and stores it in location.
type Fetcher interface {
	Fetch(src string, dest string) error
}

// HTTPFetcher is a Fetcher and fetches a file from a Http url.
type HTTPFetcher struct {
}

// Fetch a file and stores it in location.
func (hf *HTTPFetcher) Fetch(src string, dest string) error {
	resp, err := http.Get(src)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
