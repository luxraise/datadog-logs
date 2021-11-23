package logs

func MakeHost(hostname string) (h Host) {
	h.Hostname = hostname
	return
}

type Host struct {
	logCore
}

func (h *Host) NewSource(source string) *Source {
	var s Source
	s.logCore = h.logCore
	s.Source = source
	return &s
}
