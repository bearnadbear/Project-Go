package reposerviceUser

import (
	modelUser "project/model/user"

	"gorm.io/gorm"
)

type Repository interface {
	Save(user modelUser.User) (modelUser.User, error)
	FindByEmail(email string) (modelUser.User, error)
	FindByID(ID int) (modelUser.User, error)
	Update(user modelUser.User) (modelUser.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user modelUser.User) (modelUser.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (modelUser.User, error) {
	var user modelUser.User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByID(ID int) (modelUser.User, error) {
	var user modelUser.User

	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Update(user modelUser.User) (modelUser.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
