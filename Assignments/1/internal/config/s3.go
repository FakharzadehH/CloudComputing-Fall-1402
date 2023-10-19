package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

type S3 struct {
	Region     string `mapstructure:"region"`
	AccessKey  string `mapstructure:"access_key"`
	SecretKey  string `mapstructure:"secret_key"`
	Endpoint   string `mapstructure:"url"`
	BucketName string `mapstructure:"bucket_name"`
}

func (s3 S3) GenerateS3Config() aws.Config {
	s3config := aws.Config{
		Credentials:      credentials.NewStaticCredentials(s3.AccessKey, s3.SecretKey, ""),
		Endpoint:         aws.String(s3.Endpoint),
		Region:           aws.String(s3.Region),
		S3ForcePathStyle: aws.Bool(true),
	}
	return s3config
}
