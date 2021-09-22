package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net"
	"time"
)

type Response struct {
	Num  *big.Int
	Time time.Duration
}

var cache map[int]*big.Int

func main() {
	cache = make(map[int]*big.Int)
	cache[0] = big.NewInt(0)
	cache[1] = big.NewInt(1)

	fmt.Println("Launching server...")

	ln, err := net.Listen("tcp", ":8081")
	handleError(err)

	conn, err := ln.Accept()
	handleError(err)

	for {

		var msg int
		decoder := json.NewDecoder(conn)
		decod := decoder.Decode(&msg)
		handleError(decod)

		t := time.Now()
		resp := Response{
			Num:  fibonacci(msg),
			Time: time.Since(t),
		}

		encoder := json.NewEncoder(conn)
		encod := encoder.Encode(resp)
		handleError(encod)
	}
}

func handleError(err error) {
	if err == nil {
		return
	}
	// fmt.Println("ERROR: ", err.Error())
}

func fibonacci(num int) *big.Int {
	size := len(cache)

	if num < size {
		return cache[num]
	}

	a := cache[size-2]
	b := cache[size-1]

	var limit big.Int
	limit.Exp(big.NewInt(10), big.NewInt(99), nil)

	for i := size; i <= num; i++ {
		a.Add(a, b)
		cache[i] = a
		a, b = b, a
	}
	return b

}
