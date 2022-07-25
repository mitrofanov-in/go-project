package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	urls := []string{
		"https://www.boostra.ru",
		"https://fssp.gov.ru",
	}

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			doHTTP(url)
			wg.Done()
		}(url)
	}

	wg.Wait()
}

func doHTTP(url string) {
	t := time.Now()

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Смотрим смторим", url, err.Error())
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println("Статус страницы", url, time.Since(t).Milliseconds())

}
