package service

import (
	"context"
	"log"

	"github.com/themy4a4o-hub/mycrypto/internal/models"
	"github.com/themy4a4o-hub/mycrypto/internal/repository"
)

type RateService struct {
	repo *repository.RateRepository
}

func NewRateService(repo *repository.RateRepository) *RateService {
	return &RateService{repo: repo}
}

func (s *RateService) GetCurrentStats(ctx context.Context) ([]models.RateStats, error) {
	cryptos := []string{"BTC"}

	var stats []models.RateStats

	for _, crypto := range cryptos {
		lastRate, err := s.repo.GetLastRate(ctx, crypto)
		if err != nil {
			log.Printf("Ошибка получения последнего курса для %s: %v", crypto, err)
			continue
		}
		if lastRate == nil {
			log.Printf("Нет данных для %s", crypto)
			continue
		}
		minPrice, maxPrice, err := s.repo.GetMinMaxForDay(ctx, crypto)
		if err != nil {
			log.Printf("Ошибка получения мин/макс для %s: %v", crypto, err)
			continue
		}
		var changePercent *float64
		priceHourAgo, err := s.repo.GetRateOneHourAgo(ctx, crypto)
		if err == nil && priceHourAgo > 0 {
			percent := ((lastRate.Price - priceHourAgo) / priceHourAgo) * 100
			changePercent = &percent
		}
		stats = append(stats, models.RateStats{
			Cryptocurrency:    crypto,
			CurrentPrice:      lastRate.Price,
			MinPriceToday:     minPrice,
			MaxPriceToday:     maxPrice,
			ChangePercentHour: changePercent,
		})
	}
	return stats, nil
}
