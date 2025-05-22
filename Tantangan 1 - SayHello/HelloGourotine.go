package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// ----------------------------------------------------------//
func sayHello() {
	defer wg.Done()
	fmt.Println("Hello, Sayang... 😊")
}
func sayTanya1() {
	fmt.Println("Sudah makan belum?")
}
func sayTanya2() {
	fmt.Println("Kita mau kemana hari ini?")
}

//----------------------------------------------------------//

func main() {

	//----------------------------------------------------------//
	wg.Add(2)
	go sayHello() // konkuren dengan sayTanya
	go func() {
		defer wg.Done()
		sayTanya1()
		sayTanya2()
	}()
	wg.Wait()

	//*Output akan menjadi:
	/*
		Sudah makan belum?			// sayTanya1 akan selalu diatas sayTanya2
		Kita mau kemana hari ini?	// tapi akan konkuren dengan sayHello
		Hello, Sayang... 😊
	*/
	//*Atau
	/*
		Hello, Sayang... 😊
		Sudah makan belum?
		Kita mau kemana hari ini?
	*/
	//----------------------------------------------------------//

}
