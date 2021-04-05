package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/zappel/expense-server/internal/catalog"
	"github.com/zappel/expense-server/internal/catalog/endpoint"
)

func GetCategory(Getcat catalog.Service) http.HandlerFunc { //interface getcat addcat

	return httptransport.NewServer(
		//endpoint
		endpoint.GetCategory(Getcat),

		// Decoder.
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var qry catalog.GetCategoryInput
			qry.Name = chi.URLParam(r, "category")
			return &qry, nil
		},

		// Encoder.
		encodeResponse,
	).ServeHTTP
}

func AddCategory(Addcat catalog.Service) http.Handler {
	//catalog.Service itu function nya tapi dia terima receiver svc yang di main karna di catalog minta receiver

	return httptransport.NewServer(
		// Endpoint.
		endpoint.AddCategory(Addcat),

		// Decoder.
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var input catalog.AddCategoryInput
			if err := json.NewDecoder(r.Body).Decode(&input); err != nil {

			}
			return &input, nil
		},

		// Encoder.
		encodeResponse,
	)
}

func DelCategory(Delcat catalog.Service) http.Handler { //interface getcat addcat
	// Endpoint.
	return httptransport.NewServer(
		//endpoint
		endpoint.DelCategory(Delcat),

		// Decoder.
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var qry catalog.DelCategoryInput
			qry.Name = chi.URLParam(r, "category")
			return &qry, nil
		},

		// Encoder.
		encodeResponse,
	)
}

func ListCategories(showallcat catalog.Service) http.Handler { //interface getcat addcat
	// Endpoint.
	return httptransport.NewServer(
		//endpoint
		endpoint.ListCategories(showallcat),

		// Decoder.
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var qry catalog.ListCategoriesInput

			return &qry, nil
		},

		// Encoder.
		encodeResponse,
	)
}

func AddExpense(addex catalog.Service) http.Handler { //interface getcat addcat
	// Endpoint.
	return httptransport.NewServer(
		//endpoint
		endpoint.AddExpense(addex),

		// Decoder.
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var input catalog.AddExpenseInput
			if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
				fmt.Println(err)

			}

			return &input, nil
		},

		// Encoder.
		encodeResponse,
	)
}

func ListExpenses(showallex catalog.Service) http.Handler { //interface getcat addcat
	// Endpoint.
	return httptransport.NewServer(
		//endpoint
		endpoint.ListExpenses(showallex),

		// Decoder.
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var qry catalog.ListExpensesInput

			return &qry, nil
		},

		// Encoder.
		encodeResponse,
	)
}
