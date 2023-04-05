package main

import (
	"fmt"
	"math/rand"
	"time"
)

var restartNotifier = make(chan string)

func main() {
	HandlePriceWithInt()
	HandlePriceWithFloat()
	fmt.Println("---")
	
	SelfRestartApp()
}

func SelfRestartApp() {
	var fruits = []string{"Apple", "Banana", "Avocado", "Pear"}

	defer RecoverFunc()

	for {
		rand.Seed(time.Now().UnixNano())
		i := rand.Intn(5)
		fmt.Println("Special fruit for today: ", fruits[i])

		time.Sleep(500 * time.Millisecond)
	}
}

func RecoverFunc() {
	err := recover()
	if err != nil {
		fmt.Println("\nPanic occured. Restarting app..")
		go func() {
			defer func() { restartNotifier <- "App restarted" }()
			finish := time.Now().Add(500 * time.Millisecond)
			for {
				if time.Now().After(finish) {
					break
				}

				print(".")
				time.Sleep(100 * time.Millisecond)
			}
		}()
		time.Sleep(500 * time.Millisecond)
		println("..........................")
		println()
		SelfRestartApp()
	}
}
