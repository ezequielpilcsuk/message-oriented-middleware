package main

import (
	"github.com/zeromq/goczmq"
	"log"
	"strconv"
	"strings"
)

func main() {
	// Create a router socket and bind it to port 5555.
	router, err := goczmq.NewRouter("tcp://*:5555")
	if err != nil {
		log.Fatal(err)
	}
	defer router.Destroy()

	log.Println("router created and bound to port 5555")

	// Receive the message. Here we call RecvMessage, which
	// will return the message as a slice of frames ([][]byte).
	// Since this is a router socket that support async
	// request / reply, the first frame of the message will
	// be the routing frame.
	for {
		request, err := router.RecvMessage()
		if err != nil {
			log.Fatal(err)
		}

		go treatMessage(request, router)

	}
}

func treatMessage(req [][]byte, router *goczmq.Sock) {
	msg := string(req[1])
	log.Printf("router received '%s' from '%v'", msg, req[0])
	ops := strings.Split(msg, "+")
	op1, _ := strconv.Atoi(ops[0])
	op2, _ := strconv.Atoi(ops[1])
	result := op1 + op2

	// Send a reply. First we send the routing frame, which
	// lets the dealer know which client to send the message.
	// The FlagMore flag tells the router there will be more
	// frames in this message.
	err := router.SendFrame(req[0], goczmq.FlagMore)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("router sent '%v'", result)

	// Next send the reply. The FlagNone flag tells the router
	// that this is the last frame of the message.
	err = router.SendFrame([]byte(string(result)), goczmq.FlagNone)
	if err != nil {
		log.Fatal(err)
	}

}
