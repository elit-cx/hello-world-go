package main

import (
    "fmt"
    "time"
)

func helloworld() string {
	time.Sleep(10 * time.Second)
	return "Hello World!!"
}

func main() {
	fmt.Println(helloworld())
}
