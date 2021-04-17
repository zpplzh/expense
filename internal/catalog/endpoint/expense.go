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
