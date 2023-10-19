package domain

type User struct {
	ID         int           `json:"id"`
	Email      string        `json:"email"`
	LastName   string        `json:"last_name"`
	NationalID string        `json:"national_id"`
	IP         string        `json:"ip"`
	State      UserAuthState `json:"state"`
	Image1     string        `json:"image1"`
	Image2     string        `json:"image2"`
}

type UserAuthState string

const (
	UserAuthStatePending  UserAuthState = "pending"
	UserAuthStateDeclined UserAuthState = "declined"
	UserAuthStateAccepted UserAuthState = "accepted"
)
