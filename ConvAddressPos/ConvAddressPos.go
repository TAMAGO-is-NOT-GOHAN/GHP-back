package ConvAddressPos

import (
	"context"
	"log"
	"os"

	"googlemaps.github.io/maps"
)

func ConvAddressToPos(address string) Coordinate {
	c, err := maps.NewClient(maps.WithAPIKey(os.Getenv("GMAP_TOKEN")))
	if err != nil {
		log.Fatal(err)
	}

	r := &maps.GeocodingRequest{
		Address: address,
	}

	route, err := c.Geocode(context.Background(), r)
	if err != nil {
		log.Fatal(err)
	}

	coordinate := Coordinate{address, route[0].Geometry.Location.Lat, route[0].Geometry.Location.Lng}

	return coordinate
}
