package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type PostBody struct {
	URL string `json:url`
}

type Response struct {
	Error string `json:"error, omitempty"`
	Data  any    `json:"data, omitempty"`
}

func SendJson(w http.ResponseWriter, resp Response, status int) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("Failed to marshal json data", "error", err)
		SendJson(
			w,
			Response{Error: "Something Went Wrong"},
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("Failed To Write Response To Client", "error", err)
		return
	}
}
