package main

import "fmt"

func main() {
	var typeOfScan int
	var address string
	var port string
	var startPort string
	var stopPort string

	fmt.Println("Port scanner with Golang!")

	// create menu
	fmt.Println("#####")
	fmt.Println("Select your type of scanning")
	fmt.Println("Press '1' for single port scanning")
	fmt.Println("Press '2' for multiple port scanning")
	fmt.Println("Press '3' for scanning all ports")
	fmt.Println("Press '4' to exit")
	fmt.Println("#####")

	for {
		// get the type of scanning
		fmt.Print("Select: ")
		fmt.Scan(&typeOfScan)

		if typeOfScan == 1 {

			fmt.Println("Provide the IP address or hostname (e.g google.com): ")
			fmt.Scan(&address)

			fmt.Println("Provide the port: ")
			fmt.Scan(&port)

			singlePortScanning(address, port)

		} else if typeOfScan == 2 {
			fmt.Println("Provide the IP address or hostname (e.g google.com): ")
			fmt.Scan(&address)

			fmt.Println("Provide the start port: ")
			fmt.Scan(&startPort)

			fmt.Println("Provide the stop port: ")
			fmt.Scan(&stopPort)

			multiplePortsScanning(address, startPort, stopPort)

		} else if typeOfScan == 3 {
			fmt.Println("Provide the IP address or hostname (e.g google.com): ")
			fmt.Scan(&address)

			scanningAllPorts(address)

		} else if typeOfScan == 4 {
			fmt.Println("Goodbye :)")
			break
		} else {
			fmt.Println("Provide a valid numeric digit")
		}
	}

}
