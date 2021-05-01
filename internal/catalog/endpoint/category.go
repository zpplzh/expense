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
		if err != nil {
			return nil, err
		}
		return viewcat, nil
	}
}

func AddCategory(addcat catalog.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		input := request.(*catalog.AddCategoryInput)
		add, err := addcat.AddCategory(ctx, input)
		if err != nil {
			return nil, err
		}
		return add, nil
	}
}

func DelCategory(delcat catalog.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		input := request.(*catalog.DelCategoryInput)
		err := delcat.DelCategory(ctx, input)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

func ListCategories(showallcat catalog.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var input *catalog.ListCategoriesInput
		var categories []*catalog.CategoryOutput

		categories, err := showallcat.ListCategories(ctx, input)
		if err != nil {
			return nil, err
		}

		return categories, nil
	}
}

func UpdateCategory(upcat catalog.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		input := request.(*catalog.UpdateCategoryInput)
		upc, err := upcat.UpdateCategory(ctx, input)
		if err != nil {
			return nil, err
		}
		return upc, nil
	}
}
