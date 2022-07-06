package main

import (
	"fmt"
	"github.com/zeromq/goczmq"
	"log"
	"os"
	"strconv"
)

func main() {
	// Create a dealer socket and connect it to the router.
	dealer, err := goczmq.NewDealer("tcp://127.0.0.1:5555")
	if err != nil {
		log.Fatal(err)
	}
	defer dealer.Destroy()

	log.Println("dealer created and connected")

	// Send a 'Hello' message from the dealer to the router.
	// Here we send it as a frame ([]byte), with a FlagNone
	msg := ""
	if len(os.Args) <= 1 {
		fmt.Printf("No arguments passed. Requesting operation:\n")
		msg = "8+3"
	} else {
		for i := 1; i < len(os.Args); i++ {
			msg += fmt.Sprintf("%v ", os.Args[i])
		}
		msg = msg[:len(msg)-1]
		//fmt.Printf("The message was %v\n", msg)
	}

	// flag to indicate there are no more frames following.
	err = dealer.SendFrame([]byte(msg), goczmq.FlagNone)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("dealer sent '%v'", msg)

	// Receive the reply.
	reply, err := dealer.RecvMessage()
	if err != nil {
		log.Fatal(err)
	}

	//TODO: Treat response to properly display it
	resp := string(reply[0])

	got, _ := strconv.Atoi(resp)

	log.Printf("result: %v", got)
}
