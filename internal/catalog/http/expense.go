package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/zappel/expense-server/internal/catalog"
	"github.com/zappel/expense-server/internal/catalog/endpoint"
)

func AddExpense(addex catalog.Service) http.Handler { //interface getcat addcat
	// Endpoint.
	return httptransport.NewServer(
		//endpoint
		endpoint.AddExpense(addex),

		// Decoder.
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var input catalog.AddExpenseInput
			if err := json.NewDecoder(r.Body).Decode(&input); err != nil {

			}

			return &input, nil
		},

		// Encoder.
		encodeResponse,
		httptransport.ServerErrorEncoder(encodeError),
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
		httptransport.ServerErrorEncoder(encodeError),
	)
}

func GetExpense(getex catalog.Service) http.Handler { //interface getcat addcat

	return httptransport.NewServer(
		//endpoint
		endpoint.GetExpense(getex),

		// Decoder.
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var qry catalog.GetExpenseInput
			qry.Id = chi.URLParam(r, "id")
			return &qry, nil
		},

		// Encoder.
		encodeResponse,
		httptransport.ServerErrorEncoder(encodeError),
	)
}
