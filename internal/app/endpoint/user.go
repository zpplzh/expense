package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/zappel/expense-server/internal/app"
)

func SignUp(siup app.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		input := request.(*app.SignUpInput)
		add, err := siup.SignUp(ctx, input)

		return add, err
	}
}

func Login(logi app.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		input := request.(*app.LoginInput)
		ex, err := logi.Login(ctx, input)

		return ex, err
	}
}

func Logout(logou app.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		input := request.(*app.Logoutinput)
		val, err := logou.Logout(ctx, input)
		if err != nil {
			return nil, err
		}

		return val, nil
	}
}
