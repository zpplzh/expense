package catalog

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/zappel/expense-server/internal/catalog/model"
)

type (
	GetCategoryInput struct {
		UserID string `json: "Userid"`
		Name   string `json: "CategoryName"`
	}

	AddCategoryInput struct {
		Userid string `json: "userid"`
		Icon   string `json:"Icon"`
		Name   string `json:"categoryName"`
	}

	CategoryOutput struct {
		Icon string `json:"Icon"`
		Name string `json: "CategoryName"`
	}

	DelCategoryInput struct {
		Name string `json: "CategoryName"`
	}

	ListCategoriesInput struct{}

	UpdateCategoryInput struct {
		Id string `json:"Id"`
	}
)

func (r *servicedb) GetCategory(ctx context.Context, input *GetCategoryInput) (*CategoryOutput, error) {
	sid := ctx.Value("sessionid")
	gusid, err1 := model.Sessions(qm.Where("sessionid = ?", sid)).One(ctx, r.db)
	if err1 != nil {
		return nil, ErrNotFound
	}

	gcat, err := model.Categories(qm.Where("Name = ? and user_id =?", input.Name, gusid.UserID)).One(ctx, r.db)
	if err != nil {
		return nil, ErrNotFound
	}

	return &CategoryOutput{
		Name: gcat.Name,
		Icon: string(gcat.Icon),
	}, nil

}

func (r *servicedb) AddCategory(ctx context.Context, input *AddCategoryInput) (*CategoryOutput, error) {
	if input.Name == "" || input.Icon == "" {
		return nil, BadInput
	}

	sid := ctx.Value("sessionid")

	gusid, err1 := model.Sessions(qm.Where("sessionid = ?", sid)).One(ctx, r.db)
	if err1 != nil {
		return nil, ErrNotFound
	}

	c := &model.Category{
		UserID: gusid.UserID,
		Name:   input.Name,
		Icon:   input.Icon,
	}

	err := c.Insert(ctx, r.db, boil.Infer())
	if err != nil {

		return nil, ErrDuplicate
	}

	return &CategoryOutput{
		Name: input.Name,
		Icon: input.Icon,
	}, nil
}

func (r *servicedb) DelCategory(ctx context.Context, input *DelCategoryInput) error {

	_, err := model.Categories(qm.Where("Name = ?", input.Name)).DeleteAll(ctx, r.db, true)
	if err != nil {
		return ErrNotFound
	}

	return nil

}

func (r *servicedb) ListCategories(ctx context.Context, input *ListCategoriesInput) ([]*CategoryOutput, error) {

	allcatarr := []*CategoryOutput{}

	allcat, err := model.Categories(qm.Select(model.CategoryColumns.Name, model.CategoryColumns.Icon)).All(ctx, r.db)
	if err != nil {
		return nil, ErrNotFound
	}

	for _, val := range allcat {
		allcatarr = append(allcatarr, &CategoryOutput{
			Name: val.Name,
			Icon: val.Icon,
		})
	}

	return allcatarr, nil
}
