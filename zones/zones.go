package zones

import (
	"context"
	"fmt"
	"log"

	compute "google.golang.org/api/compute/v1"
)

//GetZones lists zones available in a project
func GetZones(p string) {
	ctx := context.Background()
	computeclient, err := compute.NewService(ctx)
	if err != nil {
		log.Fatal(err)
	}
	res, err := computeclient.Zones.List(p).Do()
	if err != nil {
		log.Fatal(err)
	}
	for _, zone := range res.Items {
		fmt.Printf("%#v \n", zone.Name)
	}
}
