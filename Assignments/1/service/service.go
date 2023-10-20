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
	payload.Image1.Filename = userID + "_1.jpg"
	if err := s.repos.UpsertImageIntoS3(payload.Image1); err != nil {
		return &payloads.GenericMessageResponse{
			Message: "err while inserting image1 into s3",
		}, err
	}
	payload.Image2.Filename = userID + "_2.jpg"
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
	user := domain.User{}
	if err := s.repos.GetByID(userID, &user); err != nil {
		return err
	}
	//first check if user is already authorized
	if user.State == domain.UserAuthStateAccepted {
		return nil
	}
	user_id := strconv.Itoa(userID)
	//get images urls from s3
	img1 := s.repos.GenerateS3ImageURL(user_id + "_1.jpg")
	img2 := s.repos.GenerateS3ImageURL(user_id + "_2.jpg")
	message := "احزار هویت شما با موفقیت انجام شد"
	//check if face exists in both images
	exists, face1ID, face2ID, err := s.checkFaces(img1, img2)
	if err != nil {
		return err
	}
	if !exists {
		user.State = domain.UserAuthStateDeclined
		if err := s.repos.Upsert(&user); err != nil {
			return err
		}
		message = "احراز هویت شما رد شد"
	}
	//check if images have the same face
	similar, err := s.checkSimilar(face1ID, face2ID)
	if err != nil {
		return err
	}

	if !similar {
		user.State = domain.UserAuthStateDeclined
		if err := s.repos.Upsert(&user); err != nil {
			return err
		}
		message = "احراز هویت شما رد شد"
	} else {
		user.State = domain.UserAuthStateAccepted
		if err := s.repos.Upsert(&user); err != nil {
			return err
		}
	}

	err = s.repos.SendAuthStatusEmail(user.Email, message)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) checkFaces(img1 string, img2 string) (bool, string, string, error) {
	face1, err := s.repos.GetFaceDetection(img1)
	if err != nil {
		return false, "", "", err
	}
	if face1 == "" {
		return false, "", "", nil
	}
	face2, err := s.repos.GetFaceDetection(img2)
	if err != nil {
		return false, "", "", err
	}
	if face2 == "" {
		return false, "", "", nil
	}
	return true, face1, face2, nil
}

func (s *Service) checkSimilar(img1 string, img2 string) (bool, error) {
	score, err := s.repos.GetFaceSimilarity(img1, img2)
	if err != nil {
		return false, err
	}

	if score >= 80 {
		return true, nil
	}
	return false, nil
}
