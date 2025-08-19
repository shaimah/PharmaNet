package httpserver

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"pharmanet/agent/internal/connectors"
)

type Server struct {
	Addr   string
	Token  string
	Search func(query string, limit int) ([]connectors.StockRecord, error)
}

func (s *Server) Start() error {
	http.HandleFunc("/v1/inventory/search", s.handleSearch)
	log.Println("Agent HTTP server listening on", s.Addr)
	return http.ListenAndServe(s.Addr, nil)
}

func (s *Server) handleSearch(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-Agent-Token") != s.Token {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	query := r.URL.Query().Get("query")
	limitStr := r.URL.Query().Get("limit")
	limit := 100
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = l
		}
	}

	results, err := s.Search(query, limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
