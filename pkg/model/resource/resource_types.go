// The resource package contains definitions within the context of resource
package resource

import "github.com/turbonomic/turbo-go-monitoring/pkg/model"

// List of resource types
const (
	CPU      	model.ResourceType = "CPU"
	MEM      	model.ResourceType = "MEM"
	CPU_LIMIT	model.ResourceType = "CPU_LIMIT"
	MEM_LIMIT	model.ResourceType = "MEM_LIMIT"
	MEM_REQ  	model.ResourceType = "MEM_REQ"
	CPU_REQ  	model.ResourceType = "CPU_REQ"
	MEM_PROV 	model.ResourceType = "MEM_PROV"
	CPU_PROV 	model.ResourceType = "CPU_PROV"
	NETWORK  	model.ResourceType = "NETWORK"
	DISK     	model.ResourceType = "DISK"
	OBJECT_COUNT 	model.ResourceType = "OBJECT_COUNT"
	TRANSACTIONS	model.ResourceType = "TRANSACTIONS"
)


const (
	CLUSTER		model.ConstraintType = "Cluster"
	VMPM		model.ConstraintType = "VMPM"
	APPLICATION	model.ConstraintType = "APPLICATION"
)
