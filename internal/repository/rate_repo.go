package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/themy4a4o-hub/mycrypto/internal/models"
)

type RateRepository struct {
	db *sqlx.DB
}

func NewRateRepository(db *sqlx.DB) *RateRepository {
	return &RateRepository{db: db}
}

func (r *RateRepository) SaveRate(ctx context.Context, rate models.Rates) error {
	query := `INSERT INTO cryptoapp.rates (cryptocurrency, price, timestamp) VALUES ($1, $2, $3)`
	_, err := r.db.ExecContext(ctx, query, rate.Cryptocurrency, rate.Price, rate.Timestamp)
	return err
}

func (r *RateRepository) GetLastRate(ctx context.Context, crypto string) (*models.Rates, error) {

	query := `SELECT * FROM cryptoapp.rates WHERE cryptocurrency = $1 ORDER BY timestamp DESC LIMIT 1`
	var rate models.Rates
	err := r.db.GetContext(ctx, &rate, query, crypto)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get last rate: %w", err)
	}
	return &rate, err
}
func (r *RateRepository) GetMinMaxForDay(ctx context.Context, crypto string) (min, max float64, err error) {
	query := `SELECT MIN(price), MAX(price) FROM cryptoapp.rates WHERE cryptocurrency = $1 AND timestamp >= date_trunc('day', now())`
	err = r.db.QueryRowContext(ctx, query, crypto).Scan(&min, &max)
	return min, max, err
}

func (r *RateRepository) GetRateOneHourAgo(ctx context.Context, crypto string) (float64, error) {
	var price float64
	query := `SELECT price FROM cryptoapp.rates WHERE cryptocurrency = $1 AND timestamp <= now() - interval '1 hour' ORDER BY timestamp DESC LIMIT 1`
	err := r.db.QueryRowContext(ctx, query, crypto).Scan(&price)
	return price, err
}
