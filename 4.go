package main
import (
	"fmt"
	"net/http"
	"sync"
)
func main() {
	urls := []string{
		"https://golang.org",
		"https://google.com",
		"https://github.com",
		"https://yandex.ru",
		"https://vk.com",
	}
	s := make(chan struct{}, 3)
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			s <- struct{}{}
			defer func() { <-s }()

			resp, err := http.Get(u)
			if err != nil {
				fmt.Println(u, "ошибка:", err)
				return
			}
			defer resp.Body.Close()
			fmt.Println(u, resp.Status)
		}(url)
	}
	wg.Wait()
}
