package logs

func MakeService(hostname, source, service string) (s Service) {
	s.Hostname = hostname
	s.Source = source
	s.Service = service
	return
}

type Service struct {
	logCore
}
