package models

type Config struct {
	Port                 string `yaml:"port"`
	GoogleCloudProjectID string `yaml:"google_cloud_project_id"`
	BucketName           string `yaml:"bucket_name"`
}
