package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"VidVendor/config"
	"VidVendor/handlers"
	"VidVendor/services"
)

func main() {
	configPath := flag.String("config", "config.yml", "Path to the config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatalf("Error loading config: %v\n", err)
	}

	r := http.NewServeMux()

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "pong"}`))
	})

	r.HandleFunc("/upload", handlers.UploadVideoHandler)
	r.HandleFunc("/next", handlers.GetNextVideoHandler)
	r.HandleFunc("/stop", handlers.StopStreamHandler)

	services.InitQueues(cfg)

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt, syscall.SIGTERM)

	go services.DownloadVideo(cfg, sigchan)
	go services.VideoCleanup(cfg, sigchan)

	log.Println("server running on port", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatalf("Error starting server: %v\n", err)
		services.EndStream()
		sigchan <- os.Interrupt
		close(services.DeletionQueue)
		close(services.PlaybackQueue)
		close(services.URLQueue)
	}
}
