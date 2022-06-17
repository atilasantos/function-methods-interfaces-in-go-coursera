package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func GenDisplaceFn(acceleration float64, initVelocity float64, initDisplacement float64) func(t float64) float64 {
	fn := func(time float64) float64 {
		return 0.5*acceleration*math.Pow(time, 2) + initVelocity*time + initDisplacement
	}

	return fn
}

func main() {

	var acceleration float64
	var acc string

	var initVelocity float64
	var iVel string

	var initDisplacement float64
	var iDispl string

	var time float64
	var t string

	fmt.Println("Type acceleration")
	_, err := fmt.Scan(&acc)
	acceleration, _ = strconv.ParseFloat(acc, 64)

	fmt.Println("Type initial velocity")
	_, err = fmt.Scan(&iVel)
	initVelocity, _ = strconv.ParseFloat(iVel, 64)

	fmt.Println("Type initial displacement")
	_, err = fmt.Scan(&iDispl)
	initDisplacement, _ = strconv.ParseFloat(iDispl, 64)

	fmt.Println("Type time")
	_, err = fmt.Scan(&t)
	time, _ = strconv.ParseFloat(t, 64)

	if err != nil {
		log.Fatal(err)
	}

	fn := GenDisplaceFn(acceleration, initVelocity, initDisplacement)
	fmt.Println(fn(time))

}
