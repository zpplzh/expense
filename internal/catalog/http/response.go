package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/zappel/expense-server/internal/catalog"
)

type Errorer interface {
	error() error
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {

	if e, ok := response.(Errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)

		return nil
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-XSS-Protection", "1;mode=block")
	//w.Header().Set("X-Frame-Options", "deny")

	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {

	case catalog.ErrNotFound:
		return http.StatusNotFound
	case catalog.ErrDuplicate, catalog.DataExistErr:
		return http.StatusConflict

	default:
		return http.StatusInternalServerError
	}
}
