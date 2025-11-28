package main

import (
	"flag"
	"log"

	"github.com/gin-gonic/gin"

	"VidVendor/config"
	"VidVendor/routes"
)

func main() {
	configPath := flag.String("config", "config.yml", "Path to the config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatalf("Error loading config: %v\n", err)
	}

	_ = cfg // Use config as needed

	r := gin.Default()
	routes.RegisterRoutes(r)

	log.Println("server running on port", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}
