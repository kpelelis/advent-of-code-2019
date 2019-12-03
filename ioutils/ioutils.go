package ioutils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readLines(filename string) (lines []string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func ReadInputOfInts(filename string) ([]int, error) {
	lines, err := readLines(filename)
	if err != nil {
		return nil, err
	}

	var ret []int
	for i := 0; i < len(lines); i++ {
		num, err := strconv.Atoi(lines[i])
		if err != nil {
			break
		}
		ret = append(ret, num)
	}
	return ret, err
}

func ReadIntCSV(filename string) ([][]int, error) {
	lines, err := readLines(filename)
	if err != nil {
		return nil, err
	}

	var ret [][]int
	for i := 0; i < len(lines); i++ {
		nums := strings.Split(lines[i], ",")
		var line []int
		for j := 0; j < len(nums); j++ {
			num, err := strconv.Atoi(nums[j])
			if err != nil {
				break
			}
			line = append(line, num)
		}
		ret = append(ret,line)
	}
	return ret, err
}
