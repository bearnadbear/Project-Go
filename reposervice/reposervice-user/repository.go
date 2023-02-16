package reposerviceuser

import (
	"project/model/user"

	"gorm.io/gorm"
)

type Repository interface {
	Save(user user.User) (user.User, error)
	FindByEmail(email string) (user.User, error)
	FindByID(ID int) (user.User, error)
	Update(user user.User) (user.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user user.User) (user.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (user.User, error) {
	var user user.User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByID(ID int) (user.User, error) {
	var user user.User

	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Update(user user.User) (user.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
