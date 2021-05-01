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

func DelExpense(delex catalog.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		input := request.(*catalog.DelExpenseInput)
		err := delex.DelExpense(ctx, input)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

func UpdateExpense(upex catalog.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		input := request.(*catalog.UpdateExpenseInput)
		upc, err := upex.UpdateExpense(ctx, input)
		if err != nil {
			return nil, err
		}
		return upc, nil
	}
}
