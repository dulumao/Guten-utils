package event

type Listener struct {
	Callback func(event *Event) error
	Priority int
}
type Listeners []Listener

func (self Listeners) Len() int {
	return len(self)
}

func (self Listeners) Less(i, j int) bool {
	return self[i].Priority < self[j].Priority
}

func (self Listeners) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}
