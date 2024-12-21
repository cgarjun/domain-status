#### Help
```
go run main.go --help                      
  -domain string
    	Domain to scan (default "example.com")
  -port string
    	Port to scan (default "80")
  -timer int
    	timer in seconds (default 5)
```
#### Quick test
```
go run main.go -domain google.com -timer 10
```
```
Scanning google.com on port 80 with a timer of 10
[UP] Successfully connected
Local connection 192.168.68.124:61050
Remote Connection 142.250.180.14:80
```
