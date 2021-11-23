package logs

func MakeService(hostname, source, service string, tags ...Tag) (s Service) {
	s.Hostname = hostname
	s.Source = source
	s.Service = service
	s.Tags = tags
	return
}

type Service struct {
	logCore
}
