package payloads

type SignUpRequest struct {
	Email      string `json:"email"`
	LastName   string `json:"last_name"`
	NationalID int    `json:"national_id"`
	Image1     string `json:"image1"`
	Image2     string `json:"image2"`
}

type CheckStatusRequest struct {
	NationalID int `json:"national_id"`
}

type GenericMessageResponse struct {
	Message string `json:"message"`
}
