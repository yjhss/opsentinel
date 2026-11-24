package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/yjhss/opsentinel/internal/job"
)

func request(method, path string) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(method, path, nil)

	NewHandler("v0.1.0").ServeHTTP(rec, req)
	return rec
}

func TestHealth(t *testing.T) {
	rec := request(http.MethodGet, "/healthz")

	if rec.Code != http.StatusOK {
		t.Fatalf("예상 HTTP 200, 실제 %d", rec.Code)
	}
}

func TestReady(t *testing.T) {
	rec := request(http.MethodGet, "/readyz")

	if rec.Code != http.StatusOK {
		t.Fatalf("예상 HTTP 200, 실제 %d", rec.Code)
	}
}

func TestVersion(t *testing.T) {
	rec := request(http.MethodGet, "/version")

	if rec.Code != http.StatusOK {
		t.Fatalf("예상 HTTP 200, 실제 %d", rec.Code)
	}

	var response struct {
		Version string `json:"version"`
	}

	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	if response.Version != "v0.1.0" {
		t.Fatalf("버전 불일치: %s", response.Version)
	}
}

func TestJobs(t *testing.T) {
	rec := request(http.MethodGet, "/api/jobs")

	if rec.Code != http.StatusOK {
		t.Fatalf("예상 HTTP 200, 실제 %d", rec.Code)
	}

	var jobs []job.Job

	if err := json.Unmarshal(rec.Body.Bytes(), &jobs); err != nil {
		t.Fatal(err)
	}

	if len(jobs) != 3 {
		t.Fatalf("예상 작업 3건, 실제 %d건", len(jobs))
	}

	for _, j := range jobs {
		if j.ReleaseVersion != "v0.1.0" {
			t.Fatalf("작업 버전 불일치: %s", j.ReleaseVersion)
		}
	}
}

func TestMethodNotAllowed(t *testing.T) {
	rec := request(http.MethodPost, "/healthz")

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("예상 HTTP 405, 실제 %d", rec.Code)
	}
}

func TestNotFound(t *testing.T) {
	rec := request(http.MethodGet, "/unknown")

	if rec.Code != http.StatusNotFound {
		t.Fatalf("예상 HTTP 404, 실제 %d", rec.Code)
	}
}
