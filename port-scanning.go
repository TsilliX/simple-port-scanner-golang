package main

import (
	"fmt"
	"net"
	"strconv"
)

func singlePortScanning(address string, port string) {
	fmt.Println("Im scanning for a single port...")
	internalAddr := address + ":" + port
	fmt.Println(internalAddr)

	_, err := net.Dial("tcp", internalAddr)

	if err != nil {
		fmt.Printf("The port %v its closed. \n", port)
	} else {
		fmt.Printf("The port %v is open. \n", port)
	}

}

func multiplePortsScanning(address string, firstPort string, lastPort string) {
	fmt.Println("Im scanning for multiple ports...")
	firstPortInteger, _ := strconv.ParseInt(firstPort, 10, 64)
	lastPortInteger, _ := strconv.ParseInt(lastPort, 10, 64)

	for i := firstPortInteger; i <= lastPortInteger; i++ {
		currentPort := strconv.FormatInt(i, 10)
		internalAddr := address + ":" + currentPort

		_, err := net.Dial("tcp", internalAddr)

		if err != nil {
			fmt.Printf("The port %v its closed. \n", currentPort)
		} else {
			fmt.Printf("The port %v is open. \n", currentPort)
		}

	}
}

func scanningAllPorts(address string) {
	fmt.Println("I'm scanning all the ports... Go make a coffee :)")
	for i := 1; i <= 65535; i++ {
		currentPort := strconv.Itoa(i)
		internalAddr := address + ":" + currentPort

		_, err := net.Dial("tcp", internalAddr)

		if err != nil {
			fmt.Printf("The port %v its closed. \n", currentPort)
		} else {
			fmt.Printf("The port %v is open. \n", currentPort)
		}

	}
}
