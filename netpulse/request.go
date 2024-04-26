package netpulse

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type request struct {
	Body        []byte              `json:"body,omitempty"`
	Destination string              `json:"destination"`
	Handle      string              `json:"handle"`
	Header      map[string][]string `json:"header"`
	Method      string              `json:"method"`
	URL         string              `json:"url"`
}

func reqMarshal(group, dest, handle string, in *http.Request) ([]byte, error) {
	out := &request{
		Destination: dest,
		Handle:      handle,
		Header:      map[string][]string(in.Header),
		Method:      in.Method,
	}
	if in.Body != nil {
		if body, err := ioutil.ReadAll(in.Body); err == nil {
			out.Body = body
		}
	}
	in.URL.Scheme = ""
	in.URL.Host = ""
	out.URL = in.URL.String()
	marshalled, err := json.Marshal(out)
	if err != nil {
		return nil, newError(errReqMarshal, err.Error())
	}
	return append([]byte(group+repl), zip(marshalled)...), nil
}

func reqUnmarshal(group string, p []byte) (*destination, *http.Request, error) {
	unzipped, err := unzip(p)
	if err != nil {
		return nil, nil, err.(*Error).escalate(errReqUnmarshal)
	}
	in := new(request)
	if err = json.Unmarshal(unzipped, in); err != nil {
		return nil, nil, newError(errReqUnmarshalJSON, err.Error())
	}
	out, err := http.NewRequest(in.Method, in.URL, bytes.NewBuffer(in.Body))
	if err != nil {
		return nil, nil, newError(errReqUnmarshalHTTP, err.Error())
	}
	out.Header = http.Header(in.Header)
	dest := new(destination)
	dest.group = group
	dest.handle = in.Handle
	dest.node = in.Destination
	return dest, out, nil
}
