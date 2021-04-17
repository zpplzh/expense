package endpoint

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
	"github.com/zappel/expense-server/internal/catalog"
)

func SignUp(siup catalog.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		input := request.(*catalog.SignUpInput)
		add, err := siup.SignUp(ctx, input)
		fmt.Println(input, "rr")

		return add, err
	}
}

func Login(logi catalog.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		input := request.(*catalog.LoginInput)
		ex, err := logi.Login(ctx, input)

		return ex, err
	}
}

func Logout(logou catalog.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		input := request.(*catalog.Logoutinput)
		val, err := logou.Logout(ctx, input)
		if err != nil {
			return nil, err
		}

		return val, nil
	}
}
