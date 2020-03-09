package models

import "github.com/jinzhu/gorm"

type AdminAuth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func Admin_CheckAuth(username, password string) (bool, error) {
	var adminAuth AdminAuth
	err := db.Select("id").Where(AdminAuth{Username: username, Password: password}).First(&adminAuth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if adminAuth.ID > 0 {
		return true, nil
	}

	return false, nil
}
