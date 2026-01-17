package handlers

import (
	"encoding/json"
	"net/http"

	"VidVendor/models"
	"VidVendor/services"
)

func UploadVideoHandler(w http.ResponseWriter, r *http.Request) {
	ct := r.Header.Get("Content-Type")
	if ct != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}
	var req models.UploadVideoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	uuid, err := services.AddVideoURL(req.URL)
	if err != nil {
		http.Error(w, "Failed to add video URL", http.StatusInternalServerError)
		return
	}
	err = services.EnqueueVideo(uuid)
	if err != nil {
		http.Error(w, "Failed to enqueue video", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetNextVideoHandler(w http.ResponseWriter, r *http.Request) {

}

func StopStreamHandler(w http.ResponseWriter, r *http.Request) {
	services.EndStream()
	w.WriteHeader(http.StatusOK)
}
