package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

func handleConnection(connection net.Conn, messages []string) {
	defer connection.Close()

	for {
		for _, message := range messages {
			// Generate a random float between 0 and 360 and reduce the range to 60 degrees
			bearing := rand.Float64() * 360 / 6.
			fmt.Println("Sending message:", modifyBearing(message, bearing))
			modifiedMessage := modifyBearing(message, bearing)
			_, err := connection.Write([]byte(modifiedMessage + "\n"))

			if err != nil {
				netErr, ok := err.(*net.OpError)
				if ok && netErr.Err.Error() == "broken pipe" {
					fmt.Println("Client disconnected.")
					return
				}

				fmt.Println("Error occurred:", err)
				return
			}

		}
		time.Sleep(500 * time.Millisecond)
	}
}

func modifyBearing(message string, bearing float64) string {
	headingString := "INHDT," + strconv.FormatFloat(bearing, 'f', 2, 64) + ",T"
	checksum := calculateChecksum(headingString)
	headingMessage := "$" + headingString + "*" + strconv.FormatInt(int64(checksum), 16)
	return strings.Replace(message, "$INHDT,0.00,T*XX", headingMessage, 1)
}

func calculateChecksum(data string) byte {
	// XOR all the bytes in the data
	var checksum byte
	for i := 0; i < len(data); i++ {
		checksum ^= data[i]
	}

	return checksum
}

func main() {
	// Define command-line flags
	host := flag.String("host", "127.0.0.1", "the host IP address")
	port := flag.Int("port", 5001, "the port number")
	flag.Parse()

	tcpIP := *host
	tcpPort := *port

	nmeaMessages := []string{
		"$INZDA,163611.11,10,09,2019,,*76",
		"$INGGA,163611.11,7849.766185,N,00543.603403,W,2,07,1.2,2.95,M,36.08,M,2.0,0123*4D",
		"$INGLL,7849.766185,N,00543.603403,W,163611.11,A,D*65",
		"$INVTG,61.99,T,68.23,M,5.0,N,9.2,K,D*30",
		"$INHDT,0.00,T*XX",
		"$PSXN,23,0.01,0.12,66.99,0.07*0D",
	}

	listener, err := net.Listen("tcp", tcpIP+":"+strconv.Itoa(tcpPort))
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go handleConnection(connection, nmeaMessages)

	}

}
