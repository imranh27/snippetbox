package main

import (
	"bytes"
	"net/http"
	"net/url"
	"testing"
)

func TestPing(t *testing.T) {

	//Create new application struct instance with mock loggers
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/ping")

	if code != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, code)
	}

	if string(body) != "OK" {
		t.Errorf("want body to equal %q", "OK")
	}
}

func TestShowSnippet(t *testing.T) {
	app := newTestApplication(t)

	//New test server
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	//setup tests
	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody []byte
	}{
		{"Valid ID", "/snippet/1", http.StatusOK, []byte("An old silent pond...")},
		{"Non-existent ID", "/snippet/2", http.StatusNotFound, nil},
		{"Negative ID", "/snippet/-1", http.StatusNotFound, nil},
		{"Decimal ID", "/snippet/1.23", http.StatusNotFound, nil},
		{"String ID", "/snippet/foo", http.StatusNotFound, nil},
		{"Empty ID", "/snippet/", http.StatusNotFound, nil},
		{"Trailing Slash", "/snippet/1/", http.StatusNotFound, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.get(t, tt.urlPath)

			if code != tt.wantCode {
				t.Errorf("want %d; got %d", tt.wantCode, code)
			}

			if !bytes.Contains(body, tt.wantBody) {
				t.Errorf("want body to contain %q", tt.wantBody)
			}
		})
	}
}

func TestSignupUser(t *testing.T) {
	//Create application struct with mocked dependencies.
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	_, _, body := ts.get(t, "/user/signup")
	csrfToken := extractCSRFToken(t, body)

	validUserName := "Bob"
	validUserEmail := "bob@example.com"
	validUserPassword := "validPa$$word"
	blankFieldMessage := "Cannot be blank"
	//invalidFieldMessage := "This field is invalid"

	tests := []struct {
		name         string
		userName     string
		userEmail    string
		userPassword string
		csrfToken    string
		wantCode     int
		wantBody     []byte
	}{
		{"Valid Submission", validUserName, validUserEmail, validUserPassword, csrfToken, http.StatusSeeOther, nil},
		{"Empty Name", "", validUserEmail, validUserPassword, csrfToken, http.StatusOK, []byte(blankFieldMessage)},
		{"Empty Email", validUserName, "", validUserPassword, csrfToken, http.StatusOK, []byte(blankFieldMessage)},
		{"Empty Password", validUserName, validUserEmail, "", csrfToken, http.StatusOK, []byte(blankFieldMessage)},
		//{"Invalid email (imcomplete domain)", validUserName, "bob@example.", validUserPassword, csrfToken, http.StatusOK, []byte(invalidFieldMessage)},
		//{"Invalid email (missing @)", validUserName, "bobexample.com", validUserPassword, csrfToken, http.StatusOK, []byte(invalidFieldMessage)},
		//{"Invalid email (missing local part)", validUserName, "@example.com", validUserPassword, csrfToken, http.StatusOK, []byte(invalidFieldMessage)},
		{"Short Password", validUserName, validUserEmail, "pa$$word", csrfToken, http.StatusOK, []byte("This field is too short, 10 minimum")},
		{"Duplicate Email", validUserName, "dupe@example.com", validUserPassword, csrfToken, http.StatusOK, []byte("Address is already in use")},
		{"Invalid CSRF Token", "", "", "", "wrongToken", http.StatusBadRequest, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{}
			form.Add("name", tt.userName)
			form.Add("email", tt.userEmail)
			form.Add("password", tt.userPassword)
			form.Add("csrf_token", tt.csrfToken)

			code, _, body := ts.postForm(t, "/user/signup", form)

			if code != tt.wantCode {
				t.Errorf("want %d; got %d", tt.wantCode, code)
			}

			if !bytes.Contains(body, tt.wantBody) {
				t.Errorf("want body %s to contain %q", body, tt.wantBody)
			}
		})
	}
}
