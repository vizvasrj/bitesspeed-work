package contact

import "time"

type Contact struct {
	ID             int        `json:"id"`
	PhoneNumber    string     `json:"phoneNumber"`
	Email          string     `json:"email"`
	LinkedID       int        `json:"linkedId"`
	LinkPrecedence string     `json:"linkPrecedence"`
	CreatedAt      time.Time  `json:"createdAt"`
	UpdatedAt      time.Time  `json:"updatedAt"`
	DeletedAt      *time.Time `json:"deletedAt,omitempty"`
}

type IdentifyRequest struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

type IdentifyResponse struct {
	Contact struct {
		PrimaryContactID    int      `json:"primaryContactId"`
		Emails              []string `json:"emails"`
		PhoneNumbers        []string `json:"phoneNumbers"`
		SecondaryContactIDs []int    `json:"secondaryContactIds"`
	} `json:"contact"`
}
