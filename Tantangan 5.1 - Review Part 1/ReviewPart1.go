package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var workerWg sync.WaitGroup
var mu sync.Mutex

var inputChan = make(chan int, 5)
var outputChan = make(chan int)
var workers = 3
var totalProcessed int

func main() {

	//---- Goroutine Generator ----
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(inputChan)
		defer fmt.Println("Generator Berhenti, channel input telah ditutup")

		fmt.Println("Generator Mulai Menghasilkan Angka")

		for i := 1; i <= 20; i++ {
			inputChan <- i
			fmt.Println("Generator: mengirim data:", i)
			time.Sleep(50 * time.Millisecond)
		}

	}()
	//---- Goroutine Generator ----

	//---- Goroutine Pekerja ----
	wg.Add(workers)
	workerWg.Add(workers)

	for i := 1; i <= workers; i++ {
		go func(workerID int) {
			defer wg.Done()
			defer workerWg.Done()
			defer fmt.Println("Pekerja Berhenti, channel output telah ditutup")

			for data := range inputChan {
				outputChan <- (data * data) // Menghitung angka kuadrat

				mu.Lock()
				totalProcessed++ // Menghitung total Prosess
				mu.Unlock()

				fmt.Printf("Pekerja %d: Memproses %d, hasil kuadrat %d\n", workerID, data, data*data)
				time.Sleep(75 * time.Millisecond)
			}
		}(i)
	}

	wg.Add(1) // Tambahkan 1 untuk goroutine penutup channel ke WaitGroup utama
	go func() {
		defer wg.Done()
		workerWg.Wait()
		close(outputChan)
		fmt.Println("Penutup Output Channel: Output channel telah ditutup.")
	}()
	//---- Goroutine Pekerja ----

	//---- Goroutine Pencetak Hasil ----
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("Semua data sudah dicetak")

		for v := range outputChan {
			fmt.Printf("Pencetak: menerima hasil %d\n", v)
			time.Sleep(100 * time.Millisecond)
		}

	}()
	//---- Goroutine Pencetak Hasil ----
	wg.Wait()

	fmt.Println(totalProcessed)
	if totalProcessed < 20 {
		fmt.Println("Ada Kesalahan Kode")
		fmt.Println("Latihan Reviews Part 1 : GAGAL")
	} else {
		fmt.Println("Latihan Reviews Part 1 : LULUS")
	}
}
