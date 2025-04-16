package users

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func (s *usersServer) GetUsers(rw http.ResponseWriter, r *http.Request) {
	users, err := s.querier.GetUsers(r.Context(), s.pool)
	if err != nil {
		slog.ErrorContext(r.Context(), "error getting users", "error", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(rw).Encode(users)
	if err != nil {
		slog.ErrorContext(r.Context(), "error encoding users", "error", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}
