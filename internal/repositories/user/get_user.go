package user

import (
	"context"
	"errors"

	model "github.com/wendao2000/couply/internal/models"

	"github.com/wendao2000/couply/pkg/errs"

	"gorm.io/gorm"
)

func (r *repository) GetUserByEmail(ctx context.Context, email string) (user *model.User, err *errs.Error) {
	res := r.db.Select([]string{"id", "email", "username", "password"}).First(&user, "email = ?", email)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, errs.NotFoundError
		}
		return nil, errs.InternalError.WithMessage(res.Error.Error())
	}
	return user, nil
}

func (r *repository) GetUserByUsername(ctx context.Context, username string) (user *model.User, err *errs.Error) {
	res := r.db.Select([]string{"id", "email", "username", "password"}).First(&user, "username = ?", username)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, errs.NotFoundError
		}
		return nil, errs.InternalError.WithMessage(res.Error.Error())
	}
	return user, nil
}
