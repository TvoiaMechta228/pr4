package main
import (
	"fmt"
)
func search(source string, query string, resalt chan<- string) {
	resalt <- fmt.Sprintf("найдено в %s: %s", source, query)
}
func main() {
	sources := []string{"БД1", "БД2", "API1", "API2"}
	query := "test"

	resalt := make(chan string, len(sources))

	for _, s := range sources {
		go search(s, query, resalt)
	}

	fmt.Println(<-resalt)
}
