package main

import (
	"fmt"
	"keyboard"
	"log"
	"math/rand"
)

func guess(number int, taget int) bool {
	guessed = append(guessed, number)
	if number < taget {
		fmt.Println("哎呀，你猜低了！")
		fmt.Println("guessed:", guessed)
		return false
	} else if number > taget {
		fmt.Println("哎呀，你猜高了！")
		fmt.Println("guessed:", guessed)
		return false
	} else {
		fmt.Println("干得好，你猜对了！")
		return true
	}

}

func getTarget() int {
	taget := rand.Intn(100) + 1
	return taget
}

var guessed []int
var count int = 10

func main() {
	taget := getTarget()
	success := false
	for i := 0; i < count; i++ {
		fmt.Printf("you have %v chance.\n", count-i)
		fmt.Print("plase guss a number 0-100:")
		number, err := keyboard.GetFloat()
		if err != nil {
			log.Fatal(err)
		}
		success = guess(number, taget)
		if success {
			break
		}
	}
	if !success {
		fmt.Printf("对不起。你没有猜对。它是：[%v]\n", taget)
		fmt.Println("guessed:", guessed)
	}

}
