package profile

import (
	"context"
	"log"

	model "github.com/wendao2000/couply/internal/models"
	profRepo "github.com/wendao2000/couply/internal/repositories/profile"
	"github.com/wendao2000/couply/pkg/errs"
)

func (s *service) CreateBasicProfile(ctx context.Context, req *model.BasicProfile) (prof *model.Profile, err *errs.Error) {
	if req == nil {
		log.Println("[CreateBasicProfile] Request is nil")
		return nil, errs.ParamError.WithMessage("Empty request")
	}

	if req.UserID == 0 {
		log.Println("[CreateBasicProfile] Invalid user id")
		return nil, errs.ParamError.WithMessage("Invalid user id")
	}

	prof, err = s.profRepo.InsertProfile(ctx, &profRepo.InsertProfileRequest{
		UserID:      req.UserID,
		Name:        req.Name,
		DateOfBirth: req.DateOfBirth,
	})
	if err != nil {
		// TODO: retry mechanism
		log.Println("[CreateBasicProfile] Failed while inserting profile to database, err:", err)
		return nil, err
	}

	return prof, nil
}
