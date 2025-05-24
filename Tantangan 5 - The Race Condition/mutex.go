package main

import (
	"fmt"
	"sync"
)

// Variabel global untuk demonstrasi race condition
var goCount int // Tidak dilindungi oleh mutex

// Variabel global yang dilindungi oleh mutex
var muCount int
var mu sync.Mutex // Mutex untuk melindungi muCount

func main() {
	var wg sync.WaitGroup
	jumlahGoroutine := 1000

	// Inisialisasi counter di awal main agar jelas
	goCount = 0
	muCount = 0

	wg.Add(jumlahGoroutine)
	for i := 0; i < jumlahGoroutine; i++ {
		go func() { // Tidak perlu parameter wg lagi karena wg global
			defer wg.Done()

			// ---- Tanpa Mutex ----
			goCount++
			// ---- Tanpa Mutex ----

			// ---- Dengan Mutex ----
			mu.Lock()
			muCount++
			mu.Unlock()
			// ---- Dengan Mutex ----
		}() // Perhatikan: tidak ada &wg di sini jika wg global
	}

	wg.Wait()
	fmt.Printf("Nilai tanpa Mutex: %d\n", goCount)
	fmt.Printf("Nilai dengan Mutex: %d\n", muCount)
	fmt.Println("Program Selesai.")
}
