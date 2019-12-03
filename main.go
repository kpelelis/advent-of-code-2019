package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

import (
	"github.com/kpelelis/advent-of-code-2019/ioutils"
)

func calculateFuel(mass int) (int) {
	return int(math.Floor(float64(mass) / 3)) - 2
}

func day1() {
	masses, err := ioutils.ReadInputOfInts("./day1.txt")
	if err != nil {
		log.Fatal(err)
	}

	res := 0
	for i := 0; i < len(masses); i++ {
		res += calculateFuel(masses[i])
	}
	fmt.Printf("%v\n", res)
}

func day1b() {
	masses, err := ioutils.ReadInputOfInts("./day1.txt")
	if err != nil {
		log.Fatal(err)
	}

	res := 0
	for i := 0; i < len(masses); i++ {
		for fuel := calculateFuel(masses[i]); fuel > 0; fuel = calculateFuel(fuel) {
			res += fuel
		}
	}
	fmt.Printf("%v\n", res)
}

func runProgram(program []int, noun, verb int) int{
	program[1] = noun
	program[2] = verb

	for i := 0; i < len(program); {
		switch opcode := program[i]; opcode {
		case 1:
			lhr := program[i + 1]
			rhr := program[i + 2]
			resp := program[i + 3]
			if lhr >= len(program) || rhr >= len(program) || resp >= len(program ){
				log.Fatal("Out of bounds")
				break
			}
			program[resp] = program[lhr] + program[rhr]
			i += 4
		case 2:
			lhr := program[i + 1]
			rhr := program[i + 2]
			resp := program[i + 3]
			if lhr >= len(program) || rhr >= len(program) || resp >= len(program ){
				log.Fatal("Out of bounds")
				break
			}
			program[resp] = program[lhr] * program[rhr]
			i += 4
		case 99:
			i = len(program)
			break
		}
	}
	return program[0]
}

func day2() {
	lines, err := ioutils.ReadIntCSV("./day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	res := runProgram(lines[0], 12, 2)
	fmt.Printf("%v\n", res)
}

func day2b() {
	lines, err := ioutils.ReadIntCSV("./day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Speed up with binary search
	for noun := 0; noun < len(lines[0]) - 1; noun++ {
		for verb := 0; verb < len(lines[0]) - 1; verb++ {
			program := make([]int, len(lines[0]))
			copy(program, lines[0])
			res := runProgram(program, noun, verb)
			if (res == 19690720) {
				fmt.Printf("%v\n", 100 * noun + verb)
				return
			}
		}
	}
	fmt.Printf("No result!\n")
}


func main() {
	switch day := os.Args[1]; day {
	case "1":
		day1()
	case "1b":
		day1b()
	case "2":
		day2()
	case "2b":
		day2b()
	default:
		fmt.Println("Not implemented yet")
	}
}
