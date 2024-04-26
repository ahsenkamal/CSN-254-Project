package netpulse

import "sync"

type workers struct {
	*sync.Mutex
	current int
	list    []*peer
}

func (w *workers) add(p *peer) int {
	w.Mutex.Lock()
	defer w.Mutex.Unlock()
	for _, service := range w.list {
		if service.name == p.name {
			return len(w.list)
		}
	}
	w.list = append(w.list, p)
	return len(w.list)
}

func (w *workers) available() bool {
	w.Mutex.Lock()
	defer w.Mutex.Unlock()
	return len(w.list) > 0
}

func (w *workers) next() *peer {
	w.Mutex.Lock()
	defer w.Mutex.Unlock()
	length := len(w.list)
	current := w.current
	if length == 0 {
		return nil
	}
	if current < length {
		w.current++
		return w.list[current]
	}
	w.current = 1
	return w.list[0]
}

func (w *workers) remove(name string) (int, *peer) {
	w.Mutex.Lock()
	defer w.Mutex.Unlock()
	for i, p := range w.list {
		if p.name == name {
			w.list = append(w.list[0:i], w.list[i+1:]...)
			return len(w.list), p
		}
	}
	return len(w.list), nil
}

func newWorkers() *workers {
	return &workers{
		Mutex: new(sync.Mutex),
		list:  make([]*peer, 0),
	}
}
