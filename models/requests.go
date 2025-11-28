package models

type UploadVideoRequest struct {
	UserID   string `json:"user_id" binding:"required"`
	VideoURL string `json:"video_url" binding:"required"`
}
