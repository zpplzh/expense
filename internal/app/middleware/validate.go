package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/zappel/expense-server/internal/app/model"
)

var (
	ErrNotFound = "sessionid not valid"
)

type Aut interface {
	ValidateUser(next http.Handler) http.Handler
}

type servicedb2 struct {
	db *sqlx.DB
}

func NewServicess(db *sqlx.DB) Aut {
	return &servicedb2{
		db: db,
	}
}

type outps struct {
	sesid string
}

func (r *servicedb2) ValidateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ro := req.Header.Get("sessionid")
		ctxz := context.Background()
		t := time.Now()
		ts := t.Format(time.RFC3339)

		gusid, err1 := model.Sessions(qm.Where("sessionid = ? and expiry > ? ", ro, ts)).One(ctxz, r.db)
		if err1 != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": ErrNotFound,
			})
			return
		}

		ctxs := context.WithValue(req.Context(), "uid", gusid.UserID)

		next.ServeHTTP(w, req.WithContext(ctxs))
	})
}
