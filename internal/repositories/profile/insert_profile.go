package profile

import (
	"context"
	"strings"

	model "github.com/wendao2000/couply/internal/models"

	"github.com/wendao2000/couply/pkg/errs"
	"github.com/wendao2000/couply/pkg/utils"
)

func (r *repository) InsertProfile(ctx context.Context, req *InsertProfileRequest) (prof *model.Profile, err *errs.Error) {
	dob, err2 := utils.ParseDOB(req.DateOfBirth)
	if err2 != nil {
		return nil, errs.ParamError.WithMessage(err2.Error())
	}

	prof = &model.Profile{
		UserID:      req.UserID,
		Name:        req.Name,
		DateOfBirth: dob,
	}

	if res := r.db.Create(prof); res.Error != nil {
		if strings.Contains(res.Error.Error(), "UNIQUE constraint failed") {
			return nil, errs.AuthError.WithMessage("User already exists")
		}
		return nil, errs.InternalError.WithMessage(res.Error.Error())
	}

	return prof, nil
}
