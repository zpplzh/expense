package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/zappel/expense-server/internal/catalog"
)

func GetCategory(getcat catalog.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		qry := request.(*catalog.GetCategoryInput)
		viewcat, err := getcat.GetCategory(ctx, qry)
		return viewcat, err
	}
}

func AddCategory(addcat catalog.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		input := request.(*catalog.AddCategoryInput)
		add, err := addcat.AddCategory(ctx, input)

		return add, err
	}
}

func DelCategory(delcat catalog.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		input := request.(*catalog.DelCategoryInput)
		err := delcat.DelCategory(ctx, input)

		return nil, err
	}
}
