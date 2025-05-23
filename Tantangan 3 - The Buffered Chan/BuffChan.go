package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var dataChan = make(chan string, 2) // buffered channel dengan kapasitas 2

	// ----- Gouroutine Pengirim -----
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(dataChan)
		defer fmt.Println("Pengirim Selesai mengirim semua data") // pesan ini akan muncul ketika semua data sudah terkirim dan diterima

		mesagges := []string{"data 1", "data 2", "data 3"} // slice data yang dikirim

		for _, msg := range mesagges {
			fmt.Printf("Pengirim: mencoba mengirim pesan '%s'\n", msg)
			dataChan <- msg
			fmt.Printf("Pengirim telah mengirimkan pesan '%s'\n", msg)
			time.Sleep(50 * time.Millisecond)
		}
	}()
	// ----- Gouroutine Pengirim -----

	// ----- Go routine Penerima -----
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("Channel telah ditutup, selesai membaca")

		time.Sleep(100 * time.Millisecond) // beri waktu pengirim untuk mengisi buffer
		fmt.Println("\nPenerima Mulai membaca data")

		for msg := range dataChan {
			fmt.Printf("Penerima: menerima pesan '%s'\n", msg)
			time.Sleep(100 * time.Millisecond) // untuk simulasi pproses
		}

	}()
	// ----- Go routine Penerima -----

	wg.Wait()
	fmt.Println("Progam Selesai")
}
