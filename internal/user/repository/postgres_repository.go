package repository

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/johnqr/user-service/internal/user/domain"
)

type PostgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{pool: pool}
}

func (p *PostgresRepository) Create(ctx context.Context, u *domain.User) error {
	q := `
		INSERT INTO users (id, name, email, password_hash, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}

	now := time.Now().UTC()
	u.CreatedAt = now
	u.UpdatedAt = now

	_, err := p.pool.Exec(ctx, q,
		u.ID,
		u.Name,
		u.Email,
		u.PasswordHash,
		u.CreatedAt,
		u.UpdatedAt,
	)

	if err != nil {
		// Detectar unique violation (código 23505)
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return ErrConflict
		}
		return err
	}

	return nil
}

func (p *PostgresRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	q := `
		SELECT id, name, email, password_hash, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	row := p.pool.QueryRow(ctx, q, id)
	u := &domain.User{}

	err := row.Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.PasswordHash,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return u, nil
}

func (p *PostgresRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	q := `
		SELECT id, name, email, password_hash, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	row := p.pool.QueryRow(ctx, q, email)
	u := &domain.User{}

	err := row.Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.PasswordHash,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return u, nil
}
