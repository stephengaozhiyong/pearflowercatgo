package datafile

import (
	"bufio"
	"os"
	"strconv"
)

func GetFloats(filename string) ([]float64, error) {
	var numbers []float64
	file, err := os.Open(filename)
	// fmt.Println(reflect.TypeOf(file), file)
	if err != nil {
		return nil, err
	}
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
		i++
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, err
	}
	return numbers, nil
}
