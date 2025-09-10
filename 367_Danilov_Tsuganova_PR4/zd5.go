package main

import (
 "crypto/md5"
 "fmt"
 "sync"
)
func main() {
 data := []string{"t1.txt", "t2.txt", "t3.txt", "t4.txt", "t5.txt"}
 limit := make(chan bool, 3)
 var wd sync.WaitGroup
 for _, text := range data {
  wd.Add(1)
  limit <- true 
  go func(x string) {
   defer wd.Done()
   defer func() { <-limit }() 
   hash := md5.Sum([]byte(x))
   fmt.Printf("Хеш '%s': %x\n", x, hash)
  }(text)
 }
 wd.Wait()
}