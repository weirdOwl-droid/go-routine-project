package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var pesanChan = make(chan string)

	// --- Goroutine Pengirim ---
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(pesanChan)
		defer fmt.Println("\nPengirim: Semua pesan telah terkirim dan channel ditutup.")

		for i := 1; i <= 5; i++ {
			msg := fmt.Sprintf("Pesan ke-%d", i)
			pesanChan <- msg
			fmt.Printf("Pengirim: Mengirim '%s'\n", msg)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	// --- Goroutine Pengirim ---

	// --- Goroutine Penerima ---
	// --- Goroutine Penerima 1 ---
	wg.Add(1)
	go func(id int) {
		defer wg.Done()
		for msg := range pesanChan {
			fmt.Printf("Penerima %d: Menerima '%s'\n", id, msg)
			time.Sleep(150 * time.Millisecond)
		}
		fmt.Printf("Penerima %d: Channel tertutup, berhenti menerima.\n", id)
	}(1)

	// --- Goroutine Penerima 2 ---
	wg.Add(1)
	go func(id int) {
		defer wg.Done()
		for msg := range pesanChan {
			fmt.Printf("Penerima %d: Menerima '%s'\n", id, msg)
			time.Sleep(150 * time.Millisecond)
		}
		fmt.Printf("Penerima %d: Channel tertutup, berhenti menerima.\n", id)
	}(2)
	// --- Goroutine Penerima 2 ---
	// --- Goroutine Penerima ---

	wg.Wait()
	fmt.Println("Program utama selesai.") // Pesan penutup untuk main
}
