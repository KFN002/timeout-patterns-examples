package main

import (
	"fmt"
	"net/http"
	"time"
)

func fetchURL(url string, c chan string) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		c <- fmt.Sprintf("Ошибка при получении %s: %s", url, err)
		return
	}
	defer resp.Body.Close()
	c <- fmt.Sprintf("Ответ от %s: Статус - %s", url, resp.Status)
}
func main() {
	urls := []string{
		"https://yandex.ru",
		"https://lyceum.yandex.ru",
		"https://translate.yandex.com",
		"https://ihumaunkabir.com",
	}
	c := make(chan string, len(urls))
	for _, url := range urls {
		go fetchURL(url, c)
	}
	timeout := time.After(15 * time.Second)
	for i := 0; i < len(urls); i++ {
		select {
		case result := <-c:
			fmt.Println(result)
		case <-timeout:
			fmt.Println("Произошел таймаут. Прерывание остальных запросов.")
			return
		}
	}
}
