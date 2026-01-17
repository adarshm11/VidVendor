package main

import (
	"flag"
	"log"
	"net/http"

	"VidVendor/config"
	"VidVendor/handlers"
)

func main() {
	configPath := flag.String("config", "config.yml", "Path to the config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatalf("Error loading config: %v\n", err)
	}

	_ = cfg // Use config as needed

	r := http.NewServeMux()

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "pong"}`))
	})

	r.HandleFunc("/upload", handlers.UploadVideoHandler)
	r.HandleFunc("/next", handlers.GetNextVideoHandler)
	r.HandleFunc("/stop", handlers.StopStreamHandler)

	log.Println("server running on port", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}
