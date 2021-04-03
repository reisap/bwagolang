package campaign

type Service interface {

	GetCampaigns(userID int) ([]Campaign,error)
	GetCampaignById(input GetCampaignDetailInput)(Campaign,error)
}

type campaignService struct {
	repository Repository
}

func NewService(repository Repository)*campaignService{
	return &campaignService{repository}
}


func (s *campaignService)GetCampaigns(userID int) ([]Campaign,error){

	if userID != 0 {
		campaign,err := s.repository.FindByUserID(userID)
		if err != nil {
			return campaign,err
		}

		return campaign,nil
	}

	campaign,err := s.repository.FindByAll()
	if err != nil {
		return campaign,err
	}

	return campaign,nil


}

func (s *campaignService)GetCampaignById(input GetCampaignDetailInput)(Campaign,error){

	data,err := s.repository.FindById(input.ID)
	if err != nil {
		return data,err
	}

	return data,nil
}