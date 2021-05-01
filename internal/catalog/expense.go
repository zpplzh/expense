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
		UserID      string    `json: "userid"`
		CategoryId  string    `json: "categoryod"`
		Amount      int       `json:"amount"`
		Note        string    `json:"note"`
		ExpenseDate time.Time `json:"expensedate"`
	}

	AddExpenseOutput struct{}

	ExpenseOutput struct {
		Id          string    `json: "id"`
		CategoryId  string    `json: "categoryid"`
		Amount      int       `json: "amount"`
		Note        string    `json: note`
		ExpenseDate time.Time `json: "expenseDate"`
	}

	ListExpensesInput struct{}

	DelExpenseInput struct {
		Id string `json:"id"`
	}

	UpdateExpenseInput struct {
		Id          string    `json:"id"`
		CategoryId  string    `json: "categoryid"`
		Amount      int       `json: "amount"`
		Note        string    `json: note`
		ExpenseDate time.Time `json: "expenseDate"`
	}

	UpdateExpenseOutput struct{}
)

func (r *servicedb) AddExpense(ctx context.Context, input *AddExpenseInput) (*AddExpenseOutput, error) {

	uid := ctx.Value("uid")

	cat, err := model.Categories(qm.Where("categoryid=? and user_id=?", input.CategoryId, uid.(string))).One(ctx, r.db)
	if err != nil {
		fmt.Println(err)
	}

	id := ksuid.New()

	if uid.(string) == cat.UserID || cat.Categoryid == "" {
		inputex := &model.Expense{
			ID:          id.String(),
			UserID:      uid.(string),
			Categoryid:  null.StringFrom(cat.Categoryid),
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
			CategoryId:  val.Categoryid.String,
			Amount:      val.Amount,
			Note:        val.Note.String,
			ExpenseDate: val.ExpenseDate,
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
		CategoryId:  getex.Categoryid.String,
		Amount:      getex.Amount,
		Note:        getex.Note.String,
		ExpenseDate: getex.ExpenseDate,
	}, nil

}

func (r *servicedb) UpdateExpense(ctx context.Context, input *UpdateExpenseInput) (*UpdateExpenseOutput, error) {
	if input.Amount <= 0 {
		return nil, BadInput
	}

	upex, err := model.FindExpense(ctx, r.db, input.Id)
	if err != nil {
		return nil, ErrNotFound
	}
	upex.Categoryid = null.StringFrom(input.CategoryId)
	upex.Amount = input.Amount
	upex.ExpenseDate = input.ExpenseDate
	upex.Note = null.StringFrom(input.Note)
	RowsAffected, err := upex.Update(ctx, r.db, boil.Infer())
	if err != nil && RowsAffected == 0 {
		return nil, ErrNotFound
	}

	return nil, nil

}
