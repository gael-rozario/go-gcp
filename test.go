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
	for _, instance := range instancelist {
		fmt.Println(instance)
	}
}
