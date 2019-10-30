package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"com.crawler/service"
)

func main() {
	https_url := "https://www.qubole.com"
	http_url := "http://www.qubole.com"

	domains := []string{http_url, https_url}
	uniqueUrls := make(map[string]bool, 0)

	// stops crawler with ctrl-c
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)

	go func() {
		<-sigs
		done <- true
	}()

	links := make(chan string)
	cl := service.NewCrawler(domains)

	go cl.Crawl(https_url, links)

	for {
		select {
		case str := <-links:
			if !uniqueUrls[str] {
				fmt.Printf("\n[Domain] %s", str)
			} else {
				uniqueUrls[str] = true
			}

		case <-done:
			fmt.Println("\n\nExiting Crawler...\n")
			os.Exit(0)
		}
	}
}
