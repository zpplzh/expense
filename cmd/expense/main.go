package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
	"github.com/zappel/expense-server/internal/catalog"
	httptransport "github.com/zappel/expense-server/internal/catalog/http"
)

func main() {

	db, err := sqlx.Open("postgres", "dbname=expense_db user=postgres password=password sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	svc := catalog.NewServices(db) //func

	r := chi.NewRouter()
	r.Post("/AddCategory", httptransport.AddCategory(svc).ServeHTTP) //kirim svc itu kirim receiver
	r.Get("/GetCategory/{category}", httptransport.GetCategory(svc).ServeHTTP)
	r.Post("/DelCategory/{category}", httptransport.DelCategory(svc).ServeHTTP)

	log.Println("Listening on :8080 12...")
	http.ListenAndServe(":8080", r)
}
