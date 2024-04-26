package main

import (
	"io/ioutil"
	"net/http"
	"SWE/netpulse"
)

type echoHandler struct{}

func (h *echoHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	res.Write(body)
}

func main() {
	handler := new(echoHandler)
	config := &netpulse.Config{
		Interface: "eth0",
		Handler:   handler,
		LogLevel:  "debug",
		Service:   "echo-service",
	}
	server, err := netpulse.New(config)
	if err != nil {
		panic(err.Error())
	}
	defer server.Close()
	http.ListenAndServe(":8989", handler)
}
