package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func reader(ch chan int, mutex *sync.RWMutex) {
	for {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		mutex.RLock()
		ch <- 1
		time.Sleep(time.Duration(rand.Intn(400)) * time.Millisecond)
		ch <- -1
		mutex.RUnlock()
	}
}
func writer(ch chan int, mutex *sync.RWMutex) {
	for {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		mutex.Lock()
		ch <- 1
		time.Sleep(time.Duration(rand.Intn(400)) * time.Millisecond)
		ch <- -1
		mutex.Unlock()
	}
}
func main() {
	var mutex sync.RWMutex
	var readCount, writeCount int
	var temp int
	readersChannel := make(chan int)
	writersChannel := make(chan int)

	go reader(readersChannel, &mutex)
	go reader(readersChannel, &mutex)
	go reader(readersChannel, &mutex)

	go writer(writersChannel, &mutex)
	go writer(writersChannel, &mutex)
	go writer(writersChannel, &mutex)

	for {
		select {
		case temp = <-readersChannel:
			readCount += temp
		case temp = <-writersChannel:
			writeCount += temp
		}
		fmt.Printf("%s%s\n", strings.Repeat("R", readCount), strings.Repeat("W", writeCount))
	}
}
