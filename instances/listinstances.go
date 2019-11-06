package instances

import (
	"context"
	"log"
	compute "google.golang.org/api/compute/v1"
	"sync"
)

var wg sync.WaitGroup
//GetAllInstances list all instances in the project
func GetAllInstances(project string) ([]*compute.Instance, error){
	instanceList := make([]*compute.Instance,0)
	instanceChannel := make(chan *compute.Instance, 100)
	zones := GetZones(project)
	for _, zone := range zones {
		wg.Add(1)
		go GetInstanceOverRegion(instanceChannel,project,zone)
	}
	wg.Wait()
	close(instanceChannel)
	for instance:= range instanceChannel{
		instanceList = append(instanceList,instance)	
	}

	return instanceList, nil
}

//GetInstanceOverRegion retrieves instances in one region
func GetInstanceOverRegion(instanceChannel chan *compute.Instance,project string, zone string)  {
	defer wg.Done()
	ctx := context.Background()
	computeclient, err := compute.NewService(ctx)
	if err != nil {
		log.Fatal(err)
	}
	response, err := computeclient.Instances.List(project, zone).Do()
	if err!= nil{
		log.Fatal(err)
	}
	for _,instance := range response.Items{
		instanceChannel <-instance
	}
	
}
//GetZones lists zones available in a project
func GetZones(p string) []string {
	ctx := context.Background()
	computeclient, err := compute.NewService(ctx)
	var zones []string
	if err != nil {
		log.Fatal(err)
	}
	response, err := computeclient.Zones.List(p).Do()
	if err != nil {
		log.Fatal(err)
	}
	for _, zone := range response.Items {
		zones = append(zones, zone.Name)
	}
	return zones
}
