package service

import (
	"fmt"
	"ypp-wywk-service/submail"
	"ypp-wywk-service/model"
	"time"
)

const (
	_appID = "10868"
	_appKey = "b6c87b82ccbb62d74a2a3e94734c20b1"
	_signType = "md5"
)

func (s *Service) SaveRegister(data *model.TRegister) (err error) {
	err = s.dao.AddWYWKRegister(data)
	return
}

func (s *Service) SaveWYRegister(data *model.TWyRegister) (err error) {
	isset, _ := s.dao.WYExists(data.IDCard, "id_card")
	if isset == nil {
		insert := &model.TWyRegister{
			ID: UUid(),
			TypeID: data.TypeID,
			IDType: data.IDType,
			IDCard: data.IDCard,
			Truename: data.Truename,
			CreateTime: data.CreateTime,
			Mobile: data.Mobile,
			CardNO: data.CardNO,
			Status: data.Status,
			ActiveTime: data.ActiveTime,
			CardCreateTime: data.CardCreateTime,
			CommonCode: data.CommonCode,
		}
		if IsMobile(data.Mobile) {
			insert.SendCount = 1
			insert.LastSendDate = time.Now()
		}
		return s.dao.AddWYRegister(insert)
	} else {
		if !IsMobile(isset.Mobile) && IsMobile(data.Mobile) {
			user, _ := s.dao.WYExists(isset.ID, "id")
			if user != nil {
				user.Mobile = data.Mobile
				user.MobileUpdateTime = time.Now()
				user.SendCount = 1
				return s.dao.UpdateWYRegister(user)
			}
		}
		lastSendTime := isset.LastSendDate.Unix()
		bTime := ConvertTime("2016-08-04").Unix()
		begin := lastSendTime < bTime && time.Now().Unix() >= bTime
		if IsMobile(isset.Mobile) && isset.Status == 0 && begin {
			user, _ := s.dao.WYExists(isset.ID, "id")
			if user != nil {
				user.SendCount = isset.SendCount + 1
				user.LastSendDate = time.Now()
				return s.dao.UpdateWYRegister(user)
			}
		}
	}

	return nil
}

func (s *Service) SaveHMRegister(data *model.TWyRegister) (err error) {
	isset, _ := s.dao.HMExists(data.IDCard, "id_card")
	if isset == nil {
		insert := &model.TWyRegister{
			ID: UUid(),
			TypeID: data.TypeID,
			IDType: data.IDType,
			IDCard: data.IDCard,
			Truename: data.Truename,
			CreateTime: data.CreateTime,
			Mobile: data.Mobile,
			CardNO: data.CardNO,
			Status: data.Status,
			ActiveTime: data.ActiveTime,
			CardCreateTime: data.CardCreateTime,
			CommonCode: data.CommonCode,
		}
		if IsMobile(data.Mobile) {
			insert.SendCount = 1
			insert.LastSendDate = time.Now()
		}
		return s.dao.AddHMRegister(insert)
	} else {
		if !IsMobile(isset.Mobile) && IsMobile(data.Mobile) {
			user, _ := s.dao.HMExists(isset.ID, "id")
			if user != nil {
				user.Mobile = data.Mobile
				user.MobileUpdateTime = time.Now()
				user.SendCount = 1
				return s.dao.UpdateHMRegister(user)
			}
		}
		lastSendTime := isset.LastSendDate.Unix()
		bTime := ConvertTime("2016-08-04").Unix()
		begin := lastSendTime < bTime && time.Now().Unix() >= bTime
		if IsMobile(isset.Mobile) && isset.Status == 0 && begin {
			user, _ := s.dao.HMExists(isset.ID, "id")
			if user != nil {
				user.SendCount = isset.SendCount + 1
				user.LastSendDate = time.Now()
				return s.dao.UpdateHMRegister(user)
			}
		}
	}

	return nil
}

func (s *Service) SendSms(action string, mobilephone string) bool {
	messageconfig := make(map[string]string)
	messageconfig["appid"] = _appID
	messageconfig["appkey"] = _appKey
	messageconfig["signtype"] = _signType

	switch action {
	case "anyRegisterNew":
		// messagexsend
		messagexsend := submail.CreateMessageXSend()
		submail.MessageXSendAddTo(messagexsend, mobilephone)
		submail.MessageXSendSetProject(messagexsend, "EWBaR")
		submail.MessageXSendAddVar(messagexsend, "download_url", "t.cn/RONMqYW")
		fmt.Println("MessageXSend ", submail.MessageXSendRun(submail.MessageXSendBuildRequest(messagexsend), messageconfig))
	case "anyRegisterOld":
		// messagexsend
		messagexsend := submail.CreateMessageXSend()
		submail.MessageXSendAddTo(messagexsend, mobilephone)
		submail.MessageXSendSetProject(messagexsend, "7Fbam")
		submail.MessageXSendAddVar(messagexsend, "download_url", "t.cn/ROg4xWu")
		fmt.Println("MessageXSend ", submail.MessageXSendRun(submail.MessageXSendBuildRequest(messagexsend), messageconfig))
	}

	return true
}