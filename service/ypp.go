package service

func (s *Service) YppApi(method string, request map[string]string) string {
	return s.dao.YppApi(method, request)
}