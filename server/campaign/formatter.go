package campaign

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	GoalAMount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	CampaignFormatter := CampaignFormatter{}
	CampaignFormatter.ID = campaign.ID
	CampaignFormatter.UserID = campaign.UserID
	CampaignFormatter.Name = campaign.Name
	CampaignFormatter.ShortDescription = campaign.ShortDescription
	CampaignFormatter.GoalAMount = campaign.GoalAMount
	CampaignFormatter.CurrentAmount = campaign.CurrentAmount

	CampaignFormatter.ImageUrl = ""
	if len(campaign.CampaignImages) > 0 {
		CampaignFormatter.ImageUrl = campaign.CampaignImages[0].Filename
	}
	return CampaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	var campaignsFormatter []CampaignFormatter

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}
	return campaignsFormatter
}
