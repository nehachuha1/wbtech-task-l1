package main

import (
	"fmt"
	"time"
)

const (
	workTime  = 10
	dataDelay = 500
)

type Data struct {
	Message string
}

func ReceiveData(input chan Data) {
	for data := range input {
		fmt.Printf("Message text: %v\n", data.Message)
	}
}

func SendData(output chan Data) {
	newTicker := time.NewTicker(dataDelay * time.Millisecond)

	for {
		select {
		case <-newTicker.C:
			output <- Data{Message: "message text..."}
		}
	}
}

func main() {
	dataCh := make(chan Data)

	go SendData(dataCh)
	go ReceiveData(dataCh)

	time.Sleep(5 * time.Second)
}
