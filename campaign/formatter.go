package campaign

type CampaignFormatter struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	Name string `json:"name"`
	ShortDescription string `json:"short_desription"`
	CurrentAmount int `json:"current_amount"`
	GoalAmount int `json:"goal_amount"`
	ImageUrl string `json:"image_url"`
	Slug string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter{
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.Name = campaign.Name
	campaignFormatter.UserID = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.ImageUrl = ""
	campaignFormatter.Slug = campaign.Slug

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}
	return campaignFormatter

}

func FormatCampaigns(campaign []Campaign)[]CampaignFormatter{
	if len(campaign) == 0 {
		return []CampaignFormatter{}
	}
	//var campaignsFormater []CampaignFormatter
	campaignsFormater := []CampaignFormatter{}

	for _, data := range(campaign) {
		singleCampaignFormater := FormatCampaign(data)
		campaignsFormater = append(campaignsFormater,singleCampaignFormater)
	}
	return campaignsFormater
}