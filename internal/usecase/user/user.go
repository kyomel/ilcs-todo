package user

import (
	"context"
	"os"
	"time"

	"github.com/kyomel/ilcs-todo/internal/domain/user/model"
	user "github.com/kyomel/ilcs-todo/internal/domain/user/repository"
	util "github.com/kyomel/ilcs-todo/internal/utils/jwt"
)

type useCase struct {
	userRepo   user.Repository
	ctxTimeout time.Duration
}

func NewUserUsecase(userRepo user.Repository, ctxTimeout time.Duration) Usecase {
	return &useCase{
		userRepo:   userRepo,
		ctxTimeout: ctxTimeout,
	}
}

func (uc *useCase) RegisterUser(ctx context.Context, req *model.UserRequest) (*model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.ctxTimeout)
	defer cancel()

	user, err := uc.userRepo.RegisterUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (uc *useCase) Login(ctx context.Context, req *model.Login) (*model.LoginResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.ctxTimeout)
	defer cancel()

	users, err := uc.userRepo.Login(ctx, req)
	if err != nil {
		return nil, err
	}

	if err = user.ComparePassword(users.Password, req.Password); err != nil {
		return nil, err
	}

	token, err := util.GenerateToken(users.ID, users.Email, users.FullName, os.Getenv("JWT_SECRET_KEY"))
	if err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		Token: token,
	}, nil
}
