package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"github.com/unrolled/secure"

	_ "github.com/lib/pq"
	configs "github.com/zappel/expense-server/cmd/expense/config"
	"github.com/zappel/expense-server/internal/catalog"
	httptransport "github.com/zappel/expense-server/internal/catalog/http"
	"github.com/zappel/expense-server/internal/catalog/middleware"
)

func main() {

	secureMiddleware := secure.New(secure.Options{
		FrameDeny:               true,
		CustomFrameOptionsValue: "SAMEORIGIN",
		ContentTypeNosniff:      true,
	})

	conf := configs.GetEnv()

	db, err := sqlx.Open("postgres", "dbname="+conf.DBName+" user="+conf.DBUsername+" password="+conf.DBPass+" sslmode=disable")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	svc := catalog.NewServices(db) //func
	svcm := middleware.NewServicess(db)

	r := chi.NewRouter()
	r.Use(secureMiddleware.Handler)

	r.Group(func(c chi.Router) {
		c.Use(svcm.ValidateUser)

		//category
		c.Post("/addcategory", httptransport.AddCategory(svc).ServeHTTP) // svc itu kirim receiver , addcategory(svc) return -> servehttp
		c.Get("/getcategory/{category}", httptransport.GetCategory(svc))
		c.Delete("/deletecategory/{category}", httptransport.DelCategory(svc).ServeHTTP)
		c.Get("/listcategories", httptransport.ListCategories(svc).ServeHTTP)

		//expenses
		c.Post("/addexpense", httptransport.AddExpense(svc).ServeHTTP)
		c.Get("/listexpenses", httptransport.ListExpenses(svc).ServeHTTP)
		c.Get("/getexpense/{expense}", httptransport.GetExpense(svc).ServeHTTP)

		//user
		c.Post("/logout", httptransport.Logout(svc).ServeHTTP)
	})

	//user
	r.Group(func(d chi.Router) {
		d.Post("/signup", httptransport.SignUp(svc).ServeHTTP)
		d.Post("/login", httptransport.Login(svc).ServeHTTP)

	})

	log.Println("Listening on ", conf.PortServer, "...")
	http.ListenAndServe(conf.PortServer, r)
}
