package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"sync"
)
func main() {
	files := []string{"1.go", "2.go", "3.go", "4.go"}

	jobs := make(chan string, len(files))
	resalt := make(chan string, len(files))
	var wg sync.WaitGroup
	for w := 0; w < 3; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for f := range jobs {
			file, err := os.Open(f)
			if err != nil {
				resalt <- fmt.Sprintf("%s: ошибка %v", f, err)
				continue
				}
				h := md5.New()
				io.Copy(h, file)
				file.Close()
				resalt <- fmt.Sprintf("%s: %x", f, h.Sum(nil))
			}
		}()
	}
	for _, f := range files {
		jobs <- f
	}
	close(jobs)
	wg.Wait()
	close(resalt)
	for r := range resalt {
		fmt.Println(r)
	}
}
