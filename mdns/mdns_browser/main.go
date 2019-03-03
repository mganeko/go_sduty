package main

// https://github.com/hashicorp/mdns


// /*---
import (
	"fmt"
	"sync"
	"github.com/hashicorp/mdns"
)

func main() {
	wg := &sync.WaitGroup{} 

	entriesCh := make(chan *mdns.ServiceEntry, 4)
	wg.Add(1)
	go func() {
		for entry := range entriesCh {
			fmt.Printf("Got new entry: %v\n", entry)
		}
		wg.Done()
	}()

	// Start the lookup
	fmt.Printf(" -- start lookup --\r\n")
	//mdns.Lookup("_irkit._tcp", entriesCh) // OK
	//mdns.Lookup("_airplay._tcp", entriesCh) // OK
	//mdns.Lookup("_googlecast._tcp", entriesCh) // OK
	//mdns.Lookup("_smb._tcp", entriesCh) // OK
	mdns.Lookup("_foobar._tcp", entriesCh) // OK

	//mdns.Lookup("_http._tcp", entriesCh)
	wg.Wait()
	fmt.Printf(" -- end lookup --\r\n")
	close(entriesCh)
}

//--*/


/*----
import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/grandcat/zeroconf"
)

var (
	//service  = flag.String("service", "_workstation._tcp", "Set the service category to look for devices.")
	service  = flag.String("service", "_irkit._tcp", "Set the service category to look for devices.")
	domain   = flag.String("domain", "local", "Set the search domain. For local networks, default is fine.")
	waitTime = flag.Int("wait", 10, "Duration in [s] to run discovery.")
)

// MAC
//  dns-sd -B _irkit._tcp     <-- OK
//   dns-sd -B _http._tcp

func main() {
	flag.Parse()

	// Discover all services on the network (e.g. _workstation._tcp)
	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		log.Fatalln("Failed to initialize resolver:", err.Error())
	}

	entries := make(chan *zeroconf.ServiceEntry)
	go func(results <-chan *zeroconf.ServiceEntry) {
		for entry := range results {
			log.Println(entry)
		}
		log.Println("No more entries.")
	}(entries)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(*waitTime))
	defer cancel()
	err = resolver.Browse(ctx, *service, *domain, entries)
	if err != nil {
		log.Fatalln("Failed to browse:", err.Error())
	}

	<-ctx.Done()
	// Wait some additional time to see debug messages on go routine shutdown.
	time.Sleep(1 * time.Second)
}

---*/
