package models

type UploadVideoRequest struct {
	VideoURL   string `json:"video_url" binding:"required"`
	OutputPath string `json:"output_path" binding:"required"`
}
