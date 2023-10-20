package service

import (
	"context"
	"strconv"

	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/domain"
	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/domain/payloads"
	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repos *repository.Repository
}

func New(repos *repository.Repository) *Service {
	return &Service{
		repos: repos,
	}
}

func (s *Service) SubmitRequest(ctx context.Context, payload payloads.SignUpRequest, ip string) (*payloads.GenericMessageResponse, error) {
	hashedNationalID, err := bcrypt.GenerateFromPassword([]byte(payload.NationalID), bcrypt.DefaultCost)
	if err != nil {
		return &payloads.GenericMessageResponse{
			Message: "err while hashing national id",
		}, err
	}

	user := domain.User{
		Email:      payload.Email,
		LastName:   payload.LastName,
		NationalID: string(hashedNationalID),
		IP:         ip,
		State:      domain.UserAuthStatePending,
	}
	if err := s.repos.Upsert(&user); err != nil {
		return &payloads.GenericMessageResponse{
			Message: "err while inserting request to db",
		}, err
	}
	userID := strconv.Itoa(user.ID)

	payload.Image1.Filename = userID + "_1"
	if err := s.repos.UpsertImageIntoS3(payload.Image1); err != nil {
		return &payloads.GenericMessageResponse{
			Message: "err while inserting image1 into s3",
		}, err
	}
	payload.Image2.Filename = userID + "_2"
	if err := s.repos.UpsertImageIntoS3(payload.Image2); err != nil {
		return &payloads.GenericMessageResponse{
			Message: "err while inserting image2 into s3",
		}, err
	}
	user.Image1 = payload.Image1.Filename
	user.Image2 = payload.Image2.Filename
	if err := s.repos.Upsert(&user); err != nil {
		return &payloads.GenericMessageResponse{
			Message: "err while updating user images in db",
		}, err
	}
	if err := s.repos.PublishToRabbitMQ(userID); err != nil {
		return &payloads.GenericMessageResponse{
			Message: "err while inserting user id to RabbitMQ",
		}, err
	}

	return &payloads.GenericMessageResponse{
		Message: "درخواست احراز هویت شما ثبت شد",
	}, nil
}

func (s *Service) ProccessRequest(userID int) error {
	//get user from db, get image names, get images from S3, check for faces, check for similiarity, send email if successfull
	return nil
}
