package main

import (
	"fmt"
	"time"
)

type Client struct {
	name string
}

var costumersNum int = 10
var timeBetweenCostumers int = 2
var timeOnOneCostumer int = 4
var placesInside int = 3
var names []string = []string{"Bob", "Jhon", "Mark", "Julia", "Ann", "Joe", "Alice", "Alex", "Paul", "Mr. Sanzches"}

var workDayDuration int = 30

func main() {
	fmt.Println("Process started")

	lineToBarber := make(chan Client, placesInside)

	lineOutside := make(chan Client, costumersNum)

	for i := 0; i < costumersNum; i++ {
		lineOutside <- Client{names[i%len(names)]}
	}

	stopClientsArriving := make(chan bool)
	stopBarberHaircutting := make(chan bool)

	go clientsArriving(lineToBarber, lineOutside, stopClientsArriving)
	go barberHaircutting(lineToBarber, stopBarberHaircutting)

	time.Sleep(time.Duration(workDayDuration) * time.Second)
	fmt.Println("Process finished")

}

func clientsArriving(lineToBarber chan Client, lineOutside chan Client, stop chan bool) {

	for {
		var c Client = <-lineOutside

		select {
		case lineToBarber <- c:
			fmt.Println("Client ", c.name, " came to the shop")
		default:
			fmt.Println("Client", c.name, "came to the shop, but there is no place to wait")
			lineOutside <- c
		}

		select {
		case <-time.After(time.Duration(timeBetweenCostumers) * time.Second):
		case <-stop:
			fmt.Println("Arriving was stopped")
			return

		}
	}

}

func barberHaircutting(lineToBarber chan Client, stop chan bool) {
	var barberSleeps bool = false

	for {
		var c Client
		select {
		case c = <-lineToBarber:
			if barberSleeps {
				fmt.Println(c.name, "woke up the barber")
				barberSleeps = false
			}
			fmt.Println("Start haircutting of ", c.name)
		default:
			if !barberSleeps {
				fmt.Println("Barber sleeps... zzz")
				barberSleeps = true

			}
		}

		select {
		case <-time.After(time.Duration(timeOnOneCostumer) * time.Second):
		case <-stop:
			fmt.Println("Haircutting was stopped")
			return

		}
	}
}
