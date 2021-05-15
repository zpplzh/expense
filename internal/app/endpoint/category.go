package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/zappel/expense-server/internal/app"
)

func GetCategory(getcat app.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		qry := request.(*app.GetCategoryInput)
		viewcat, err := getcat.GetCategory(ctx, qry)
		if err != nil {
			return nil, err
		}
		return viewcat, nil
	}
}

func AddCategory(addcat app.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		input := request.(*app.AddCategoryInput)
		add, err := addcat.AddCategory(ctx, input)
		if err != nil {
			return nil, err
		}
		return add, nil
	}
}

func DelCategory(delcat app.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		input := request.(*app.DelCategoryInput)
		del, err := delcat.DelCategory(ctx, input)
		if err != nil {
			return nil, err
		}

		return del, nil
	}
}

func ListCategories(showallcat app.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var input *app.ListCategoriesInput
		var categories []*app.CategoryOutput

		categories, err := showallcat.ListCategories(ctx, input)
		if err != nil {
			return nil, err
		}

		return categories, nil
	}
}

func UpdateCategory(upcat app.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		input := request.(*app.UpdateCategoryInput)
		upc, err := upcat.UpdateCategory(ctx, input)
		if err != nil {
			return nil, err
		}
		return upc, nil
	}
}
