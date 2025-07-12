package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	StatusCode int
	Duration   time.Duration
}

func worker(url string, requests int, wg *sync.WaitGroup, results chan<- Result) {
	defer wg.Done()
	for i := 0; i < requests; i++ {
		start := time.Now()
		resp, err := http.Get(url)
		duration := time.Since(start)

		if err != nil {
			results <- Result{StatusCode: 0, Duration: duration}
			continue
		}
		results <- Result{StatusCode: resp.StatusCode, Duration: duration}
		resp.Body.Close()
	}
}

func main() {
	url := flag.String("url", "", "URL do serviço a ser testado")
	totalRequests := flag.Int("requests", 100, "Número total de requisições")
	concurrency := flag.Int("concurrency", 10, "Número de requisições simultâneas")
	flag.Parse()

	if *url == "" {
		fmt.Println("A URL deve ser informada com o parâmetro --url")
		return
	}

	requestsPerWorker := *totalRequests / *concurrency
	extra := *totalRequests % *concurrency

	startTime := time.Now()

	var wg sync.WaitGroup
	results := make(chan Result, *totalRequests)

	for i := 0; i < *concurrency; i++ {
		reqs := requestsPerWorker
		if i < extra {
			reqs++
		}
		wg.Add(1)
		go worker(*url, reqs, &wg, results)
	}

	wg.Wait()
	close(results)
	totalTime := time.Since(startTime)

	statusCount := make(map[int]int)
	var total200 int
	for r := range results {
		statusCount[r.StatusCode]++
		if r.StatusCode == 200 {
			total200++
		}
	}

	// Relatório
	fmt.Println("========== Relatório ==========")
	fmt.Printf("Tempo total: %v\n", totalTime)
	fmt.Printf("Total de requisições: %d\n", *totalRequests)
	fmt.Printf("Sucesso (200 OK): %d\n", total200)
	fmt.Println("Distribuição de status HTTP:")
	for code, count := range statusCount {
		fmt.Printf("  %d: %d\n", code, count)
	}
	fmt.Println("================================")
}
