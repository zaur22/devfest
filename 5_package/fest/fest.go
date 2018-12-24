package fest

func NewDevFest(year int) DevFest {
	return DevFest{
		festival: festival{
			Topics: make(map[string]string),
		},
	}
}

type festival struct {
	Topics   map[string]string
	Visitors []string
}

type DevFest struct {
	festival
	Year int
}

func (f *festival) AddVisitor(name string) {
	f.Visitors = append(f.Visitors, name)
}
