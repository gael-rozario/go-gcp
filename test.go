package main

import (
	"github.com/gael-rozario/go-gcp/instances"
	"log"
	"fmt"
)

func main() {
	project := "angelic-bond-246708"
	result,err  := instances.GetAllInstances(project)
	if err!= nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
