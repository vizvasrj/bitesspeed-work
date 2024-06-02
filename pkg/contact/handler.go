package contact

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Identify(c *gin.Context) error {
	var req IdentifyRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}

	contact, err := getOrCreateContact(db, req.Email, req.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}

	response := IdentifyResponse{
		Contact: struct {
			PrimaryContactID    int      `json:"primaryContactId"`
			Emails              []string `json:"emails"`
			PhoneNumbers        []string `json:"phoneNumbers"`
			SecondaryContactIDs []int    `json:"secondaryContactIds"`
		}{
			PrimaryContactID:    contact.ID,
			Emails:              []string{contact.Email},
			PhoneNumbers:        []string{contact.PhoneNumber},
			SecondaryContactIDs: []int{},
		},
	}

	// Fetch secondary contacts
	secondaryContacts, err := getSecondaryContacts(db, contact.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}

	for _, sc := range secondaryContacts {
		response.Contact.Emails = append(response.Contact.Emails, sc.Email)
		response.Contact.PhoneNumbers = append(response.Contact.PhoneNumbers, sc.PhoneNumber)
		response.Contact.SecondaryContactIDs = append(response.Contact.SecondaryContactIDs, sc.ID)
	}

	c.JSON(http.StatusOK, response)
	return nil
}
