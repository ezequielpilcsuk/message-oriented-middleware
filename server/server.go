package main

import (
	"fmt"
	"github.com/zeromq/goczmq"
	"log"
	"net"
	"strconv"
	"strings"
)

func main() {
	// Create a router socket and bind it to port 5555.
	fmt.Println(strconv.Itoa(-4))
	router, err := goczmq.NewRouter("tcp://*:5555")
	if err != nil {
		log.Fatal(err)
	}
	defer router.Destroy()

	log.Println("Router created and bound to port 5555")

	for {
		request, err := router.RecvMessage()
		if err != nil {
			log.Fatal(err)
		}
		go treatRequest(request, router)

	}
}

func treatRequest(req [][]byte, router *goczmq.Sock) {
	//TODO: Change message layout, Validate message, Accept different operations
	client := net.IPv4(req[0][0], req[0][1], req[0][2], req[0][3])
	msg := string(req[1])
	log.Printf("router received '%s' from '%v'", msg, client)

	if !validateRequest(msg) {
		log.Printf("invalid request '%s' from '%v'", msg, client)
		sendMessage(req[0], "invalid request", router)
		return
	}

	result := 0
	components := strings.Split(msg, " ")
	switch components[0] {
	case "add":
		for i := 1; i < len(components); i++ {
			value, _ := strconv.Atoi(components[i])
			result += value
			fmt.Printf("components[i] = %v\nresult = %v\nelem = %v\n", components[i], result, value)
		}
	case "sub":
		result, _ = strconv.Atoi(components[1])
		for i := 2; i < len(components); i++ {
			value, _ := strconv.Atoi(components[i])
			result -= value
			fmt.Printf("components[i] = %v\nresult = %v\nelem = %v\n", components[i], result, value)
		}
	case "mul":
		result, _ = strconv.Atoi(components[1])
		for i := 2; i < len(components); i++ {
			value, _ := strconv.Atoi(components[i])
			result *= value
			fmt.Printf("components[i] = %v\nresult = %v\nelem = %v\n", components[i], result, value)
		}
	case "div":
		result, _ = strconv.Atoi(components[1])
		for i := 2; i < len(components); i++ {
			value, _ := strconv.Atoi(components[i])
			result /= value
			fmt.Printf("components[i] = %v\nresult = %v\nelem = %v\n", components[i], result, value)
		}
	default:
		log.Printf("operation not supported from '%v'", client)
		sendMessage(req[0], "invalid operation", router)
		return
	}

	log.Printf("Sending %v\n", result)
	sendMessage(req[0], strconv.Itoa(result), router)

}

func validateRequest(msg string) bool {
	components := strings.Split(msg, " ")
	if len(components) <= 2 {
		return false
	}
	return true
}

func sendMessage(sender []byte, message string, router *goczmq.Sock) {
	err := router.SendFrame(sender, goczmq.FlagMore)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("router sent '%v'", message)
	
	err = router.SendFrame([]byte(message), goczmq.FlagNone)
	if err != nil {
		log.Fatal(err)
	}
}
