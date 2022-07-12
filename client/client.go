package main

import (
	"fmt"
	"github.com/zeromq/goczmq"
	"log"
	"os"
	"strconv"
)

func main() {
	dealer, err := goczmq.NewDealer("tcp://127.0.0.1:5555")
	if err != nil {
		log.Fatal(err)
	}
	defer dealer.Destroy()

	log.Println("dealer created and connected")

	if len(os.Args) < 4 {
		fmt.Printf("invalid request, all valid requests use at least 3 arguments:\n")
		//TODO: could ask for user input in case not enough arguments
		return
	}

	msg := ""
	for i := 1; i < len(os.Args); i++ {
		msg += fmt.Sprintf("%v ", os.Args[i])
	}
	msg = msg[:len(msg)-1]

	err = dealer.SendFrame([]byte(msg), goczmq.FlagNone)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("dealer sent '%v'", msg)

	reply, err := dealer.RecvMessage()
	if err != nil {
		log.Fatal(err)
	}

	resp := string(reply[0])

	got, _ := strconv.Atoi(resp)

	log.Printf("result: %v", got)
}
