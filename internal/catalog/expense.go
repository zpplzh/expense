package catalog

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/ksuid"

	null "github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/zappel/expense-server/internal/catalog/model"
)

type (
	GetExpenseInput struct {
		Id string `json:"Id"`
	}

	AddExpenseInput struct {
		UserID      string    `json: "Userid"`
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

func (r *servicedb) AddExpense(ctx context.Context, input *AddExpenseInput) (*AddExpenseOutput, error) {

	if input.Icon == "" || input.Name == "" || checkInput(input.Name) == false {
		return nil, BadInput
	}

	uid := ctx.Value("uid")

	ex, err2 := model.Categories(qm.Where("name=? and user_id=?", input.Name, uid.(string))).One(ctx, r.db)
	if err2 != nil {
		return nil, ErrNotFound
	}

	fmt.Println(ex.Icon, ex.UserID)

	if ex.Icon == input.Icon && ex.Name == input.Name && uid.(string) == ex.UserID {
		id := ksuid.New()

		inputex := &model.Expense{
			UserID:      uid.(string),
			ID:          id.String(),
			Icon:        input.Icon,
			Name:        input.Name,
			Amount:      input.Amount,
			Note:        null.StringFrom(input.Note),
			ExpenseDate: input.ExpenseDate,
		}

		err := inputex.Insert(ctx, r.db, boil.Infer())
		if err != nil {

			return nil, ErrDuplicate
		}
	} else {
		return nil, ErrNotFound
	}

	return nil, nil
}

func (r *servicedb) ListExpense(ctx context.Context, input *ListExpensesInput) ([]*ExpenseOutput, error) {
	allexarr := []*ExpenseOutput{}
	uid := ctx.Value(string("uid"))
	allex, err := model.Expenses(qm.Select("*"), qm.Where("user_id=?", uid.(string))).All(ctx, r.db)
	if err != nil {
		return nil, ErrNotFound
	}

	for _, val := range allex {
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
	uid := ctx.Value(string("uid"))
	_, err := model.Expenses(qm.Where("Id = ? AND user_id=?", input.Id, uid.(string))).DeleteAll(ctx, r.db, true)
	if err != nil {
		return ErrNotFound
	}

	return nil
}

func (r *servicedb) GetExpense(ctx context.Context, input *GetExpenseInput) (*ExpenseOutput, error) {
	getex, err := model.Expenses(qm.Where("Id = ?", input.Id)).One(ctx, r.db)
	if err != nil {
		return nil, ErrNotFound
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
