package repositories

import (
	"github.com/code-guy21/PingUp/server/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func CreateUser(user *models.User) error {
	return DB.Create(user).Error
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}