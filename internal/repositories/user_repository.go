package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/RobertJaskolski/go-REST-api/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepositoryInterface interface {
	Create(ctx context.Context, dto *models.CreateUserDTO) (*models.User, error)
	GetOne(ctx context.Context, id int) (*models.User, error)
	Update(ctx context.Context, id int, dto *models.UpdateUserDTO) (*models.User, error)
	Delete(ctx context.Context, id int) error
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
	err := r.db.QueryRow(ctx, createUserQuery, args).Scan(&user)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return nil, errors.New(pgErr.Detail)
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetOne(ctx context.Context, id int) (*models.User, error) {
	args := pgx.NamedArgs{
		"ID": id,
	}

	user := new(models.User)
	err := r.db.QueryRow(ctx, getUserQuery, args).Scan(&user)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return nil, errors.New(pgErr.Detail)
	}

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New(fmt.Sprintf("User with ID %d not found", id))
		}
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Update(ctx context.Context, id int, dto *models.UpdateUserDTO) (*models.User, error) {
	args := pgx.NamedArgs{
		"ID":        id,
		"Email":     dto.Email,
		"FirstName": dto.FirstName,
		"LastName":  dto.LastName,
		"TimeZone":  dto.TimeZone,
		"Mobile":    dto.Mobile,
		"Role":      dto.Role,
		"IsActive":  dto.IsActive,
	}

	user := new(models.User)
	fmt.Println(args)
	err := r.db.QueryRow(ctx, updateUserQuery, args).Scan(&user)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return nil, errors.New(pgErr.Detail)
	}

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New(fmt.Sprintf("User with ID %d not found", id))
		}
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Delete(ctx context.Context, id int) error {
	args := pgx.NamedArgs{
		"ID": id,
	}

	a, err := r.db.Exec(ctx, deleteUserQuery, args)

	if a.RowsAffected() == 0 {
		return errors.New(fmt.Sprintf("User with ID %d not exist", id))
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return errors.New(pgErr.Detail)
	}

	if err != nil {
		return err
	}

	return nil
}
