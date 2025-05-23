package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func HitungJumlah(id int) {
	defer wg.Done()

	sum := 0
	for i := 1; i <= id; i++ {
		sum += i
	}
	fmt.Printf("Goroutine ID: %d telah selesai di jalankan, jumlah 1 sampai dengan %d adalah %d\n", id, id, sum)
	time.Sleep(100 * time.Millisecond)
}

func main() {
	var goCount = 10

	wg.Add(goCount)
	for i := 1; i <= goCount; i++ {
		go HitungJumlah(i)
	}

	wg.Wait()
	fmt.Println("Semua Go routine telah selasai di jalankan...")
}
