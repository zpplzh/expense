package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/zappel/expense-server/internal/catalog/model"
)

var (
	ErrNotFound = "sessionid not valid"
)

type Aut interface {
	Validateus(next http.Handler) http.Handler
}

type servicedb2 struct {
	db *sqlx.DB
}

func NewServicess(db *sqlx.DB) Aut {
	return &servicedb2{
		db: db,
	}
}

func (r *servicedb2) Validateus(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ro := req.Header.Get("sessionid")
		ctxz := context.Background()

		exists, err1 := model.Sessions(qm.Where("sessionid=?", ro)).Exists(ctxz, r.db)
		if err1 != nil || exists == false {
			//t := ErrNotFound
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": ErrNotFound,
			})
			return

		}

		next.ServeHTTP(w, req)
	})
}
