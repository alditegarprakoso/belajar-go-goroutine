package belajargogoroutine

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGetGomaxprocs(t *testing.T) {
	totalCPU := runtime.NumCPU()
	fmt.Println("Total CPU :", totalCPU)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread :", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine :", totalGoroutine)
}
func TestChangeThreadNumber(t *testing.T) {
	totalCPU := runtime.NumCPU()
	fmt.Println("Total CPU :", totalCPU)

	runtime.GOMAXPROCS(20) // Mengubah jumlah thread
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread :", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine :", totalGoroutine)
}
