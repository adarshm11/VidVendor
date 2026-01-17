package services

// goroutine to read from UUID and serve videos back to the handler

import (
	"log"

	"VidVendor/config"
)

func PlayVideo(cfg *config.Config) {
	for {
		uuid, ok := <-PlaybackQueue
		if !ok {
			log.Printf("PlayQueue is closed, exiting...")
			return
		}

		log.Printf("Playing video with UUID: %s\n", uuid)
		// TODO: extract video document + MP4 from GCS and serve it back to the handler
	}
}
