package campaign

import "time"

type Campaign struct {
	ID int
	UserID int
	Name string
	ShortDescription string
	Description string
	BackerCount int
	CurrentAmount int
	GoalAmount int
	Perks string
	Slug string
	CreatedAt time.Time
	UpdatedAt time.Time
	CampaignImages []CampaignImages
}

type CampaignImages struct {
	ID int
	CampaignID int
	FileName string
	IsPrimary string
	CreatedAt time.Time
	UpdatedAt time.Time

}