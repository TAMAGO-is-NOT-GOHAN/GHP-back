package shortestroute

import (
	"context"
	"log"
	"os"

	"googlemaps.github.io/maps"
)

func SearchShortestRoute(origin string, destination string) int {
	c, err := maps.NewClient(maps.WithAPIKey(os.Getenv("GMAP_TOKEN")))
	if err != nil {
		log.Fatal(err)
	}
	r := &maps.DirectionsRequest{
		Origin:      origin,
		Destination: destination,
	}
	route, _, err := c.Directions(context.Background(), r)
	if err != nil {
		log.Fatal(err)
	}

	return int(route[0].Legs[0].Duration.Minutes())
}
