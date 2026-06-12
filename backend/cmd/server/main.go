package main

import (
	"log"

	"github.com/candelatorrez/northwind-app/db/seed"
	"github.com/candelatorrez/northwind-app/internal/api"
	"github.com/candelatorrez/northwind-app/internal/config"
	"github.com/candelatorrez/northwind-app/internal/database"
	"github.com/candelatorrez/northwind-app/internal/repository"
	"github.com/candelatorrez/northwind-app/internal/service"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load("../.env"); err != nil {
		log.Println("warning: .env not found")
	}

	cfg := config.Load()

	db, err := database.Connect(
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.DBUser,
		cfg.DBPassword,
	)

	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatal(err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("database connected")

	if err := database.Migrate(db); err != nil {
		log.Fatal(err)
	}

	log.Println("database migrated")

	if err := seed.Run(db); err != nil {
		log.Fatal("seed failed:", err)
	}

	clientRepo := repository.NewClientRepository(db)
	invoiceRepo := repository.NewInvoiceRepository(db)
	riskRepo := repository.NewRiskSnapshotRepository(db)
	actionRepo := repository.NewCollectionActionRepository(db)

	clientService := service.NewClientService(clientRepo)
	dashboardService := service.NewDashboardService(clientRepo, invoiceRepo, riskRepo)
	actionService := service.NewCollectionActionService(actionRepo, clientRepo)
	invoiceService := service.NewInvoiceService(db)

	clientHandler := api.NewClientHandler(clientService)
	dashboardHandler := api.NewDashboardHandler(dashboardService)
	actionHandler := api.NewCollectionActionHandler(actionService)
	invoiceHandler := api.NewInvoiceHandler(invoiceService)

	router := api.NewRouter()

	api.RegisterRoutes(
		router,
		api.Handlers{
			ClientHandler:           clientHandler,
			DashboardHandler:        dashboardHandler,
			CollectionActionHandler: actionHandler,
			InvoiceHandler:          invoiceHandler,
		},
	)

	log.Printf("server running on :%s", cfg.AppPort)

	if err := router.Run(":" + cfg.AppPort); err != nil {
		log.Fatal(err)
	}
}
