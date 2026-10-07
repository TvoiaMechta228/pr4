package main
import (
	"fmt"
	"sync"
)
func main() {
	job := make(chan int, 10)
	resalt := make(chan int, 10)
	var wg sync.WaitGroup
	for w := 0; w < 3; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range job {
				resalt <- j * j
			}
		}()
	}
	for i := 1; i <= 10; i++ {
		job<- i
	}
	close(job)
	wg.Wait()
	close(resalt)
	for r := range resalt {
		fmt.Println(r)
	}
}
