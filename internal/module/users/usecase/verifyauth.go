package usecase

import (
	"context"
	"hidroponic/internal/errors"
	"hidroponic/internal/module/users/models"

	"golang.org/x/crypto/bcrypt"
)

func (u Usecase) VerifyUsernamePassword(ctx context.Context, username, password string) (*models.User, *errors.BaseError) {
	user, err := u.repo.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if !checkPasswordHash(password, user.Password) {
		return nil, errors.ErrorInvalidPassword.New("password do not match")
	}

	return u.mapEntityToModel(user), nil
}

func checkPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
