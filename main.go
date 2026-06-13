package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
	"time"
)

func argsHandler() (string, string, []string) {
	host := flag.String("host", "localhost", "Hôte du serveur")
	portsStr := flag.String("port", "8080", "Liste des ports séparés par des virgules (ex: 8080,8081,9000)")
	datatype := flag.String("data", "tcp", "TCP/UDP")
	flag.Parse()
	portsList := strings.Split(*portsStr, ",")
	return *datatype, *host, portsList
}

func connectToPort(url string, port string, data_type string) (net.Conn, error) {
	start := time.Now()
	conn, err := net.Dial(data_type, url+":"+port)
	elapsed := time.Since(start)

	if err != nil {
		fmt.Println("Error connecting:", err)
		return nil, err
	}

	fmt.Println("OK -", url+":"+port, "took", elapsed)

	return conn, nil
}

func startListeners(ports []string, host string, datatype string) {
	// _ index of ports list declared but unused
	for _, port := range ports {
		_, err := connectToPort(host, port, datatype)
		if err != nil {
			continue
		}
	}
}

func main() {
	datatype, host, ports := argsHandler()

	startListeners(ports, host, datatype)
}
