package user

import (
	"context"
	"log"

	model "github.com/wendao2000/couply/internal/models"
	repoUser "github.com/wendao2000/couply/internal/repositories/user"

	"github.com/wendao2000/couply/pkg/errs"
	"github.com/wendao2000/couply/pkg/utils"
)

// CreateUser creates new user account with email/username and password, returns user data with JWT token
func (s *service) CreateUser(ctx context.Context, req *model.Auth) (user *model.User, userToken *model.UserToken, err *errs.Error) {
	if req == nil {
		log.Println("[CreateUser] Request is nil")
		return nil, nil, errs.ParamError.WithMessage("Empty request")
	}

	if len(req.Email) == 0 && len(req.Username) == 0 {
		log.Println("[CreateUser] Both email and username is empty")
		return nil, nil, errs.ParamError.WithMessage("Email and/or username cannot be empty")
	}

	// validate email format
	if len(req.Email) > 0 {
		if !utils.ValidateEmail(req.Email) {
			log.Println("[CreateUser] Invalid email format")
			return nil, nil, errs.ParamError.WithMessage("Invalid email")
		}
	}

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

// Authenticate validates user credentials (email/username + password) and returns user data with new JWT token
func (s *service) Authenticate(ctx context.Context, req *model.Auth) (user *model.User, userToken *model.UserToken, err *errs.Error) {
	if req == nil {
		log.Println("[Authenticate] Request is nil")
		return nil, nil, errs.ParamError.WithMessage("Empty request")
	}

	if len(req.Email) == 0 && len(req.Username) == 0 {
		log.Println("[Authenticate] Both email and username is empty")
		return nil, nil, errs.ParamError.WithMessage("Email and/or username cannot be empty")
	}

	// validate email format
	if len(req.Email) > 0 {
		if !utils.ValidateEmail(req.Email) {
			log.Println("[Authenticate] Invalid email format")
			return nil, nil, errs.ParamError.WithMessage("Invalid email")
		}
	}

	if len(req.Password) == 0 {
		log.Println("[Authenticate] Password is empty")
		return nil, nil, errs.ParamError.WithMessage("Password cannot be empty")
	}

	if len(req.Email) > 0 { // Authenticate by email
		user, err = s.userRepo.GetUserByEmail(ctx, req.Email)
	} else { // Authenticate by username
		user, err = s.userRepo.GetUserByUsername(ctx, req.Username)
	}
	if err != nil {
		log.Println("[Authenticate] Failed while getting user from database, err:", err)
		return nil, nil, err
	}

	// Validate password
	err = s.validatePassword(user.Password, s.saltPassword(req.Password))
	if err != nil {
		log.Println("[Authenticate] Failed hwile validating user password, err:", err)
		return nil, nil, err
	}

	userToken, err = s.generateToken(user)
	if err != nil {
		log.Println("[Authenticate] Failed while generating jwt token, err:", err)
		return nil, nil, err
	}

	return user, userToken, nil
}
