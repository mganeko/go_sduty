package main

// https://github.com/hashicorp/mdns

// --- register

import (
	"fmt"
	//"sync"
	"time"
	"os"
	"github.com/hashicorp/mdns"
)

func main() {
	// Setup our service export
	host, _ := os.Hostname()
	info := []string{"My awesome service"}
	service, _ := mdns.NewMDNSService(host, "_foobar._tcp", "", "", 8000, nil, info)

	// Create the mDNS server, defer shutdown
	for {
		server, _ := mdns.NewServer(&mdns.Config{Zone: service})
		defer server.Shutdown()

		fmt.Printf("start msdn register _foobar._tcp\r\n")
		time.Sleep(30 * time.Second)
		break
	}

	fmt.Printf("end msdn register _foobar._tcp\r\n")
}

