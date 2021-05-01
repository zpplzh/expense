package catalog

import (
	"context"
	"fmt"

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
		Categoryid string `json:"categoryid"`
		Name       string `json:"name"`
		Icon       string `json:"icon"`
	}

	UpdateCategoryOutput struct{}
)

func (r *servicedb) GetCategory(ctx context.Context, input *GetCategoryInput) (*CategoryOutput, error) {
	uid := ctx.Value("uid")

	gcat, err := model.Categories(qm.Where("Name = ? and user_id =?", input.Name, uid.(string))).One(ctx, r.db)
	if err != nil {
		return nil, ErrNotFound
	}

	return &CategoryOutput{
		Name: gcat.Name,
		Icon: string(gcat.Icon),
	}, nil

}

func (r *servicedb) AddCategory(ctx context.Context, input *AddCategoryInput) (*CategoryOutput, error) {
	if input.Name == "" || input.Icon == "" || checkInput(input.Name) == false {
		return nil, BadInput
	}

	uid := ctx.Value(string("uid"))

	exists, err1 := model.Categories(qm.Where("user_id=? AND icon =? AND name=?", uid.(string), input.Icon, input.Name)).Exists(ctx, r.db)
	if err1 != nil || exists == true {
		return nil, DataExistErr
	}

	c := &model.Category{

		UserID: uid.(string),
		Name:   input.Name,
		Icon:   input.Icon,
	}

	fmt.Println(uid.(string))

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
	uid := ctx.Value(string("uid"))

	exists, err1 := model.Categories(qm.Where("user_id=? AND name=?", uid.(string), input.Name)).Exists(ctx, r.db)
	if err1 != nil || exists == false {
		return ErrNotFound
	}

	_, err := model.Categories(qm.Where("name = ? AND user_id=?", input.Name, uid.(string))).DeleteAll(ctx, r.db, true)
	if err != nil {
		fmt.Println(err)
		return ErrNotFound
	}

	return nil

}

func (r *servicedb) ListCategories(ctx context.Context, input *ListCategoriesInput) ([]*CategoryOutput, error) {
	allcatarr := []*CategoryOutput{}
	uid := ctx.Value(string("uid"))
	allcat, err := model.Categories(qm.Select(model.CategoryColumns.Name, model.CategoryColumns.Icon), qm.Where("user_id=?", uid.(string))).All(ctx, r.db)
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

func (r *servicedb) UpdateCategory(ctx context.Context, input *UpdateCategoryInput) (*UpdateCategoryOutput, error) {
	if input.Name == "" || input.Icon == "" || !checkInput(input.Name) {
		return nil, BadInput
	}
	fmt.Println(checkInput(input.Name))

	upca, err := model.FindCategory(ctx, r.db, input.Categoryid, name string, selectCols ...string)
	if err != nil {
		return nil, ErrNotFound
	}
	upca.Name = input.Name
	upca.Icon = input.Icon
	RowsAffected, err := upca.Update(ctx, r.db, boil.Infer())
	if err != nil && RowsAffected == 0 {
		return nil, ErrNotFound
	}

	return nil, nil

}
