package main

import (
	"log"
	"time"
	"math/rand"
)

func problem1() {

	log.Printf("problem1: started --------------------------------------------")

	//
	// Todo:
	//
	// Quit all go routines after
	// a total of exactly 100 random
	// numbers have been printed.
	//
	// Do not change the 25 in loop!
	//

	for inx := 0; inx < 10; inx++ {

		go printRandom1(inx)

	}

	//
	// Todo:
	//
	// Remove this quick and dirty sleep
	// against a synchronized wait until all
	// go routines are finished.
	//

	time.Sleep(5 * time.Second)

	log.Printf("problem1: finised --------------------------------------------")
}

func printRandom1(slot int) {

	//
	// Do not change 25 into 10!
	//

	for inx := 0; inx < 25; inx++ {

		log.Printf("problem1: slot=%03d count=%05d rand=%f", slot, inx, rand.Float32())

	}
}