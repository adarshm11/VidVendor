package services

import (
	"VidVendor/config"
)

var (
	URLQueue      chan string
	PlaybackQueue chan string
)

func InitQueues(cfg *config.Config) {
	URLQueue = make(chan string, cfg.URLQueueBufferSize)
	PlaybackQueue = make(chan string, cfg.PlaybackQueueBufferSize)
}
