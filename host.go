package logs

func MakeHost(hostname string, tags ...Tag) (h Host) {
	h.Hostname = hostname
	h.Tags = tags
	return
}

type Host struct {
	logCore
}

func (h *Host) MakeService(service string, tags ...Tag) (s Service) {
	s.logCore = h.copy()
	s.Service = service
	s.Tags = append(s.Tags, tags...)
	return
}
