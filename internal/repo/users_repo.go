package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/murashi19/koda-b8-backend1/internal/models"
)

type UserRepo struct {
	// data *[]models.User
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) Create(ctx context.Context, data *models.User) (*models.User, error) {

	sql := `
		INSERT INTO users(email, password, username, phone) VALUES
		($1, $2, $3, $4)
		RETURNING id, email, password, username,phone,created_at,updated_at;
	`
	user, err := oneRow[models.User](ctx, r.db, sql, data.Email, data.Password, data.Username, data.Phone)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, fmt.Errorf("Email already registered")
		}
		return nil, err
	}
	return user, nil
}

func (r *UserRepo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
	SELECT
		id,
		email,
		password,
		username,
		phone,
		created_at,
		updated_at
	FROM users
	WHERE email = $1
	`

	return oneRow[models.User](ctx, r.db, query, email)
}

func (r *UserRepo) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	sql := `SELECT
			id,
			email,
			username,
			phone
			FROM users
			ORDER BY id`
	return rows[models.User](ctx, r.db, sql)
}
