package items

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func (s *itemsServer) GetItems(rw http.ResponseWriter, r *http.Request) {
	items, err := s.querier.GetItems(r.Context(), s.pool)
	if err != nil {
		slog.ErrorContext(r.Context(), "error getting items", "error", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(rw).Encode(items)
	if err != nil {
		slog.ErrorContext(r.Context(), "error encoding items", "error", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}
