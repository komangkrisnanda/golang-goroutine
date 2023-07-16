package golanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T){
	channel := make(chan string)
	defer close(channel)

	// channel <- "John"

	// var data string

	// data = <- channel

	// //OR

	// //data := <- channel

	// fmt.Println(<- channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "John Wick"
		fmt.Println("Data sent.")
	}()

	data := <- channel
	fmt.Println(data)


}