package repositories

import (
	"context"
	"github.com/RobertJaskolski/go-REST-api/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepositoryInterface interface {
	Create(ctx, dto *models.CreateUserDTO) (*models.User, error)
}

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(ctx context.Context, dto *models.CreateUserDTO) (*models.User, error) {
	query := `INSERT INTO users (email, first_name, last_name, time_zone, mobile, role, is_active, password) VALUES (@Email, @FirstName, @LastName, @TimeZone, @Mobile, @Role, @IsActive, @Password) RETURNING (id, email, first_name, last_name, time_zone, mobile, role, is_active, created_at, modified_at);`
	args := pgx.NamedArgs{
		"Email":     dto.Email,
		"FirstName": dto.FirstName,
		"LastName":  dto.LastName,
		"TimeZone":  dto.TimeZone,
		"Mobile":    dto.Mobile,
		"Role":      dto.Role,
		"IsActive":  dto.IsActive,
		"Password":  dto.Password,
	}

	user := new(models.User)
	err := r.db.QueryRow(ctx, query, args).Scan(&user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
