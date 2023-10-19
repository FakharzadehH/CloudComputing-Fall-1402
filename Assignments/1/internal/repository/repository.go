package repository

import (
	"mime/multipart"

	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/config"
	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/domain"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Upsert(user *domain.User) error {
	return r.db.Save(user).Error
}

func (r *Repository) GetByID(id int, user *domain.User) error {
	return r.db.First(user, id).Error
}

func (r *Repository) GetByNationalID(nationalID string, user *domain.User) error {

	return r.db.First(user, "national_id = ?", nationalID).Error
}

func (r *Repository) UpsertImageIntoS3(img *multipart.FileHeader) error {
	cfg := config.GetConfig()
	s3Session, err := session.NewSessionWithOptions(session.Options{
		Config:  cfg.S3.GenerateS3Config(),
		Profile: "filebase",
	})
	if err != nil {
		return err
	}
	s3Client := s3.New(s3Session)
	bucket := aws.String(cfg.S3.BucketName)
	file, err := img.Open()
	if err != nil {
		return err
	}
	defer file.Close()
	s3Input := &s3.PutObjectInput{
		Body:   file,
		Bucket: bucket,
		Key:    aws.String(img.Filename),
	}
	if _, err := s3Client.PutObject(s3Input); err != nil {
		return err
	}
	return nil
	//check https://docs.filebase.com/code-development-+-sdks/code-development/aws-sdk-go-golang

}
