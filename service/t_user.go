package service

import (
	"ypp-wywk-service/model"
)

func (s *Service) UserExists(mobilephone string) (user *model.TUser, err error) {
	user, err = s.dao.UserExists(mobilephone)
	return
}