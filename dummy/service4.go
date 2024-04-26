package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"SWE/netpulse"
)

func main() {
	service := "addition-service"
	config := &netpulse.Config{Interface: "eth0", LogLevel: "listen"}
	client, err := netpulse.New(config)
	if err != nil {
		panic(err.Error())
	}
	defer client.Close()
	client.WaitFor(service)
	input := "12 13"
	body := bytes.NewBuffer([]byte(input))
	request, _ := http.NewRequest("POST", "netpulse://"+service+"/", body)
	response, err := client.Do(request)
	if err != nil {
		panic(err.Error())
	}
	output, _ := ioutil.ReadAll(response.Body)
	if string(output) == "25" {
		fmt.Println("It works. The output is ", string(output), ".")
	} else {
		fmt.Println("It doesn't work. The output is ", string(output), ".")
	}
}
