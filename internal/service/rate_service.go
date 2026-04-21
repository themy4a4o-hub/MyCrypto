package service

import (
	"context"

	"github.com/themy4a4o-hub/mycrypto/internal/repository"
)

type RateService struct {
	repo *repository.RateRepository
}

func NewRateService(repo *repository.RateRepository) *RateService {
	return &RateService{repo: repo}
}

func (s *RateService) GetCurrentStats(ctx context.Context) ([]models.RateStats, error) {
	query:= `SELECT * FROM rates `
}
