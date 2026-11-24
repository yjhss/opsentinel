package api

import (
	"encoding/json"
	"net/http"

	"github.com/yjhss/opsentinel/internal/job"
)

type Handler struct {
	version string
}

func NewHandler(version string) http.Handler {
	h := &Handler{version: version}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{
			"status": "alive",
		})
	})

	mux.HandleFunc("GET /readyz", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{
			"status": "ready",
		})
	})

	mux.HandleFunc("GET /version", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{
			"version": h.version,
		})
	})

	mux.HandleFunc("GET /api/jobs", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, job.SampleJobs(h.version))
	})

	return mux
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		// 헤더가 이미 전송된 후 발생한 오류이므로
		// 여기에서는 응답 상태 코드를 다시 변경하지 않습니다.
		return
	}
}
