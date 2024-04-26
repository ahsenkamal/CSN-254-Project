package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"SWE/netpulse"
)

func main() {
	service := "echo-service"
	config := &netpulse.Config{Interface: "eth0", LogLevel: "listen"}
	client, err := netpulse.New(config)
	if err != nil {
		panic(err.Error())
	}
	defer client.Close()
	client.WaitFor(service)
	input := "This is the value I am inputting."
	body := bytes.NewBuffer([]byte(input))
	request, _ := http.NewRequest("POST", "netpulse://"+service+"/", body)
	response, err := client.Do(request)
	if err != nil {
		panic(err.Error())
	}
	output, _ := ioutil.ReadAll(response.Body)
	if string(output) == input {
		fmt.Println("It works.")
	} else {
		fmt.Println("It doesn't work.")
	}
}
