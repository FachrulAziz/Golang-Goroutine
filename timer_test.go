package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)
}

func TestAfter(t *testing.T) {
	channel := time.After(6 * time.Second)
	fmt.Println(time.Now())

	time := <-channel
	fmt.Println(time)
}
func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(7*time.Second, func() {
		fmt.Println(time.Now())
	})
	fmt.Println(time.Now())

	group.Wait()
}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	// ticker.Stop() berfungsi untuk menghentikan ticker yang terus berjalan
	go func() {
		time.Sleep(9 * time.Second)
		ticker.Stop()
	}()
	for time := range ticker.C {
		fmt.Println(time)
	}
}
