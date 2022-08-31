package logs

func MakeService(hostname, service string, tags ...Tag) (s Service) {
	s.Hostname = hostname
	s.Service = service
	s.Tags = tags
	return
}

type Service struct {
	logCore
}
