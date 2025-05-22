package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var dataChan = make(chan string)

	wg.Add(2)
	go func() {
		defer wg.Done()
		dataChan <- "Data Rahasia"
		fmt.Println("Data terkirim")
	}()
	go func() {
		defer wg.Done()
		data := <-dataChan
		fmt.Printf("Data diterima: %s", data)
	}()
	wg.Wait()

}
