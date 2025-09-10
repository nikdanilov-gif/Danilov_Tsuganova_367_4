package main

import (
 "fmt"
 "net/http"
 "sync"
 "time"
)
func main() {
 urls := []string{
  "https://youtube.com",
  "https://www.dns-shop.ru",
  "https://market.yandex.ru",
  "https://flk.vostok-electra.ru/",
  "https://www.google.com/",
  "https://github.com",
  "https://workspace.ktk-45.ru",
 }

 jobs := make(chan string, len(urls))
 results := make(chan string, len(urls))

 var wd sync.WaitGroup

 for x := 1; x <= 3; x++ {
  wd.Add(1)
  go func(workerID int) {
   defer wd.Done()

   for url := range jobs {
    fmt.Printf("Проверка %d статуса: %s\n", workerID, url)
    

    client := http.Client{
     Timeout: 5 * time.Second,
    }

    resp, err := client.Get(url)
    var status string
    if err != nil {
     status = fmt.Sprintf("ОШИБКА: %v", err)
    } else {
     defer resp.Body.Close()
     status = fmt.Sprintf("Статус: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
    }
    results <- fmt.Sprintf("Итог:%d: %s - %s", workerID, url, status)
   }
  }(x)
 }

 for _, url := range urls {
  jobs <- url
 }
 close(jobs)
 wd.Wait()

 close(results)

 fmt.Println("\n URL:")
 for result := range results {
  fmt.Println(result)
 }
}