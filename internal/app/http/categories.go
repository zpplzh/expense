package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/zappel/expense-server/internal/app"
	"github.com/zappel/expense-server/internal/app/endpoint"
)

func GetCategory(Getcat app.Service) http.HandlerFunc { //interface getcat addcat

	return httptransport.NewServer(
		//endpoint
		endpoint.GetCategory(Getcat),

		// Decoder.
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var qry app.GetCategoryInput
			qry.Name = chi.URLParam(r, "id")
			return &qry, nil
		},

		// Encoder.
		encodeResponse,

		httptransport.ServerErrorEncoder(encodeError),
	).ServeHTTP
}

func AddCategory(Addcat app.Service) http.Handler {
	//app.Service itu function nya tapi dia terima receiver svc yang di main karna di app minta receiver

	return httptransport.NewServer(
		// Endpoint.
		endpoint.AddCategory(Addcat),

		// Decoder.
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var input app.AddCategoryInput
			if err := json.NewDecoder(r.Body).Decode(&input); err != nil {

			}
			return &input, nil
		},

		// Encoder.
		encodeResponse,

		httptransport.ServerErrorEncoder(encodeError),
	)
}

func DelCategory(Delcat app.Service) http.Handler { //interface getcat addcat
	// Endpoint.
	return httptransport.NewServer(
		//endpoint
		endpoint.DelCategory(Delcat),

		// Decoder.
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var qry app.DelCategoryInput
			qry.Categoryid = chi.URLParam(r, "id")
			return &qry, nil
		},

		// Encoder.
		encodeResponse,

		httptransport.ServerErrorEncoder(encodeError),
	)
}

func ListCategories(showallcat app.Service) http.Handler { //interface getcat addcat
	// Endpoint.
	return httptransport.NewServer(
		//endpoint
		endpoint.ListCategories(showallcat),

		// Decoder.
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var qry app.ListCategoriesInput

			return &qry, nil
		},

		// Encoder.
		encodeResponse,

		httptransport.ServerErrorEncoder(encodeError),
	)
}

func UpdateCategory(upcat app.Service) http.Handler {
	//app.Service itu function nya tapi dia terima receiver svc yang di main karna di app minta receiver
	return httptransport.NewServer(
		// Endpoint.
		endpoint.UpdateCategory(upcat),

		// Decoder.
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var input app.UpdateCategoryInput
			input.Categoryid = chi.URLParam(r, "id")
			if err := json.NewDecoder(r.Body).Decode(&input); err != nil {

			}
			return &input, nil
		},

		// Encoder.
		encodeResponse,

		httptransport.ServerErrorEncoder(encodeError),
	)
}
