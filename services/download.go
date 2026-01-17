package services

import (
	"VidVendor/config"
	"VidVendor/utils"

	"io"
	"log"
	"os"

	"github.com/kkdai/youtube/v2"
)

var client = youtube.Client{}

func DownloadVideo(cfg *config.Config) error {
	for {
		url, ok := <-URLQueue
		if !ok {
			log.Printf("URLQueue is closed, exiting...")
			return nil
		}
		video, err := client.GetVideo(url)
		if err != nil {
			log.Printf("Failed to upload video: %v", err)
			continue
		}
		formats := video.Formats.WithAudioChannels() // get only formats with audio
		stream, _, err := client.GetStream(video, &formats[0])
		if err != nil {
			log.Printf("Failed to get video stream: %v", err)
			continue
		}
		defer stream.Close()

		videoId := utils.GenerateUUID()
		log.Printf("Downloading video ID: %s\n", videoId)

		videoPath := videoId + ".mp4"
		videoFile, err := os.Create(videoPath)
		if err != nil {
			log.Printf("Failed to create video file: %v", err)
			continue
		}
		defer videoFile.Close()

		_, err = io.Copy(videoFile, stream)
		if err != nil {
			log.Printf("Failed to save video stream: %v", err)
			os.Remove(videoPath)
			continue
		}

		PlaybackQueue <- videoId

		log.Printf("Video %s downloaded successfully\n", videoId)
	}
}
