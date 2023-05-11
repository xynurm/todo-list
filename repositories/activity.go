package repositories

import (
	"test/models"

	"gorm.io/gorm"
)

type ActivityRepository interface {
	GetAll() ([]models.Activity, error)
	Save(activity models.Activity) (models.Activity, error)
	FindByID(ID int) (models.Activity, error)
	Update(activity models.Activity) (models.Activity, error)
	Delete(activity models.Activity) (models.Activity, error)
}

func RepositoryActivity(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]models.Activity, error) {
	var activity []models.Activity
	err := r.db.Find(&activity).Error
	return activity, err
}
func (r *repository) Save(activity models.Activity) (models.Activity, error) {
	err := r.db.Create(&activity).Error
	return activity, err
}
func (r *repository) FindByID(ID int) (models.Activity, error) {
	var activity models.Activity

	err := r.db.First(&activity, ID).Error
	return activity, err
}

func (r *repository) Update(activity models.Activity) (models.Activity, error) {
	err := r.db.Save(&activity).Error

	return activity, err
}

func (r *repository) Delete(activity models.Activity) (models.Activity, error) {
	err := r.db.Delete(&activity).Error

	return activity, err
}
