package main

import (
	"fmt";
	"strconv";
)

type WireDirection struct {
	direction string
	distance int
}

func parseInput(wire1 [9]string, wire2 [8] string) (wire1Parse []WireDirection, wire2Parse []WireDirection) {
	index := 0
	wire1Parsed := make([]WireDirection, 0)
	wire2Parsed := make([]WireDirection, 0)
	for index < (len(wire1) - 1) {
		wire1code := wire1[index]
		directioncode := wire1code[0:1]
		distance, err := strconv.Atoi(wire1code[1:])
		if err != nil { return wire1Parsed, wire2Parsed }
		wire1Parsed = append(wire1Parsed, WireDirection{direction: directioncode, distance: distance})
		index += 1
	}
	index = 0
	for index < (len(wire2) - 1) {
		wire2code := wire2[index]
		directioncode := wire2code[0:1]
		distance, err := strconv.Atoi(wire2code[1:])
		if err != nil { return wire1Parsed, wire2Parsed }
		wire2Parsed = append(wire2Parsed, WireDirection{direction: directioncode, distance: distance})
		index += 1
	}
	return wire1Parsed, wire2Parsed
}

func main() {
	wire1 := [9]string {"R75","D30","R83","U83","L12","D49","R71","U7","L72"}
	wire2 := [8]string {"U62","R66","U55","R34","D71","R55","D58","R83"}
	//nums, err := readFile("input.txt")
	//if err != nil { panic(err) }
	wire1Parse, wire2Parsed := parseInput(wire1,wire2)
	fmt.Println(wire1Parse, wire2Parsed)
}