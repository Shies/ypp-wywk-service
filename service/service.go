package service

import (
	"sync"
	"ypp-wywk-service/conf"
	"ypp-wywk-service/dao"
)


// Service biz service def.
type Service struct {
	c    *conf.Config
	dao	 *dao.Dao
	wait *sync.WaitGroup
	once sync.Once
}

// New new a Service and return.
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:    c,
		dao:  dao.New(c),
		wait: new(sync.WaitGroup),
	}

	return s
}


func (s *Service) Close() {
}

// Ping check server ok.
func (s *Service) Ping() (err error) {
	return
}