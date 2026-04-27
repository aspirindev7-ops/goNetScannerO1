package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

// Result represents the status of a single port
type Result struct {
	Port    int
	State   string
	Service string
}

func main() {
	target := "127.0.0.1" // Localhost for safety, can be changed to any IP
	startPort := 1
	endPort := 1024

	fmt.Printf("--- Go-Scan: High Performance Network Discovery ---\n")
	fmt.Printf("Scanning %s (Ports %d-%d)...\n", target, startPort, endPort)
	start := time.Now()

	var wg sync.WaitGroup
	results := make(chan Result, endPort-startPort+1)

	// Worker pool to handle concurrency
	for i := startPort; i <= endPort; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			address := fmt.Sprintf("%s:%d", target, port)
			
			// DialTimeout is used to efficiently check for open ports
			conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
			if err != nil {
				return
			}
			conn.Close()
			results <- Result{Port: port, State: "Open"}
		}(i)
	}

	// Wait for all goroutines and close channel
	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Printf("%-10s | %-10s\n", "PORT", "STATE")
	fmt.Println("-----------------------")
	for res := range results {
		fmt.Printf("%-10d | %-10s\n", res.Port, res.State)
	}

	duration := time.Since(start)
	fmt.Printf("\nScan completed in: %v\n", duration)
}
