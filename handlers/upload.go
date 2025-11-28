package handlers

// handles upload requests
import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kkdai/youtube/v2"

	"VidVendor/models"
	"VidVendor/storage"
)

var client = youtube.Client{}

func UploadVideo(c *gin.Context, cfg *models.Config) {
	var req models.UploadVideoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := DownloadYoutubeVideo(req.VideoURL, cfg)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Video uploaded successfully"})
}

func DownloadYoutubeVideo(videoUrl string, cfg *models.Config) error {
	video, err := client.GetVideo(videoUrl)
	if err != nil {
		return fmt.Errorf("failed to upload video: %v", err)
	}
	formats := video.Formats.WithAudioChannels() // get only formats with audio
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		return fmt.Errorf("failed to get video streams: %v", err)
	}
	defer stream.Close()

	// store the video file in cloud storage
	videoId := uuid.New().String()
	fmt.Printf("Downloading video ID: %s\n", videoId)

	videoPath := videoId + ".mp4"
	videoFile, err := os.Create(videoPath)
	if err != nil {
		return fmt.Errorf("failed to create video file: %v", err)
	}
	defer videoFile.Close()
	defer os.Remove(videoPath)

	_, err = io.Copy(videoFile, stream)
	if err != nil {
		return fmt.Errorf("failed to save video stream: %v", err)
	}

	err = storage.UploadToGCS(videoPath, cfg)
	if err != nil {
		return fmt.Errorf("failed to upload to cloud storage: %v", err)
	}

	fmt.Printf("Video %s downloaded successfully\n", videoId)
	return nil
}
