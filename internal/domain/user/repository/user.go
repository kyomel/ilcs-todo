package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kyomel/ilcs-todo/internal/domain/user/model"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) Repository {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) RegisterUser(ctx context.Context, req *model.UserRequest) (*model.User, error) {
	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Email:    req.Email,
		Password: hashedPassword,
		FullName: req.FullName,
	}

	userID := uuid.New()

	query := `
		INSERT INTO users (id, email, password_hash, full_name)
		VALUES ($1, $2, $3, $4)
		RETURNING id, email, password_hash, full_name
	`

	err = r.db.QueryRowContext(ctx, query,
		userID,
		user.Email,
		user.Password,
		user.FullName,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.FullName,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepo) Login(ctx context.Context, req *model.Login) (*model.User, error) {
	user := &model.User{}

	query := `
		SELECT id, email, password_hash, full_name
		FROM users
		WHERE email = $1
	`

	err := r.db.QueryRowContext(ctx, query, req.Email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.FullName,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
