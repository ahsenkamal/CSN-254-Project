package netpulse

import "net/http"

type whisperer interface {
	Whisper(addr string, payload []byte) error
}

type writer struct {
	http.ResponseWriter
	group     string
	output    *response
	peer      string
	whisperer whisperer
}

func (w *writer) Header() http.Header {
	return w.output.Header
}

func (w *writer) Write(data []byte) (int, error) {
	if w.output.Code == 0 {
		w.WriteHeader(http.StatusOK)
	}
	header := w.Header()
	if header.Get("Content-Type") == "" {
		header.Add("Content-Type", http.DetectContentType(data))
	}
	w.output.Body = data
	payload := resMarshal(w.group, w.output)
	if err := w.whisperer.Whisper(w.peer, payload); err != nil {
		return 0, newError(errResWhisper, err.Error())
	}
	return len(data), nil
}

func (w *writer) WriteHeader(code int) {
	w.output.Code = code
}

func newWriter(node whisperer, dest *destination) *writer {
	return &writer{
		group: dest.group,
		output: &response{
			Handle: dest.handle,
			Header: http.Header(make(map[string][]string)),
		},
		peer:      dest.node,
		whisperer: node,
	}
}
