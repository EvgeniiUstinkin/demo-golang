package response

import (
	"time"

	sresponse "github.com/finchatapp/finchat-api/internal/entities/_shared/models/response"
)

type User struct {
	ID            int        `json:"id"`
	UUID          string     `json:"uuid"`
	IsActive      bool       `json:"isActive"`
	FirebaseID    *string    `json:"-"`
	StripeID      *string    `json:"-"`
	FirstName     string     `json:"firstName"`
	LastName      string     `json:"lastName"`
	Phonenumber   string     `json:"phoneNumber"`
	CountryCode   string     `json:"countryCode"`
	Email         string     `json:"email"`
	Username      *string    `json:"username,omitempty"`
	IsVerified    bool       `json:"isVerified"`
	Type          string     `json:"userType"`
	ProfileAvatar *string    `json:"profileAvatar"`
	LastSeen      time.Time  `json:"lastSeen"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
	DeletedAt     *time.Time `json:"deletedAt,omitempty"`
}

type UserList struct {
	Paging sresponse.Paging `json:"paging"`
	Items  []User           `json:"items"`
}
