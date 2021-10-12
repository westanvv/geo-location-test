package main

import (
	"github.com/oschwald/geoip2-golang" //nolint:goimports
	"net"
)

func main() {
}

type GeoLocation struct {
	Latitude  float64
	Longitude float64
	IP        string
}

func getGeoInfo(
	db *geoip2.Reader,
	ipAddress string,
) (GeoLocation, error) {
	ip := net.ParseIP(ipAddress)
	record, err := db.City(ip)

	var geoLocation GeoLocation

	geoLocation.IP = ipAddress
	geoLocation.Latitude = record.Location.Latitude
	geoLocation.Longitude = record.Location.Longitude

	return geoLocation, err
}
