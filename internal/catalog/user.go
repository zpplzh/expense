package catalog

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/segmentio/ksuid"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/zappel/expense-server/internal/catalog/model"
)

type (
	SignUpInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	SignUpOutput struct{}

	LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginOutput struct {
		Sessionid string `json:"sessionid"`
	}

	Logoutinput struct {
		Sessionid string `json:"sessionid"`
	}
	Logoutoutput struct {
	}
)

func (r *servicedb) SignUp(ctx context.Context, input *SignUpInput) (*SignUpOutput, error) {
	if checkEmail(input.Email) == false || len(input.Password) < 7 {
		return nil, BadInput
	}

	var h *Hash

	p := h.HashandSalt(input.Password)

	uid := ksuid.New()

	inputus := &model.User{
		UserID:   uid.String(),
		Email:    input.Email,
		Password: p,
	}

	err := inputus.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		fmt.Println(err)
		return nil, ErrDuplicate
	}

	return nil, nil
}

func (r *servicedb) Login(ctx context.Context, input *LoginInput) (*LoginOutput, error) {
	//if checkEmail(input.Email) == false {	return nil, BadInput}

	var h *Hash

	gus, err := model.Users(qm.Where("Email = ?", input.Email)).One(ctx, r.db)
	if err != nil {
		return nil, ErrNotFound
	}
	match := h.CheckPass(input.Password, gus.Password)
	if !match {
		return nil, ErrAuth

	}

	tampsess := ""
	count := 0
	for {

		tampsess = RandSessionid()
		exists, err1 := model.Sessions(qm.Where("sessionid=?", string(tampsess))).Exists(ctx, r.db)
		if err1 == nil && exists == false {
			fmt.Println(tampsess, exists)
			break
		} else {
			count++
			fmt.Println(count, tampsess)
			if count == 5 {
				break
			}
		}
	}
	t := strings.Replace(tampsess, "\u0000", "", -1)

	inpt := &model.Session{
		Sessionid: string(t),
		UserID:    gus.UserID,
		Expiry:    time.Now().Add(time.Minute * 15),
	}

	errd := inpt.Insert(ctx, r.db, boil.Infer())
	if errd != nil {
		return nil, ErrDuplicate

	}

	return &LoginOutput{
		Sessionid: tampsess,
	}, nil
}

func (r *servicedb) Logout(ctx context.Context, input *Logoutinput) (*Logoutoutput, error) {

	_, err := model.Sessions(qm.Where("sessionid = ?", input.Sessionid)).DeleteAll(ctx, r.db, true)
	if err != nil {
		return nil, ErrNotFound
	}

	return nil, nil
}
