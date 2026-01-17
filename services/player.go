package services

import (
	"log"
	"os"

	"VidVendor/config"
	"VidVendor/utils"
)

func AddVideoURL(url string) (string, error) {
	uuid := utils.GenerateUUID()
	URLQueue <- url
	log.Printf("Video URL added to queue: %s with ID: %s\n", url, uuid)
	return uuid, nil
}

func EnqueueVideo(uuid string) error {
	PlaybackQueue <- uuid
	log.Printf("Video ID enqueued for playback: %s\n", uuid)
	return nil
}

// Called when the user skips the current video or the current video ends
func GetNextVideo() string {
	uuid, ok := <-PlaybackQueue
	if !ok {
		log.Printf("PlaybackQueue is closed, exiting...")
		return ""
	}

	log.Printf("Next video ID: %s\n", uuid)
	return uuid
}

// Goroutine that continuously listens to the DeletionQueue and deletes videos
func VideoCleanup(cfg *config.Config) {
	for uuid := range DeletionQueue {
		videoPath := cfg.OutputDirectory + "/" + uuid + ".mp4"
		if err := os.Remove(videoPath); err != nil {
			log.Printf("Failed to delete video %s: %v", uuid, err)
		} else {
			log.Printf("Video %s deleted successfully", uuid)
		}
	}
}

// Empties the PlaybackQueue and schedules all videos for deletion
func EndStream() {
	for len(PlaybackQueue) > 0 {
		DeletionQueue <- <-PlaybackQueue
	}
}
