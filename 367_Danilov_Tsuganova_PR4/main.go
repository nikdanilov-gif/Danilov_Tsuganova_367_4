package main

import ("fmt"
		"time"
		"sync")

func main(){
	var wg sync.WaitGroup
	wg.Add(1)
	go func(){
		defer wg.Done()
		for i:=1; i<=5;i++{
			fmt.Println(i)	
			time.Sleep(1*time.Second)
		}
	}()
	wg.Wait()
}
