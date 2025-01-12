package user

import (
	"context"

	"github.com/kyomel/ilcs-todo/internal/domain/user/model"
)

type Repository interface {
	RegisterUser(ctx context.Context, req *model.UserRequest) (*model.User, error)
	Login(ctx context.Context, req *model.Login) (*model.User, error)
}
