package main

import (
	"SWE/netpulse"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type additionHandler struct{}

func (h *additionHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	numbers := strings.Split(string(body), " ")

	num1, err := strconv.Atoi(numbers[0])
	if err != nil {
		http.Error(res, "Invalid input", http.StatusBadRequest)
		return
	}

	num2, err := strconv.Atoi(numbers[1])
	if err != nil {
		http.Error(res, "Invalid input", http.StatusBadRequest)
		return
	}

	result := num1 + num2
	res.Write([]byte(strconv.Itoa(result)))
	res.Write(body)
}

func main() {
	handler := new(additionHandler)
	config := &netpulse.Config{
		Interface: "eth0",
		Handler:   handler,
		LogLevel:  "debug",
		Service:   "addition-service",
	}
	server, err := netpulse.New(config)
	if err != nil {
		panic(err.Error())
	}
	defer server.Close()
	http.ListenAndServe(":9999", handler)
}
