package models

type AdminUserInfo struct {
	ID           int    `gorm:"primary_key" json:"id"`
	Name         string `json:"name"`
	Avatar       string `json:"avatar"`
	Introduction string `json:"introduction"`
	Role         string `json:"role"`
}

func Admin_GetUserInfo(id int) (*AdminUserInfo, error) {
	var userInfo AdminUserInfo
	err := db.Model(&AdminUserInfo{}).Where("id = ?", id).First(&userInfo).Error
	if err != nil {
		return nil, err
	}
	return &userInfo, nil
}
