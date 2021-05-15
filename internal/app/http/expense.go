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

func AddExpense(addex app.Service) http.Handler { //interface getcat addcat
	// Endpoint.
	return httptransport.NewServer(
		//endpoint
		endpoint.AddExpense(addex),

		// Decoder.
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var input app.AddExpenseInput
			if err := json.NewDecoder(r.Body).Decode(&input); err != nil {

			}

			return &input, nil
		},

		// Encoder.
		encodeResponse,
		httptransport.ServerErrorEncoder(encodeError),
	)
}

func ListExpenses(showallex app.Service) http.Handler { //interface getcat addcat
	// Endpoint.
	return httptransport.NewServer(
		//endpoint
		endpoint.ListExpenses(showallex),

		// Decoder.
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var qry app.ListExpensesInput

			return &qry, nil
		},

		// Encoder.
		encodeResponse,
		httptransport.ServerErrorEncoder(encodeError),
	)
}

func GetExpense(getex app.Service) http.Handler { //interface getcat addcat

	return httptransport.NewServer(
		//endpoint
		endpoint.GetExpense(getex),

		// Decoder.
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var qry app.GetExpenseInput
			qry.Id = chi.URLParam(r, "id")
			return &qry, nil
		},

		// Encoder.
		encodeResponse,
		httptransport.ServerErrorEncoder(encodeError),
	)
}

func DelExpense(Delex app.Service) http.Handler { //interface getcat addcat
	// Endpoint.
	return httptransport.NewServer(
		//endpoint
		endpoint.DelExpense(Delex),

		// Decoder.
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var qry app.DelExpenseInput
			qry.Id = chi.URLParam(r, "id")
			return &qry, nil
		},

		// Encoder.
		encodeResponse,

		httptransport.ServerErrorEncoder(encodeError),
	)
}

func UpdateExpense(upex app.Service) http.Handler {
	//app.Service itu function nya tapi dia terima receiver svc yang di main karna di app minta receiver
	return httptransport.NewServer(
		// Endpoint.
		endpoint.UpdateExpense(upex),

		// Decoder.
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var input app.UpdateExpenseInput
			input.Id = chi.URLParam(r, "id")
			if err := json.NewDecoder(r.Body).Decode(&input); err != nil {

			}
			return &input, nil
		},

		// Encoder.
		encodeResponse,

		httptransport.ServerErrorEncoder(encodeError),
	)
}

func AddExpenseBatch(addex app.Service) http.Handler {
	// Endpoint.
	return httptransport.NewServer(
		//endpoint
		endpoint.AddExpenseBatch(addex),

		// Decoder.
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var input app.AddExpenseBatchInput
			if err := json.NewDecoder(r.Body).Decode(&input); err != nil {

			}

			return &input, nil
		},

		// Encoder.
		encodeResponse,
		httptransport.ServerErrorEncoder(encodeError),
	)
}
