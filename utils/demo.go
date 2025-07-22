package utils

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var data = make(map[int]string)

func writeData(key int, value string) {
	data[key] = value
}

func updateValueIncorrectly(val string) string {
	return "updated-" + val
}

func Demo() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <key> <value>")
		return
	}

	key, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid key:", os.Args[1])
		return
	}

	value := os.Args[2]

	go writeData(key, value)

	fmt.Println("Value after updateValueIncorrectly call: ", updateValueIncorrectly(value))

	time.Sleep(2 * time.Second)

	fmt.Println("Reading data from map:", data[key])
}
