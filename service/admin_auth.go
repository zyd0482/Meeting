package service

import "meeting/models"

type AdminAuth struct {
	Username string
	Password string
}

func (a *AdminAuth) Check() (bool, error) {
	return models.Admin_CheckAuth(a.Username, a.Password)
}
