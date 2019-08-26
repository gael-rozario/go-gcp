package main

import (
	"fmt"
	"log"

	"github.com/gael-rozario/go-gcp/instances"
)

func main() {
	project := "angelic-bond-246708"
	instancelist, err := instances.GetAllInstances(project)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%#v", instancelist.Items)
	for _, instance := range instancelist.Items{
		for _, interface := range instance.NetworkInterfaces{
			
		}
	}
}
