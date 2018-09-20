package service

import (
	"fmt"
	"encoding/json"
)

const (
	_forbiddenNum = "9991"
)

func (s *Service) SetUserCache(mobile string, redis map[string]string) {
	wy, _ := s.dao.WYExists(mobile, "mobile")
	if wy != nil && wy.SendCount == 1 {
		v, _ := json.Marshal(redis)
		s.dao.SetUserCache(mobile, string(v))
	}
	return
}

func (s *Service) CheckIsYpp(status int8, mobile string) {
	if status > 0 {
		if exists := s.dao.CheckSendCache(mobile); !exists {
			action := "wywk_check_ypp_send"
			hm, _ := s.dao.HMExists(mobile, "mobile")
			if hm == nil {
				s.SendSms(action, mobile)
			} else {
				if hm.CommonCode != _forbiddenNum {
					s.SendSms(action, mobile)
				}
			}
		}
	}
	return
}

func (s *Service) CheckReserve(cardno string) {
	if cardno != "" {
		exists := s.dao.CheckReserveCache(cardno)
		if exists {
			var req = make(map[string]string)
			req["cardno"] = cardno
			res := s.YppApi("WangYuOpenClientCallBack", req)
			fmt.Println(res)
		}
	}
	return
}
