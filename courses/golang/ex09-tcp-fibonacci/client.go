package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/big"
	"net"
	"os"
	"strconv"
	"time"
)

type Response struct {
	Num  *big.Int
	Time time.Duration
}

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	handleError(err)

	defer conn.Close()

	// var resp Response
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		input, err := strconv.ParseInt(scan.Text(), 10, 64)
		handleError(err)

		encoder := json.NewEncoder(conn)
		encod := encoder.Encode(input)
		handleError(encod)

		var msg Response
		decoder := json.NewDecoder(conn)
		decod := decoder.Decode(&msg)
		handleError(decod)

		fmt.Printf("%s %d\n", msg.Time, msg.Num)
	}
}

func handleError(err error) {
	if err == nil {
		return
	}
	fmt.Println("ERROR: ", err.Error())
}
