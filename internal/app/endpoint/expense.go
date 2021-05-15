package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/zappel/expense-server/internal/app"
)

func AddExpense(addex app.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		input := request.(*app.AddExpenseInput)
		add, err := addex.AddExpense(ctx, input)

		return add, err
	}
}

func ListExpenses(showallex app.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var input *app.ListExpensesInput
		var expenses []*app.ExpenseOutput

		expenses, err := showallex.ListExpense(ctx, input)

		return expenses, err
	}
}

func GetExpense(getex app.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		qry := request.(*app.GetExpenseInput)
		viewex, err := getex.GetExpense(ctx, qry)
		return viewex, err
	}
}

func DelExpense(delex app.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		input := request.(*app.DelExpenseInput)
		del, err := delex.DelExpense(ctx, input)
		if err != nil {
			return nil, err
		}

		return del, nil
	}
}

func UpdateExpense(upex app.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		input := request.(*app.UpdateExpenseInput)
		upc, err := upex.UpdateExpense(ctx, input)
		if err != nil {
			return nil, err
		}
		return upc, nil
	}
}

func AddExpenseBatch(addex app.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		input := request.(*app.AddExpenseBatchInput)
		add, err := addex.AddExpenseBatch(ctx, input)

		return add, err
	}
}
