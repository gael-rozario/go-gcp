package instances

import (
	"context"

	compute "google.golang.org/api/compute/v1"
)

//GetAllInstances list all instances in the project
func GetAllInstances(project string) (list *compute.InstanceList, err error) {
	ctx := context.Background()
	computeclient, err := compute.NewService(ctx)
	if err != nil {
		return nil, err
	}
	res, err := computeclient.Instances.List(project, "us-east1-b").Do()
	if err != nil {
		return nil, err
	}
	return res, err
}
