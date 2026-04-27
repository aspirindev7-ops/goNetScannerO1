# Go-Scan: High-Performance Concurrent Port Scanner

Go-Scan is a light-weight, blazing-fast network port discovery tool built in Go. This project demonstrates the power of Go's concurrency model (Goroutines and Channels) applied to network engineering tasks.

## Why Go?
Go is the language of modern cloud infrastructure (Docker, Kubernetes). This project showcases:
- **Goroutines:** Running thousands of network checks in parallel with minimal memory overhead.
- **WaitGroups:** Synchronizing concurrent operations.
- **Networking (net package):** Low-level TCP dialing and timeout management.

## Features
- **Concurrent Execution:** Scans 1000+ ports in seconds.
- **Timeout Management:** Efficiently handles non-responsive hosts.
- **Resource Efficiency:** Uses minimal CPU/RAM compared to traditional scripting.

## Usage
Ensure you have Go installed:
```bash
go run goscan.go
```
