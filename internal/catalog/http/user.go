package http

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/zappel/expense-server/internal/catalog"
	"github.com/zappel/expense-server/internal/catalog/endpoint"
)

func SignUp(sgnup catalog.Service) http.Handler {
	return httptransport.NewServer(
		//endpoint
		endpoint.SignUp(sgnup),

		//decoder
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var inp catalog.SignUpInput

			if err := json.NewDecoder(r.Body).Decode(&inp); err != nil {

			}
			return &inp, nil
		},

		//encoder
		encodeResponse,
	)
}
