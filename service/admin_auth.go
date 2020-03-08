package service

import "meeting/models"

type AdminAuth struct {
	Username string
	Password string
}

func (a *AdminAuth) Check() (bool, error) {
	return models.CheckAdminAuth(a.Username, a.Password)
}
