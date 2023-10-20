package payloads

import (
	"mime/multipart"

	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/domain"
)

type SignUpRequest struct {
	Email      string                `json:"email"`
	LastName   string                `json:"last_name"`
	NationalID string                `json:"national_id"`
	Image1     *multipart.FileHeader `json:"image1"`
	Image2     *multipart.FileHeader `json:"image2"`
}
type CheckStatusRequest struct {
	NationalID string `json:"national_id"`
}

type GenericMessageResponse struct {
	Message string `json:"message"`
}

type FaceDetectionResponse struct {
	Result domain.FaceDetectionResult `json:"result"`
}

type FaceSimilarityResponse struct {
	Result domain.FaceSimilarityResult `json:"result"`
}
