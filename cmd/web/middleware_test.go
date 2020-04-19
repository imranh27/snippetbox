package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSecureHeaders(t *testing.T) {
	//initialise test request
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	//create mock handler
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	//Pass the mock handler to our secure headers middleware
	secureHeaders(next).ServeHTTP(rr, r)

	//collect results
	rs := rr.Result()

	//Check middleware
	frameOptionsGot := rs.Header.Get("X-Frame-Options")
	frameOptionsWant := "deny"
	if frameOptionsGot != frameOptionsWant {
		t.Errorf("want %q; got %q", frameOptionsWant, frameOptionsGot)
	}

	xssProtectionGot := rs.Header.Get("X-XSS-Protection")
	xssProtectionWant := "1;mode=block"
	if xssProtectionGot != xssProtectionWant {
		t.Errorf("want %q; got %q", xssProtectionWant, xssProtectionGot)
	}

	//test for called next handler
	if rs.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, rs.StatusCode)
	}

	defer rs.Body.Close()
	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "OK" {
		t.Errorf("want body to equal %q", "OK")
	}

}
