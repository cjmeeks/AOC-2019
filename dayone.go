package main

import (
	//"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readFile(fname string) (nums []int, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil { return nil, err }

	lines := strings.Split(string(b), "\r\n")
	nums = make([]int, 0, len(lines))

	for _, l := range lines {
		if len(l) == 0 { continue }
		n, err := strconv.Atoi(l)
		if err != nil { return nil, err }
		nums = append(nums, n)
	}

	return nums, nil
}

func getFuelOfModule(moduleMass int) (newFuelCount int){
	newFuelCount = 0
	currentMass := moduleMass
	for {
		fuelMass := (currentMass / 3) - 2
		if fuelMass <= 0 {
			break
		}
		newFuelCount += fuelMass
		currentMass = fuelMass
	}
	return newFuelCount
}

//func main() {
//	nums, err := readFile("input.txt")
//	if err != nil { panic(err) }
//	var sum int
//	for _, num := range nums {
//		calc := getFuelOfModule(num)
//		sum += calc
//	}
//	fmt.Println(sum)
//	fmt.Println(getFuelOfModule(1969))
//}