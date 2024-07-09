package usecase

import (
	"hidroponic/internal/module/users/entities"
	"hidroponic/internal/module/users/models"
	"hidroponic/internal/module/users/ports"
	"time"
)

type Usecase struct {
	repo ports.Repository
}

func New(repo ports.Repository) ports.Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (Usecase) mapEntityToModel(userEntity *entities.User) *models.User {
	return &models.User{
		ID:        userEntity.ID,
		Name:      userEntity.Name,
		Username:  userEntity.Username,
		Email:     userEntity.Email,
		CreatedAt: userEntity.CreatedAt.Format(time.RFC3339),
	}
}
