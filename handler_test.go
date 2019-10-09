package pharmacyfinder

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

func TestPharmacyHandler(t *testing.T) {

	t.Run("pharmacies/postcode delegates to finder valid postcode", func(t *testing.T) {

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

	t.Run("pharmacies/postcode passes postcode path param valid postcode", func(t *testing.T) {

		finder := StubFinderFunc(func(postcode string) []FindResult {
			return []FindResult{FindResult{Pharmacy: Pharmacy{Name: postcode}}}
		})

		handler := NewPharmacyHandler(finder)
		request, _ := http.NewRequest(http.MethodGet, "/pharmacies/postcode/m44bf", nil)
		response := httptest.NewRecorder()

		handler.ServeHTTP(response, request)

		var got []FindResult
		json.NewDecoder(response.Body).Decode(&got)
		postcode := got[0].Pharmacy.Name
		if postcode != "m44bf" {
			t.Errorf("did not pass postcode param %s", postcode)
		}
	})

	t.Run("pharmacies/postcode handles empty result valid postcode", func(t *testing.T) {

		finder := StubFinderFunc(func(postcode string) []FindResult {
			return []FindResult{}
		})

		handler := NewPharmacyHandler(finder)
		request, _ := http.NewRequest(http.MethodGet, "/pharmacies/postcode/m444bf", nil)
		response := httptest.NewRecorder()

		handler.ServeHTTP(response, request)

		var got []FindResult
		json.NewDecoder(response.Body).Decode(&got)
		if len(got) != 0 {
			t.Errorf("did not handle empty result %v", got)
		}
	})

	t.Run("pharmacies/postcode returns bad request invalid postcode", func(t *testing.T) {

		finder := StubFinderFunc(func(postcode string) []FindResult {
			return []FindResult{}
		})

		handler := NewPharmacyHandler(finder)
		request, _ := http.NewRequest(http.MethodGet, "/pharmacies/postcode/1", nil)
		response := httptest.NewRecorder()

		handler.ServeHTTP(response, request)

		if response.Code != 400 {
			t.Errorf("did not get 400 was %v", response.Code)
		}
	})
}
