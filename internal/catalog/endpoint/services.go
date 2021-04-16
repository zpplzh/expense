package endpoint

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
	"github.com/zappel/expense-server/internal/catalog"
)

func AddExpense(addex catalog.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		input := request.(*catalog.AddExpenseInput)
		_, err := addex.AddExpense(ctx, input)

		return nil, err
	}
}

func ListExpenses(showallex catalog.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var input *catalog.ListExpensesInput
		var expenses []*catalog.ExpenseOutput

		expenses, err := showallex.ListExpense(ctx, input)

		return expenses, err
	}
}

func GetExpense(getex catalog.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		qry := request.(*catalog.GetExpenseInput)
		viewex, err := getex.GetExpense(ctx, qry)
		fmt.Println(viewex)
		return viewex, err
	}
}

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
		var input *catalog.Logoutinput
		val, err := logou.Logout(ctx, input)
		if err != nil {
			return nil, err
		}

		return val, nil
	}
}
