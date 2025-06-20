package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/sm888sm/backend-marketplace/internal/config"
	"github.com/sm888sm/backend-marketplace/internal/routes"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Gagal load config: %v", err)
	}

	fmt.Println("DB_USER dari main:", cfg.DBUser)

	db, err := config.InitDB(cfg)
	if err != nil {
		log.Fatalf("Gagal koneksi database: %v", err)
	}
	log.Println("Koneksi database berhasil.")

	router := routes.SetupRouter(cfg, db)

	go func() {
		log.Printf("Server berjalan di port %s...", cfg.ServerPort)
		if err := router.Run(":" + cfg.ServerPort); err != nil {
			log.Fatalf("Gagal menjalankan server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Server dimatikan...")
}
