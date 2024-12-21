package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

func checkStatus(domain string, port string) string {
	address := domain + ":" + port
	timeout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp", address, timeout)
	var status string

	if err != nil {
		status = "closed"
		status = fmt.Sprintf("[DOWN] Connection to %s on port %s failed: %s\n", domain, port, err)
	} else {
		status = "open"

		status = fmt.Sprintf("[UP] Successfully connected\nLocal connection %s\nRemote Connection %s\n", conn.LocalAddr(), conn.RemoteAddr())
		conn.Close()
	}
	return status
}

func main() {
	domain := flag.String("domain", "example.com", "Domain to scan")
	port := flag.String("port", "80", "Port to scan")
	flag.Parse()

	fmt.Printf("Scanning %s on port %s\n", *domain, *port)
	fmt.Printf(checkStatus(*domain, *port))

}
