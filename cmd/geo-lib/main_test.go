package main

import (
	"github.com/Pallinder/go-randomdata" //nolint:goimports
	"github.com/oschwald/geoip2-golang"
	"log"
	"os"
	"sync"
	"testing" //nolint:goimports
)

const IPCount = 50000

func BenchmarkWith1_Goroutine(b *testing.B) {
	pwd, _ := os.Getwd()
	db, err := geoip2.Open(pwd + "/../../data/GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for i := 0; i < b.N; i++ {
		test(1, db)
	}
}

func BenchmarkWith10_Goroutine(b *testing.B) {
	pwd, _ := os.Getwd()
	db, err := geoip2.Open(pwd + "/../../data/GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for i := 0; i < b.N; i++ {
		test(10, db)
	}
}

func test(threadCount int, db *geoip2.Reader) {
	var ch = make(chan string, IPCount)
	var wg sync.WaitGroup

	// This starts threadCount number of goroutines that wait for something to do
	wg.Add(threadCount)
	for i := 0; i < threadCount; i++ {
		go func() {
			for {
				ip, ok := <-ch
				if !ok { // if there is nothing to do and the channel has been closed then end the goroutine
					wg.Done()
					return
				}
				getGeoInfo(db, ip)
				//record, _ := getGeoInfo(db, ip)
				//fmt.Printf("%v, [%v, %v]\n", record.IP, record.Latitude, record.Longitude)
			}
		}()
	}

	// Now the jobs can be added to the channel, which is used as a queue
	for i := 0; i < IPCount; i++ {
		ch <- randomdata.IpV4Address() // add IP to the queue
	}

	close(ch) // This tells the goroutines there's nothing else to do
	wg.Wait() // Wait for the threads to finish
}
