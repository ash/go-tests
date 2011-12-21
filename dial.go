package main

import "fmt"
import "net"

func main() {
    conn, err := net.Dial("tcp", "89.149.254.47:80")
    fmt.Printf("%v\n%v\n", conn, err)
    
    buf := make([]byte, 1024)
    
    conn.Write([]byte("GET / HTTP/1.1\nHost: shitov.ru\n\n"))
    for {
	n, _ := conn.Read(buf)
	if (n == 0) {
	    break
	}
	fmt.Printf("%s", buf[0:n])
    }
    conn.Close()
}
