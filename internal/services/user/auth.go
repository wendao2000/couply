package user

import (
	"context"
	"log"

	model "github.com/wendao2000/couply/internal/models"
	repoUser "github.com/wendao2000/couply/internal/repositories/user"

	"github.com/wendao2000/couply/pkg/errs"
)

func (s *service) CreateUser(ctx context.Context, req *model.Auth) (user *model.User, userToken *model.UserToken, err *errs.Error) {
	if req == nil {
		log.Println("[CreateUser] Request is nil")
		return nil, nil, errs.ParamError.WithMessage("Empty request")
	}
	if len(req.Email) == 0 && len(req.Username) == 0 {
		log.Println("[CreateUser] Both email and username is empty")
		return nil, nil, errs.ParamError.WithMessage("Email and/or username cannot be empty")
	}
	// TODO: validate email format
	if len(req.Password) == 0 {
		log.Println("[CreateUser] Password is empty")
		return nil, nil, errs.ParamError.WithMessage("Password cannot be empty")
	}

	user, err = s.userRepo.InsertUser(ctx, &repoUser.InsertUserRequest{
		Email:          req.Email,
		Username:       req.Username,
		SaltedPassword: s.saltPassword(req.Password),
	})
	if err != nil {
		log.Println("[CreateUser] Failed while inserting user to database, err:", err)
		return nil, nil, err
	}

	userToken, err = s.generateToken(user)
	if err != nil {
		log.Println("[CreateUser] Failed while generating jwt token, err:", err)
		return nil, nil, err
	}

	return user, userToken, nil
}

func (s *service) Authenticate(ctx context.Context, req *model.Auth) (user *model.User, userToken *model.UserToken, err *errs.Error) {
	if req == nil {
		log.Println("[CreateUser] Request is nil")
		return nil, nil, errs.ParamError.WithMessage("Empty request")
	}
	if len(req.Email) == 0 && len(req.Username) == 0 {
		log.Println("[CreateUser] Both email and username is empty")
		return nil, nil, errs.ParamError.WithMessage("Email and/or username cannot be empty")
	}
	// TODO: validate email format
	if len(req.Password) == 0 {
		log.Println("[CreateUser] Password is empty")
		return nil, nil, errs.ParamError.WithMessage("Password cannot be empty")
	}

	if len(req.Email) > 0 { // Authenticate by email
		user, err = s.userRepo.GetUserByEmail(ctx, req.Email)
	} else { // Authenticate by username
		user, err = s.userRepo.GetUserByUsername(ctx, req.Username)
	}
	if err != nil {
		log.Println("[CreateUser] Failed while getting user from database, err:", err)
		return nil, nil, err
	}

	// Validate password
	err = s.validatePassword(user.Password, s.saltPassword(req.Password))
	if err != nil {
		log.Println("[CreateUser] Failed hwile validating user password, err:", err)
		return nil, nil, err
	}

	userToken, err = s.generateToken(user)
	if err != nil {
		log.Println("[CreateUser] Failed while generating jwt token, err:", err)
		return nil, nil, err
	}

	return user, userToken, nil
}
