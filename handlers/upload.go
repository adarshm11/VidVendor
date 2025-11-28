package handlers

// handles upload requests
import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kkdai/youtube/v2"

	"VidVendor/models"
)

var client = youtube.Client{}

func UploadVideo(c *gin.Context, outputDir string) {
	var req models.UploadVideoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := DownloadYoutubeVideo(req.VideoURL, req.OutputPath)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Video uploaded successfully"})
}

func DownloadYoutubeVideo(videoUrl, outputPath string) error {
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

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, stream) // Save the stream to the output file
	if err != nil {
		return fmt.Errorf("failed to save video stream: %v", err)
	}

	fmt.Printf("Video uploaded successfully: %s\n", video.ID)
	return nil
}
