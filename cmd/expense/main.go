package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"github.com/sethvargo/go-envconfig"

	_ "github.com/lib/pq"
	"github.com/zappel/expense-server/internal/catalog"
	httptransport "github.com/zappel/expense-server/internal/catalog/http"
)

type MyConfig struct {
	PortServer string `env:"PORT"`
	DBName     string `env:"DBNAME"`
	DBUsername string `env:"DBUSER"`
	DBPass     string `env:"DBPASS"`
}

func main() {

	var conf MyConfig
	ctx := context.Background()

	if err := envconfig.Process(ctx, &conf); err != nil {
		log.Fatal(err)
	}

	db, err := sqlx.Open("postgres", "dbname="+conf.DBName+" user="+conf.DBUsername+" password="+conf.DBPass+" sslmode=disable")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	svc := catalog.NewServices(db) //func

	r := chi.NewRouter()
	//path, httphandlerfunc.ServeHTTP
	r.Post("/addcategory", httptransport.AddCategory(svc).ServeHTTP) // svc itu kirim receiver , addcategory(svc) return -> servehttp
	r.Get("/getcategory/{category}", httptransport.GetCategory(svc))
	r.Post("/deletecategory/{category}", httptransport.DelCategory(svc).ServeHTTP)
	r.Get("/listcategories", httptransport.ListCategories(svc).ServeHTTP)
	r.Post("/addexpense", httptransport.AddExpense(svc).ServeHTTP)
	r.Get("/listexpenses", httptransport.ListExpenses(svc).ServeHTTP)
	r.Get("/getexpense/{expense}", httptransport.GetExpense(svc).ServeHTTP)

	log.Println("Listening on ", conf.PortServer, "...")
	http.ListenAndServe(conf.PortServer, r)
}
