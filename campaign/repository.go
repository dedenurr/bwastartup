package campaign

import "gorm.io/gorm"

// Definisi kontrak interface repository
type Repository interface{
	FindAll () ([]Campaign, error)
	FindByUserId(userID int) ([]Campaign, error)
}

// buat struct repository untuk akses ke Database
type repository struct{
	db *gorm.DB
}

// agar bisa di aksese diluar package struct diatas
func NewRepository(db *gorm.DB) *repository  {
	return &repository{db}	
}

func (r *repository) FindAll() ([]Campaign, error)  {
	// definisikan data yang mau dicari
	var campaigns []Campaign

	// query db gorm
	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	
	// jika terjadi error
	if err != nil {
		return campaigns, err
	}

	// jika tidak ada error 
	return campaigns, nil
}

func (r * repository) FindByUserId(userID int) ([]Campaign, error)  {
	// definisikan data yang mau dicari
	var campaigns []Campaign

	// query db gorm find by ID
	err := r.db.Where("user_id = ?", userID).Preload("CampaignImages","campaign_images.is_primary = 1").Find(&campaigns).Error

	// jika terjadi error
	if err != nil {
		return campaigns, err
	}

	// jika tidak ada error 
	return campaigns, nil
}