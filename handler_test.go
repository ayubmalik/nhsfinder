package nhsfinder

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubFinderFunc func(string) []FindResult

func (sff StubFinderFunc) ByPostcode(postcode string) []FindResult {
	return sff(postcode)
}

func TestPharmacies(t *testing.T) {

	t.Run("returns valid pharmacies", func(t *testing.T) {

		finder := StubFinderFunc(func(postcode string) []FindResult {
			return []FindResult{FindResult{Distance: 1.0, Pharmacy: Pharmacy{Name: "pharmacy1"}}}
		})

		handler := NewPharmacyHandler(finder)
		request, _ := http.NewRequest(http.MethodGet, "/pharmacies/postcode/m44bf", nil)
		response := httptest.NewRecorder()

		handler.ServeHTTP(response, request)

		gotStatus := response.Code
		wantStatus := http.StatusOK
		if gotStatus != wantStatus {
			t.Errorf("did not get correct status, got %d, want %d", gotStatus, wantStatus)
		}

		var got []FindResult
		json.NewDecoder(response.Body).Decode(&got)

		if len(got) != 1 {
			t.Errorf("did not find 2 pharmacies!")
		}

		if got[0].Pharmacy.Name != "pharmacy1" {
			t.Errorf("did not get valid json pharmacy")
		}
	})
}
