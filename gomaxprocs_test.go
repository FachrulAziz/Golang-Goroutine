package belajargolanggoroutine

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGomaxprocs(t *testing.T) {
	TotalCPU := runtime.NumCPU()
	fmt.Println("Total CPU = ", TotalCPU)

	TotalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread = ", TotalThread)

	TotalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine = ", TotalGoroutine)
}
