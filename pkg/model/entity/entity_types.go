// The entity package contains definitions within the context of entity.
package entity

import "github.com/turbonomic/turbo-go-monitoring/pkg/model"

// List of entity types
const (
	NODE model.EntityType = "Node"
	POD  model.EntityType = "Pod"
	APP  model.EntityType = "App"
	VIRTUAL_CLUSTER model.EntityType = "VirtualCluster"
	CONTAINER model.EntityType = "Container"
	SERVICE model.EntityType = "Service"
	LOAD_BALANCER model.EntityType = "LoadBalancer"
)
