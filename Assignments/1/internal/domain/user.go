package domain

type User struct {
	ID         int    `json:"id"`
	Email      string `json:"email"`
	LastName   string `json:"last_name"`
	NationalID int    `json:"national_id"`
	IP         string `json:"ip"`
	State      string `json:"state"`
	Image1     string `json:"image1"`
	Image2     string `json:"image2"`
}
