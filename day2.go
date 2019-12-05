package main

import (
	"fmt"
)

func parseInput(nums [165]int) (answer int) {
	index := 0
	//original := nums
	code := nums[index]
	for code != 99 {
		code = nums[index]
		switch code {
		case 1:
			num1 := nums[nums[index + 1]]
			num2 := nums[nums[index + 2]]
			storePos := nums[index + 3]
			nums[storePos] = num1 + num2
			fmt.Println(num1," : ", num2, " : ",storePos, " = ", nums[storePos], " index = ", index )
			index += 4
			continue
		case 2:
			num1 := nums[nums[index + 1]]
			num2 := nums[nums[index + 2]]
			storePos := nums[index + 3]
			nums[storePos] = num1 * num2
			fmt.Println(num1," : ", num2, " : ",storePos, " = ", nums[storePos], " index = ", index )
			index += 4
			continue
		}
	}
	return nums[0]
}

func main() {
	input := [165]int {1,41,12,3,1,1,2,3,1,3,4,3,1,5,0,3,2,1,10,19,1,6,19,23,1,10,23,27,2,27,13,31,1,31,6,35,2,6,35,39,1,39,5,43,1,6,43,47,2,6,47,51,1,51,5,55,2,55,9,59,1,6,59,63,1,9,63,67,1,67,10,71,2,9,71,75,1,6,75,79,1,5,79,83,2,83,10,87,1,87,5,91,1,91,9,95,1,6,95,99,2,99,10,103,1,103,5,107,2,107,6,111,1,111,5,115,1,9,115,119,2,119,10,123,1,6,123,127,2,13,127,131,1,131,6,135,1,135,10,139,1,13,139,143,1,143,13,147,1,5,147,151,1,151,2,155,1,155,5,0,99,2,0,14,0}
	//nums, err := readFile("input.txt")
	//if err != nil { panic(err) }
	answer := parseInput(input)
	fmt.Println(answer)
}
//19690720
//19229920
//14621910
//14621918
//6327510