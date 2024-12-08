package user

import (
	"context"
	"strings"

	model "github.com/wendao2000/couply/internal/models"

	"github.com/wendao2000/couply/pkg/errs"

	"golang.org/x/crypto/bcrypt"
)

func (r *repository) InsertUser(ctx context.Context, req *InsertUserRequest) (user *model.User, err *errs.Error) {
	pwhashed, err2 := bcrypt.GenerateFromPassword([]byte(req.SaltedPassword), bcrypt.DefaultCost)
	if err2 != nil {
		return nil, errs.InternalError.WithMessage(err2.Error())
	}

	user = &model.User{
		Email:    req.Email,
		Username: req.Username,
		Password: pwhashed,
	}

	if res := r.db.Create(user); res.Error != nil {
		if strings.Contains(res.Error.Error(), "UNIQUE constraint failed") {
			return nil, errs.AuthError.WithMessage("User already exists")
		}
		return nil, errs.InternalError.WithMessage(res.Error.Error())
	}

	return user, nil
}
