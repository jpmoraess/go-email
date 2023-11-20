package campaign

import (
	"errors"
	"go-email/internal/dto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type CampaignRepositoryMock struct {
	mock.Mock
}

func (r *CampaignRepositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

var (
	input = dto.NewCampaignInput{
		Name:    "Teste Y",
		Content: "<html></html>",
		Emails:  []string{"teste@mail.com"},
	}

	campaignRepositoryMock = new(CampaignRepositoryMock)

	service = CampaignService{
		CampaignRepository: campaignRepositoryMock,
	}
)

/*
func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	id, err := service.Create(input)

	assert.NotNil(id)
	assert.Nil(err)
}
*/

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)
	input.Name = ""

	_, err := service.Create(input)

	assert.NotNil(err)
	assert.Equal("name is required", err.Error())
}

func Test_Create_SaveCampaign(t *testing.T) {
	campaignRepositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != input.Name || campaign.Content != input.Content || len(campaign.Contacts) != len(input.Emails) {
			return false
		}
		return true
	})).Return(nil)

	service.Create(input)

	campaignRepositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	campaignRepositoryMock.On("Save", mock.Anything).Return(errors.New("error to save on database"))

	_, err := service.Create(input)

	assert.Equal("error to save on database", err.Error())
}
