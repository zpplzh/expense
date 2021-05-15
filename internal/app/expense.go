package app

import (
	"context"
	"time"

	"github.com/segmentio/ksuid"
	null "github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/zappel/expense-server/internal/app/model"
)

type (
	GetExpenseInput struct {
		Id string `json:"Id"`
	}

	AddExpenseInput struct {
		UserID     string  `json: "userid"`
		CategoryId *string `json: "categoryid"`
		//harus pake pointer kalau ga pake pointer pas isi nya null di request jadi error, string yang isi nya "" itu tetap di anggap ada isi
		Amount      int       `json:"amount"`
		Note        string    `json:"note"`
		ExpenseDate time.Time `json:"expensedate"`
	}

	AddExpenseOutput struct {
		Id string `json: "id"`
	}

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
	DelExpenseOutput struct {
		Id string `json:"id"`
	}

	UpdateExpenseInput struct {
		Id          string    `json:"id"`
		CategoryId  string    `json: "categoryid"`
		Amount      int       `json: "amount"`
		Note        string    `json: note`
		ExpenseDate time.Time `json: "expenseDate"`
	}

	UpdateExpenseOutput struct {
		Id string `json: "id"`
	}

	AddExpenseBatchInput struct {
		Data []*AddExpenseInput `json:"data"`
	}

	AddExpenseBatchOutput struct {
	}
)

func (r *servicedb) AddExpense(ctx context.Context, input *AddExpenseInput) (*AddExpenseOutput, error) {

	uid := ctx.Value("uid")

	id := ksuid.New()

	inputex := &model.Expense{
		ID:     id.String(),
		UserID: uid.(string),
		// Categoryid:  null.StringFrom(cat.Categoryid),
		Amount:      input.Amount,
		Note:        null.StringFrom(input.Note),
		ExpenseDate: input.ExpenseDate,
	}

	if input.CategoryId != nil {
		cat, err := model.Categories(qm.Where("categoryid=? and user_id=?", input.CategoryId, uid.(string))).One(ctx, r.db)
		if err != nil {
			return nil, ErrNotFound
		}
		if uid == cat.UserID {
			inputex.Categoryid = null.StringFrom(cat.Categoryid)
		}
	}

	err := inputex.Insert(ctx, r.db, boil.Infer())
	if err != nil {

		return nil, ErrDuplicate
	}

	return &AddExpenseOutput{
		Id: id.String(),
	}, nil
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

func (r *servicedb) DelExpense(ctx context.Context, input *DelExpenseInput) (*DelExpenseOutput, error) {
	uid := ctx.Value(string("uid"))
	_, err := model.Expenses(qm.Where("Id = ? AND user_id=?", input.Id, uid.(string))).DeleteAll(ctx, r.db, true)
	if err != nil {
		return nil, ErrNotFound
	}
	return &DelExpenseOutput{
		Id: input.Id,
	}, nil
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

	return &UpdateExpenseOutput{
		Id: input.Id,
	}, nil

}

func (r *servicedb) AddExpenseBatch(ctx context.Context, input *AddExpenseBatchInput) (*AddExpenseBatchOutput, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	uid := ctx.Value("uid")

	for _, val := range input.Data {
		id := ksuid.New()
		inputex := &model.Expense{
			ID:          id.String(),
			UserID:      uid.(string),
			Amount:      val.Amount,
			Note:        null.StringFrom(val.Note),
			ExpenseDate: val.ExpenseDate,
		}

		if val.CategoryId != nil {
			cat, err := model.Categories(qm.Where("categoryid=? and user_id=?", val.CategoryId, uid.(string))).One(ctx, r.db)
			if err != nil {
				return nil, ErrNotFound
			}
			if uid == cat.UserID {
				inputex.Categoryid = null.StringFrom(cat.Categoryid)
			}
		}

		err := inputex.Insert(ctx, tx, boil.Infer())
		if err != nil {
			//tx.Rollback()
			return nil, ErrDuplicate
		}

	}

	err1 := tx.Commit()
	if err1 != nil {
		return nil, ErrNotFound
	}

	return &AddExpenseBatchOutput{}, nil
}
