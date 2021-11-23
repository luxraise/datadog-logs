package logs

func MakeSource(hostname, source string) (s Source) {
	s.Hostname = hostname
	s.Source = source
	return
}

type Source struct {
	logCore
}

func (s *Source) MakeService(service string) (svc Service) {
	svc.logCore = s.logCore
	svc.Service = service
	return
}
