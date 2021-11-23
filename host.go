package logs

func MakeHost(hostname string, tags ...Tag) (h Host) {
	h.Hostname = hostname
	h.Tags = tags
	return
}

type Host struct {
	logCore
}

func (h *Host) MakeSource(source string) (s Source) {
	s.logCore = h.copy()
	s.Source = source
	return
}
