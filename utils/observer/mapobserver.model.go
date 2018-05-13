package observer

type MapPublisher struct {
	Observers map[Observer]struct{}
}

func (obs *MapPublisher) Register(o Observer) {
	obs.Observers[o] = struct{}{}
}

func (obs *MapPublisher) Deregister(o Observer) {
	delete(obs.Observers, o)
}

func (obs *MapPublisher) Notify(e Event) {
	for o := range obs.Observers {
		o.OnNotify(e)
	}
}
