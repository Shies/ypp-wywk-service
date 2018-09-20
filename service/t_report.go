package service

import (
	"time"
	"ypp-wywk-service/model"
)

const (
	_dateFormat = "2006-01-02"
)

func (s *Service) SendCount(source string) {
	now := time.Now().Format(_dateFormat)
	if count, _ := s.dao.SmsCount(source, now); count > 0 {
		sendCount, _ := s.dao.SmsByOne(source, now)
		s.dao.UpdateReport(int(sendCount)+1, source, now)
		return
	}

	report := &model.TReport{
		CreateDate: time.Now(),
		Source: source,
		SendCount: 1,
	}
	s.dao.AddReport(report)
}
