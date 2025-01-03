package campaign

import (
	"go-email/internal/dto"
	internalerrors "go-email/internal/internal-errors"
)

type CampaignService struct {
	CampaignRepository CampaignRepository
}

func (cs *CampaignService) Create(input dto.NewCampaignInput) (string, error) {
	campaign, err := NewCampaign(input.Name, input.Content, input.Emails)
	if err != nil {
		return "", err
	}
	err = cs.CampaignRepository.Save(campaign)
	if err != nil {
		return "", internalerrors.ErrInternal
	}
	return campaign.ID, nil
}
