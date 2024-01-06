package database

import (
	"go-email/internal/domain/campaign"
)

type CampaignRepositoryAdapter struct {
	campaigns []campaign.Campaign
}

func (c *CampaignRepositoryAdapter) Save(campaign *campaign.Campaign) error {
	c.campaigns = append(c.campaigns, *campaign)
	return nil
}

func (c *CampaignRepositoryAdapter) Get() []campaign.Campaign {
	return c.campaigns
}
