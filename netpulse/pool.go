package netpulse

import "sync"

type pool struct {
	*sync.Mutex
	workers map[string]*workers // map[service-type]service-workers
}

func (p *pool) add(service string) {
	p.Lock()
	defer p.Unlock()
	if p.workers[service] == nil {
		p.workers[service] = newWorkers()
	}
}

func (p *pool) get(service string) (*workers, bool) {
	p.Lock()
	defer p.Unlock()
	w, ok := p.workers[service]
	return w, ok
}

func (p *pool) remove(service string) {
	p.Lock()
	defer p.Unlock()
	if p.workers[service] != nil {
		delete(p.workers, service)
	}
}
