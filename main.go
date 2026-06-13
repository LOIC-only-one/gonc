package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

func argsHandler() (string, string, []string, string) {
	host := flag.String("host", "localhost", "Hôte du serveur")
	portsStr := flag.String("port", "8080", "Liste ports séparés par virgules")
	datatype := flag.String("data", "tcp", "tcp|udp")
	mode := flag.String("mode", "scan", "scan|listen")
	flag.Parse()

	portsList := strings.Split(*portsStr, ",")
	return *datatype, *host, portsList, *mode
}

func connectToPort(host string, port string, dataType string, wg *sync.WaitGroup) {
	defer wg.Done()

	if dataType == "udp" {
		addr := host + ":" + port
		conn, err := net.DialTimeout("udp", addr, 500*time.Millisecond)
		if err != nil {
			fmt.Println("FAIL -", addr)
			return
		}
		defer conn.Close()

		conn.SetDeadline(time.Now().Add(500 * time.Millisecond))
		_, err = conn.Write([]byte("ping"))
		if err != nil {
			fmt.Println("FAIL -", addr)
			return
		}

		buf := make([]byte, 32)
		_, err = conn.Read(buf)
		if err != nil {
			fmt.Println("OPEN|FILTERED -", addr)
			return
		}
		fmt.Println("OK   -", addr)
		return
	}

	start := time.Now()
	conn, err := net.DialTimeout(dataType, host+":"+port, 500*time.Millisecond)
	elapsed := time.Since(start)
	if err != nil {
		fmt.Println("FAIL -", host+":"+port)
		return
	}
	conn.Close()
	fmt.Println("OK   -", host+":"+port, "(", elapsed, ")")
}

func scanPorts(host string, ports []string, dataType string) {
	var wg sync.WaitGroup

	for _, port := range ports {
		wg.Add(1)
		go connectToPort(host, strings.TrimSpace(port), dataType, &wg)
	}

	wg.Wait()
}

func startListener(ports []string) {
	var wg sync.WaitGroup

	for _, port := range ports {
		wg.Add(1)
		go func(p string) {
			defer wg.Done()

			listener, err := net.Listen("tcp", ":"+p)
			if err != nil {
				fmt.Println("Listen error on port", p, ":", err)
				return
			}
			defer listener.Close()
			fmt.Println("OK - listening on port:", p)

			for {
				conn, err := listener.Accept()
				if err != nil {
					fmt.Println("Accept error:", err)
					continue
				}
				fmt.Println("OK - conn from:", conn.RemoteAddr().String(), "on port:", p)
				conn.Close()
			}
		}(port)
	}

	wg.Wait()
}

func main() {
	dataType, host, ports, mode := argsHandler()

	if mode == "listen" {
		startListener(ports)
	} else {
		scanPorts(host, ports, dataType)
	}
}
