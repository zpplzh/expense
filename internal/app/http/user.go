package http

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/zappel/expense-server/internal/app"
	"github.com/zappel/expense-server/internal/app/endpoint"
)

func SignUp(sgnup app.Service) http.Handler {
	return httptransport.NewServer(
		//endpoint
		endpoint.SignUp(sgnup),

		//decoder
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var inp app.SignUpInput

			if err := json.NewDecoder(r.Body).Decode(&inp); err != nil {

			}
			return &inp, nil
		},

		//encoder
		encodeResponse,
		httptransport.ServerErrorEncoder(encodeError),
	)
}

func Login(logi app.Service) http.Handler {
	return httptransport.NewServer(
		//endpoint
		endpoint.Login(logi),

		//decoder
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var inp app.LoginInput

			if err := json.NewDecoder(r.Body).Decode(&inp); err != nil {

			}
			return &inp, nil
		},

		//encoder
		encodeResponse,
		httptransport.ServerErrorEncoder(encodeError),
	)
}

func Logout(logo app.Service) http.Handler {
	return httptransport.NewServer(
		//endpoint
		endpoint.Logout(logo),

		//decoder
		func(_ context.Context, r *http.Request) (interface{}, error) {
			ro := r.Header.Get("sessionid")
			inp := app.Logoutinput{
				Sessionid: ro,
			}

			if err := json.NewDecoder(r.Body).Decode(&inp); err != nil {

			}
			return &inp, nil
		},

		//encoder
		encodeResponse,
		httptransport.ServerErrorEncoder(encodeError),
	)
}
