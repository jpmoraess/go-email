package campaign

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string
}

type Campaign struct {
	ID        string
	Name      string
	Content   string
	Contacts  []Contact
	CreatedOn time.Time
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	if name == "" {
		return nil, errors.New("name is required")
	} else if content == "" {
		return nil, errors.New("content is required")
	} else if len(emails) == 0 {
		return nil, errors.New("contacts is required")
	}

	contacts := make([]Contact, len(emails))
	for index, email := range emails {
		contacts[index].Email = email
	}

	return &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		Contacts:  contacts,
		CreatedOn: time.Now(),
	}, nil
}