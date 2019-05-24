package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// генерация
	go func() {
		for z:= 0; ; z++ {
			naturals <- z
		}
		close(naturals)
	}()
		
	// возведение в квадрат
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// печать(указываем значение v, это число возведения после которого программа остановится)
	for v:=0; v<=10 ;v++ {
		fmt.Println(<-squares)
		time.Sleep(1 * time.Second)
	}
}
