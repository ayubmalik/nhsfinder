package nhsfinder_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ayubmalik/nhsfinder"
)

func TestPharmaciesServer(t *testing.T) {
	finder, _ := nhsfinder.NewInMemFinder()

	handler := nhsfinder.NewPharmacyHandler(finder)

	server := httptest.NewServer(handler)
	defer server.Close()

	searchURL := server.URL + "/pharmacies/postcode/M44BF"
	res, err := http.Get(searchURL)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	got := res.StatusCode
	want := 200
	if got != want {
		t.Fatalf("got status %d wanted %d", got, want)
	}

	decoder := json.NewDecoder(res.Body)
	var result []nhsfinder.FindResult
	decoder.Decode(&result)

	got = len(result)
	want = 10
	if len(result) != 10 {
		t.Fatalf("got %d results want %d", got, want)
	}
}
