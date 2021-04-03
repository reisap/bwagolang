package campaign


import "gorm.io/gorm"

type Repository interface {
	//Save(campaign Campaign) (Campaign,error)
	FindByUserID (userID int) ([]Campaign,error)
	FindByAll () ([]Campaign,error)

}

type repository struct{
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository)FindByUserID (userID int) ([]Campaign,error){
	var campaigns []Campaign
	//perload digunakan untuk menunjukan relasi antar table, CampaignImages -> nama di struct dan campaign_images == adalah nama table diikuti nama field asli di database
	err := r.db.Where("user_id = ?",userID).Preload("CampaignImages","campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns,err
	}

	return campaigns,nil

}

func (r *repository)FindByAll () ([]Campaign,error){
	var campaigns []Campaign
	err := r.db.Find(&campaigns).Error
	if err != nil {
		return campaigns,err
	}

	return campaigns,nil
}