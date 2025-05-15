package average2

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func Average(nums ...float64) float64 {
	var sum float64 = 0
	for _, n := range nums {
		sum += n
	}
	count := float64(len(nums))
	avg := sum / count
	return avg
}

func main() {
	args := os.Args[1:]
	var nums []float64
	for _, arg := range args {
		n, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, n)
	}
	fmt.Printf("平均数：%v\n", Average(nums...))
}
