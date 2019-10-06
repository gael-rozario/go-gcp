package instances

import (
	"context"
	"log"
	compute "google.golang.org/api/compute/v1"
)

//GetAllInstances list all instances in the project
func GetAllInstances(project string) ([]*compute.Instance, error){
	instanceList := make([]*compute.Instance,0)
	ctx := context.Background()
	computeclient, err := compute.NewService(ctx)
	zones := GetZones(project)
	if err != nil {
		return nil, err
	}
	for _, zone := range zones {
		response, err := computeclient.Instances.List(project, zone).Do()
		if err != nil {
			return nil, err
		}
		for _,instance := range response.Items{
			instanceList= append(instanceList,instance)
		}
	}
	return instanceList, nil
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
