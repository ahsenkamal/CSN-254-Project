package netpulse

import "sync"

type notifier struct {
	*sync.Mutex
	active bool
	stream chan struct{}
}

func (n *notifier) activate() {
	n.Lock()
	defer n.Unlock()
	n.active = true
}

func (n *notifier) deactivate() {
	n.Lock()
	defer n.Unlock()
	n.active = false
}

func (n *notifier) notify() {
	n.Lock()
	defer n.Unlock()
	if n.active {
		n.stream <- struct{}{}
	}
}
