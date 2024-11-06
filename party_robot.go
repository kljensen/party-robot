package partyrobot

import (
	"fmt"
	"math"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	return fmt.Sprintf("Complete this function")
}

// ComputeDistance calculates the distance between two points.
func ComputeDistance(x1, y1, x2, y2 float64) float64 {
	ydelta := y2 - y1
	xdelta := x2 - x1
	squareDistance := math.Pow(ydelta, 2) + math.Pow(xdelta, 2)
	return math.Sqrt(squareDistance)
}
