package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/themy4a4o-hub/mycrypto/internal/config"
	"github.com/themy4a4o-hub/mycrypto/internal/db"
	"github.com/themy4a4o-hub/mycrypto/internal/models"
	"github.com/themy4a4o-hub/mycrypto/internal/repository"
)

func main() {
	cfg := config.Load()

	sqlDB, err := db.NewPostgres(*cfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB connected")
	defer sqlDB.Close()

	sqlxDB := sqlx.NewDb(sqlDB, "postgres")
	ctx := context.Background()
	
	testRate := models.Rates{
		Id:             1,
		Cryptocurrency: "BTC",
		Price:          50000.50,
		Timestamp:      time.Now(),
	}

	rateRepo := repository.NewRateRepository(sqlxDB)
	err = rateRepo.SaveRate(ctx, testRate)
	if err != nil {
		log.Fatalf("Ошибка сохранения в базу данных %v", err)
	}
	log.Printf("Успешно сохранено в базу данных")

	savedRate, err := rateRepo.GetLastRate(ctx, "BTC")
	if err != nil {
		log.Fatalf("Ошибка получения последней записи из базы данных %v", err)
	}
	log.Printf("Успешное чтение из базы данных: %+v", savedRate)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Println("Shutting gracefully down...")
}