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
	//path, httphandlerfunc.ServeHTTP
	r.Post("/addcategory", httptransport.AddCategory(svc).ServeHTTP) //kirim svc itu kirim receiver , addcategory(svc) return baru jalanin servehttp
	r.Get("/getcategory/{category}", httptransport.GetCategory(svc))
	r.Post("/deletecategory/{category}", httptransport.DelCategory(svc).ServeHTTP)
	r.Get("/listcategories", httptransport.ListCategories(svc).ServeHTTP)
	r.Post("/addexpense", httptransport.AddExpense(svc).ServeHTTP)
	r.Get("/listexpenses", httptransport.ListExpenses(svc).ServeHTTP)

	log.Println("Listening on :8080...")
	http.ListenAndServe(":8080", r)
}
