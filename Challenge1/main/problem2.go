package main

import (
	"log"
	"time"
	"math/rand"
)

func problem2() {

	log.Printf("problem2: started --------------------------------------------")

	//
	// Todo:
	//
	// Throttle all go subroutines in a way,
	// that every one second one random number
	// is printed.
	//

	for inx := 0; inx < 10; inx++ {

		go printRandom2(inx)

	}

	//
	// Todo:
	//
	// Remove this quick and dirty sleep
	// against a synchronized wait until all
	// go routines are finished.
	//
	// Same as problem1...
	//

	time.Sleep(5 * time.Second)

	log.Printf("problem2: finished -------------------------------------------")
}

func printRandom2(slot int) {

	for inx := 0; inx < 10; inx++ {

		log.Printf("problem2: slot=%03d count=%05d rand=%f", slot, inx, rand.Float32())

	}
}