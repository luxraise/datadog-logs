package logs

func MakeSource(hostname, source string) (s Source) {
	s.Hostname = hostname
	s.Source = source
	return
}

type Source struct {
	logCore
}

func (s *Source) NewService(service string) *Service {
	var svc Service
	svc.logCore = s.logCore
	svc.Service = service
	return &svc
}
