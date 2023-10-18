package payloads

import "mime/multipart"

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
