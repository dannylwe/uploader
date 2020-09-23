package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	// "github.com/danny/services/config"
	// "github.com/danny/services/handlers"
	"github.com/danny/services/model"
)



func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestHealthCheckHandler(t *testing.T) {
    
    req, err := http.NewRequest("GET", "/health-check", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(HealthCheckHandler)
	
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    expected := `{"alive": true}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestRedirectUpload(t *testing.T) {
	
	req:= httptest.NewRequest(http.MethodGet, "/", nil)
	
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RedirectToUpload)
	handler.ServeHTTP(rr, req)
	
	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("handler returned wrong status code: got %v want %v",
		status, http.StatusOK)
	}
}

func TestUploadHandler(t *testing.T) {
	
	req:= httptest.NewRequest(http.MethodGet, "/upload", nil)
	
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UploadHandler)
	handler.ServeHTTP(rr, req)
	
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
		status, http.StatusOK)
	}
}

func TestGetTopFiveProfitableItems_WrongMethod(t *testing.T) {
	
	req:= httptest.NewRequest(http.MethodGet, "/topfive", nil)
	
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetTopFiveProfitableItems)
	handler.ServeHTTP(rr, req)
	
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
		status, http.StatusOK)
	}
}

func TestGetTopFiveProfitableItems(t *testing.T) {

	model.SQLConn()

	var jsonStr = []byte(`{"startDate": "2016-01-09", "endDate": "2016-10-19"}`)
	req:= httptest.NewRequest(http.MethodPost, "/topfive", bytes.NewBuffer(jsonStr))
	
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetTopFiveProfitableItems)
	handler.ServeHTTP(rr, req)
	
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
		status, http.StatusOK)
	}
}

func TestGetProfitsByDate(t *testing.T) {

	model.SQLConn()

	var jsonStr = []byte(`{"startDate": "2016-01-09", "endDate": "2016-10-19"}`)
	req, _:= http.NewRequest(http.MethodPost, "/profit", bytes.NewBuffer(jsonStr))
	
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetTopFiveProfitableItems)
	handler.ServeHTTP(rr, req)
	
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
		status, http.StatusOK)
	}
}