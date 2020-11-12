package main

import (
	"fmt"
	"time"
)

func main() {
	akpingChan := make(chan int, 1)
	akpongChan := make(chan int, 1)

	go akping(akpingChan, akpongChan)
	go akpong(akpongChan, akpingChan)

	akpingChan <- 1

	select {}
}

func akping(akpingChan <-chan int, akpongChan chan<- int) {
	for {
		ball := <-akpingChan

		fmt.Println("Ping", ball)
		time.Sleep(1 * time.Second)

		akpongChan <- ball + 1
	}
}

func akpong(akpongChan <-chan int, akpingChan chan<- int) {
	for {
		ball := <-akpongChan

		fmt.Println("akPong", ball)
		time.Sleep(1 * time.Second)

		akpingChan <- ball + 1
	}
}
