package catalog

import (
	"context"
	"math/rand"
	"time"

	"github.com/jmoiron/sqlx"

	null "github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/zappel/expense-server/internal/catalog/model"
	"github.com/zappel/expense-server/internal/catalog/pkgs"
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

type (
	GetExpenseInput struct {
		Id string `json:"Id"`
	}

	AddExpenseInput struct {
		Id          string    `json: "Id"`
		Icon        string    `json:"Icon"`
		Name        string    `json:"CategoryName"`
		Amount      int       `json:"Amount"`
		Note        string    `json:"Note"`
		ExpenseDate time.Time `json:"ExpenseDate"`
	}

	AddExpenseOutput struct{}

	ExpenseOutput struct {
		Id          string    `json: "id"`
		Icon        string    `json: "icon"`
		Name        string    `json: "categoryName"`
		Amount      int       `json: "amount"`
		Note        string    `json: note`
		ExpenseData time.Time `json: "expenseDate"`
	}

	ListExpensesInput struct{}

	DelExpenseInput struct {
		Id string `json:"id"`
	}

	UpdateExpenseInput struct {
		Id string `json:"id"`
	}
)

type (
	SignUpInput struct {
		User_id  string `json:"user_id"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	SignUpOutput struct{}
)

type Service interface {
	GetCategory(ctx context.Context, input *GetCategoryInput) (*CategoryOutput, error)
	AddCategory(ctx context.Context, input *AddCategoryInput) (*CategoryOutput, error)
	DelCategory(ctx context.Context, input *DelCategoryInput) error
	//UpdateCategory(ctx context.Context, input *UpdateCategoryInput) error
	ListCategories(ctx context.Context, input *ListCategoriesInput) ([]*CategoryOutput, error)

	GetExpense(ctx context.Context, input *GetExpenseInput) (*ExpenseOutput, error)
	AddExpense(ctx context.Context, input *AddExpenseInput) (*AddExpenseOutput, error)
	ListExpense(ctx context.Context, input *ListExpensesInput) ([]*ExpenseOutput, error)
	DelExpense(ctx context.Context, input *DelExpenseInput) error
	//UpdateExpense(ctx context.Context, input *UpdateExpenseInput) error

	SignUp(ctx context.Context, input *SignUpInput) (*SignUpOutput, error)
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
		Icon: string(gcat.Icon.String),
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
		Icon: input.Icon,
	}, nil
}

func (r *servicedb) DelCategory(ctx context.Context, input *DelCategoryInput) error {

	_, err := model.Categories(qm.Where("Name = ?", input.Name)).DeleteAll(ctx, r.db, true)
	if err != nil {
		return err
	}

	return nil

}

func (r *servicedb) ListCategories(ctx context.Context, input *ListCategoriesInput) ([]*CategoryOutput, error) {

	allcatarr := []*CategoryOutput{}

	allcat, err := model.Categories(qm.Select(model.CategoryColumns.Name, model.CategoryColumns.Icon)).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	for _, val := range allcat {
		allcatarr = append(allcatarr, &CategoryOutput{
			Name: val.Name,
			Icon: val.Icon.String,
		})
	}

	return allcatarr, nil
}

func (r *servicedb) AddExpense(ctx context.Context, input *AddExpenseInput) (*AddExpenseOutput, error) {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, 30)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}

	inputex := &model.Expense{
		ID:          string(b),
		Icon:        input.Icon,
		Name:        input.Name,
		Amount:      input.Amount,
		Note:        null.StringFrom(input.Note),
		ExpenseDate: input.ExpenseDate,
	}

	err := inputex.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *servicedb) ListExpense(ctx context.Context, input *ListExpensesInput) ([]*ExpenseOutput, error) {

	allexarr := []*ExpenseOutput{}

	allcat, err := model.Expenses(qm.Select("*")).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	for _, val := range allcat {
		allexarr = append(allexarr, &ExpenseOutput{
			Id:          val.ID,
			Name:        val.Name,
			Icon:        val.Icon,
			Amount:      val.Amount,
			Note:        val.Note.String,
			ExpenseData: val.ExpenseDate,
		})
	}

	return allexarr, nil
}

func (r *servicedb) DelExpense(ctx context.Context, input *DelExpenseInput) error {
	// kalau ga pake bintang
	_, err := model.Expenses(qm.Where("Id = ?", input.Id)).DeleteAll(ctx, r.db, true)
	if err != nil {
		return err
	}

	return nil
}

func (r *servicedb) GetExpense(ctx context.Context, input *GetExpenseInput) (*ExpenseOutput, error) {
	getex, err := model.Expenses(qm.Where("Id = ?", input.Id)).One(ctx, r.db)
	if err != nil {
		return nil, err
	}

	return &ExpenseOutput{
		Id:          getex.ID,
		Icon:        getex.Icon,
		Name:        getex.Name,
		Amount:      getex.Amount,
		Note:        getex.Note.String,
		ExpenseData: getex.ExpenseDate,
	}, nil

}

func (r *servicedb) SignUp(ctx context.Context, input *SignUpInput) (*SignUpOutput, error) {
	var h *pkgs.Hash

	p := h.HashandSalt(input.Password)

	inputus := &model.User{
		UserID:   input.User_id,
		Email:    input.Email,
		Password: p,
	}

	err := inputus.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return nil, nil
}
