package repository

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/config"
	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/domain"
	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	amqp "github.com/rabbitmq/amqp091-go"
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
	s3Client, err := config.GenerateS3Client()
	if err != nil {
		return err
	}
	bucket := aws.String(config.GetConfig().S3.BucketName)
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
	objectAcl := &s3.PutObjectAclInput{
		Bucket: bucket,
		Key:    aws.String(img.Filename),
		ACL:    aws.String("public-read"),
	} // set access to public
	_, err = s3Client.PutObjectAcl(objectAcl)
	if err != nil {
		logger.Logger().Errorw("error while updating acl", "error", err)
	}
	return nil
	//checkhttps://docs.arvancloud.ir/fa/developer-tools/sdk/object-storage/

}

func (r *Repository) GetImageFromS3(objName string) (*os.File, error) {
	s3Client, err := config.GenerateS3Client()
	if err != nil {
		return nil, err
	}
	s3Obj := &s3.GetObjectInput{
		Bucket: aws.String(config.GetConfig().S3.BucketName),
		Key:    aws.String(objName),
	}
	output, err := s3Client.GetObject(s3Obj)
	if err != nil {
		return nil, err
	}
	img, err := os.Create(objName)
	defer img.Close()
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(img, output.Body); err != nil {
		return nil, err
	}
	return img, nil
}

func (r *Repository) GenerateS3ImageURL(objName string) string {
	s3Config := config.GetConfig().S3
	url := fmt.Sprintf("%s/%s/%s", s3Config.Endpoint, s3Config.BucketName, objName)
	return url
}

func (r *Repository) PublishToRabbitMQ(message string) error {
	rabbMQ := config.GetConfig().RabbitMQ
	conn, err := amqp.Dial(rabbMQ.GetURI())
	if err != nil {
		logger.Logger().Debugw("rabbitMQ Connection err", "error", err)
		return err
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	queue, err := ch.QueueDeclare(rabbMQ.QueueName, false, false, false, false, nil)
	if err != nil {
		return err
	}

	err = ch.PublishWithContext(context.TODO(),
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		return err
	}
	return nil

}
