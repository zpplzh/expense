package catalog

import (
	"context"

	"github.com/jmoiron/sqlx"
	null "github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/zappel/expense-server/internal/catalog/model"
)

type (
	GetCategoryInput struct {
		Name string `json: "CategoryName"`
	}

	AddCategoryInput struct {
		Icon string `json:"Icon"`
		Name string `json:"categoryName"`
	}

	CategoryOutput struct {
		Icon null.String `json:"Icon"`
		Name string      `json: "CategoryName"`
	}

	DelCategoryInput struct {
		Name string `json: "CategoryName"`
	}
)

type Service interface {
	GetCategory(ctx context.Context, input *GetCategoryInput) (*CategoryOutput, error)
	AddCategory(ctx context.Context, input *AddCategoryInput) (*CategoryOutput, error)
	DelCategory(ctx context.Context, input *DelCategoryInput) error
}

type servicedb struct {
	db *sqlx.DB
}

func NewServices(db *sqlx.DB) Service {
	return &servicedb{
		db: db,
	}
}

func (r *servicedb) GetCategory(ctx context.Context, input *GetCategoryInput) (*CategoryOutput, error) {

	gcat, err := model.Categories(qm.Where("Name = ?", input.Name)).One(ctx, r.db)
	if err != nil {
		return nil, err
	}

	return &CategoryOutput{
		Name: gcat.Name,
		Icon: gcat.Icon,
	}, nil

}

func (r *servicedb) AddCategory(ctx context.Context, input *AddCategoryInput) (*CategoryOutput, error) {

	c := &model.Category{
		Name: input.Name,
		Icon: null.StringFrom(input.Icon),
	}

	err := c.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return &CategoryOutput{
		Name: input.Name,
		Icon: null.StringFrom(input.Icon),
	}, nil
}

func (r *servicedb) DelCategory(ctx context.Context, input *DelCategoryInput) error {

	_, err := model.Categories(qm.Where("Name = ?", input.Name)).DeleteAll(ctx, r.db, true)
	if err != nil {
		return err
	}

	return nil

}
