package campaign

type CampaignRepository interface {
	Save(campaign *Campaign) error
}
