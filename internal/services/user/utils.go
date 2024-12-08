package user

import (
	"errors"
	"time"

	constant "github.com/wendao2000/couply/internal/constants"
	model "github.com/wendao2000/couply/internal/models"

	"github.com/wendao2000/couply/pkg/errs"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// saltPassword adds salt to password for additional security
func (s *service) saltPassword(password string) string {
	return password + s.cfg.AuthPassSalt
}

// validatePassword validates if provided password matches stored hash
func (s *service) validatePassword(realPassword []byte, password string) (err *errs.Error) {
	err2 := bcrypt.CompareHashAndPassword(realPassword, []byte(password))
	if err2 != nil {
		if errors.Is(err2, bcrypt.ErrMismatchedHashAndPassword) {
			return errs.AuthError.WithMessage("invalid password")
		}
		return errs.InternalError.WithMessage(err2.Error())
	}
	return nil
}

// generateToken generates JWT token containing user info and expiry
func (s *service) generateToken(user *model.User) (userToken *model.UserToken, err *errs.Error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     now.Add(constant.AuthTokenExp).Unix(),
		"iat":     now.Unix(),
		"iss":     s.cfg.AuthJWTIssuer,
	})

	accessToken, err2 := token.SignedString([]byte(s.cfg.AuthJWTSecret))
	if err2 != nil {
		return nil, errs.InternalError.WithMessage(err2.Error())
	}

	// TODO: implement refresh token

	return &model.UserToken{
		Token:     accessToken,
		TokenType: constant.AuthBearer,
		ExpiresIn: int64(constant.AuthTokenExp / time.Second), // expires_in in second
	}, nil
}
