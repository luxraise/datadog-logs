package logs

func MakeSource(hostname, source string, tags ...Tag) (s Source) {
	s.Hostname = hostname
	s.Source = source
	s.Tags = tags
	return
}

type Source struct {
	logCore
}

func (s *Source) MakeService(service string) (svc Service) {
	svc.logCore = s.copy()
	svc.Service = service
	return
}
