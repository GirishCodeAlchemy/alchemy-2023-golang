// Package main provides ...
package library

import (
	"fmt"
	"time"
)

func sendmessage(Channel chan string) {
	fmt.Println("Sending message")
	for {
		Channel <- "Hello" + time.Now().String()
		fmt.Println("message sent @" + time.Now().String())
		time.Sleep(time.Second * 1)

	}
}

func recivemessage(channel chan string) {
	fmt.Println("Reciving message")
	for {
		fmt.Println("old", <-channel)
	}
}

func SendmessageAck(msg chan string, ack chan string) {
	fmt.Println("Sending message")
	for {
		msg <- "Sent message " + time.Now().String()
		time.Sleep(time.Second * 5)

		ackmsg := <-ack
		fmt.Println("sentAck -->", ackmsg)

	}

}

func RecivemessageAck(msg chan string, ack chan string) {
	fmt.Println("Reciving message")
	for {
		msgRec := <-msg
		fmt.Println("receACK -->", msgRec)
		ack <- "Reciieved message @" + time.Now().String()
		fmt.Println("receACK -->", ack)
	}

}

func ChannelFunctions() {
	channel := make(chan string)

	go sendmessage(channel)
	go recivemessage(channel)

	// // Buffered channel
	// channel1 := make(chan string, 10)
	// go sendmessage(channel1)
	// go recivemessage(channel1)

	var stopchannel string
	fmt.Scanln(&stopchannel)
	msgChan := make(chan string)
	ackChan := make(chan string)
	go SendmessageAck(msgChan, ackChan)
	go RecivemessageAck(msgChan, ackChan)

	fmt.Scanln(&stopchannel)

}

// When a channel needs to send a signal of an event without the need for sending any data. From the below piece of code, we can see that we are sending a signal using sending empty struct to the channel which is sent to the workerRoutine.

func workerRoutine(ch chan struct{}) {
	// Receive message from main program.
	<-ch
	println("Signal Received")

	// Send a message to the main program.
	close(ch)
}

func ChannelWithNodata() {
	//Create channel
	ch := make(chan struct{})

	//define workerRoutine
	go workerRoutine(ch)

	// Send signal to worker goroutine
	ch <- struct{}{}

	// Receive a message from the workerRoutine.
	<-ch
	println("Signal Received")

}
