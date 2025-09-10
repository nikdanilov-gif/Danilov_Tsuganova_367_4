package main

import ("fmt"
		"time")


func main(){
tick := time.Tick(200 * time.Millisecond)
initTime := time.Now()
ogranich:=make(chan int, 15)
	for zapros:=1; zapros<=15; zapros++{
	<-tick
	fmt.Println("Запрос №",zapros,"выполнен")
	ogranich<-zapros
	}
fmt.Println("Затрачено времени:", time.Since(initTime))
}