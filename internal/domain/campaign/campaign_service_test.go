package campaign

import (
	"errors"
	"go-email/internal/dto"
	internalerrors "go-email/internal/internal-errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type campaignRepositoryMock struct {
	mock.Mock
}

func (r *campaignRepositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *campaignRepositoryMock) Get() []Campaign {
	return nil
}

var (
	input = dto.NewCampaignInput{
		Name:    "Teste Y",
		Content: "<html></html>",
		Emails:  []string{"teste@mail.com"},
	}

	service = CampaignService{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	campaignRepositoryMock := new(campaignRepositoryMock)
	service.CampaignRepository = campaignRepositoryMock
	campaignRepositoryMock.On("Save", mock.Anything).Return(nil)

	id, err := service.Create(input)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)

	_, err := service.Create(dto.NewCampaignInput{})

	assert.False(errors.Is(internalerrors.ErrInternal, err))
}

func Test_Create_SaveCampaign(t *testing.T) {
	campaignRepositoryMock := new(campaignRepositoryMock)
	service.CampaignRepository = campaignRepositoryMock
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
	campaignRepositoryMock := new(campaignRepositoryMock)
	service.CampaignRepository = campaignRepositoryMock
	campaignRepositoryMock.On("Save", mock.Anything).Return(errors.New("error to save on database"))

	_, err := service.Create(input)

	assert.True(errors.Is(internalerrors.ErrInternal, err))
}
