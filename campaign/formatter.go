package campaign

import "strings"

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

type UserCampaignFormatter struct {
	Name string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type ImagesCampaignFormatter struct {
	ImageUrl string `json:"image_url"`
	IsPrimary bool  `json:"is_primary"`
}

type DetailCampaignFormatter struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	Name string `json:"name"`
	ShortDescription string `json:"short_desription"`
	Description string `json:"description"`
	ImageUrl string `json:"image_url"`
	BackerCount int `json:"backer_count"`
	CurrentAmount int `json:"current_amount"`
	GoalAmount int `json:"goal_amount"`
	Perks []string `json:"perks"`
	Slug string `json:"slug"`
	CampaignImages []ImagesCampaignFormatter `json:"images"`
	User UserCampaignFormatter `json:"user"`
}

func FormatDetailCampaign(campaign Campaign)DetailCampaignFormatter{
	campaignFormatter := DetailCampaignFormatter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.UserID = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.ImageUrl = ""
	campaignFormatter.Slug = campaign.Slug

	var stringSplit []string

	for _,perk := range(strings.Split(campaign.Perks, ",")){
		stringSplit = append(stringSplit,strings.TrimSpace(perk))
	}
	campaignFormatter.Perks = stringSplit
	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	user := campaign.User
	campaignUser := UserCampaignFormatter{}
	campaignUser.Name = user.Name
	campaignUser.ImageUrl = user.AvatarFileName
	campaignFormatter.User = campaignUser


	arrImages := []ImagesCampaignFormatter{}

	for _,data := range(campaign.CampaignImages){
		imagesCampaignFormatter := ImagesCampaignFormatter{}
		imagesCampaignFormatter.ImageUrl = data.FileName
		isPrimary := false
		if data.IsPrimary == 1 {
			isPrimary = true
		}
		imagesCampaignFormatter.IsPrimary = isPrimary
		arrImages = append(arrImages,imagesCampaignFormatter)

	}

	campaignFormatter.CampaignImages = arrImages



	return campaignFormatter
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