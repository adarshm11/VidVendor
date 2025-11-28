package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"

	"VidVendor/handlers"
	"VidVendor/models"
)

var cfg models.Config

func LoadConfig(path string) (*models.Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func main() {
	configPath := flag.String("config", "config.yaml", "Path to the config file")
	flag.Parse()

	cfg, err := LoadConfig(*configPath)
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	_ = cfg // Use config as needed

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/upload", func(c *gin.Context) {
		handlers.UploadVideo(c, cfg)
	})

	fmt.Println("server running on port", cfg.Port)
	r.Run(":" + cfg.Port)
}
