package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func checkSite(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	start := time.Now()
	resp, err := http.Get(url)
	duration := time.Since(start)

	if err != nil {
		fmt.Printf("❌ %s is down! (%v)\n", url, err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("✅ %s returned %d in %v\n", url, resp.StatusCode, duration)
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://github.com",
		"https://google.com",
		"https://httpbin.org/delay/2",
	}

	var wg sync.WaitGroup

	fmt.Println("Starting health checks...")
	start := time.Now()

	for _, url := range urls {
		wg.Add(1)
		go checkSite(url, &wg)
	}

	wg.Wait()
	fmt.Printf("All checks completed in %v!\n", time.Since(start))
}
