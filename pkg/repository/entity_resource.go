package repository

import (
	"github.com/turbonomic/turbo-go-monitoring/pkg/model"
	"fmt"
	"bytes"
	"github.com/golang/glog"
)

type ResourceIdentifier struct {
	ResourceType model.ResourceType
	Key string
}

type MetricKey struct {
	ResourceIdentifier *ResourceIdentifier
	PropType     model.MetricPropType
}

// Represents the individual Commodity or resource that an entity buys or sells in the Repository.
// Each commodity has a set of properties whose value is monitored and obtained from environment
type EntityResource struct {
	ResourceId *ResourceIdentifier
	Props map[model.MetricPropType]model.MetricValue	//*MonitoredMetricProp
}
// Map of Entity Resources and the identifier for the resource
type EntityResourceMap map[*ResourceIdentifier]*EntityResource

// ----------------------------------------------------------------------
func (resourceId *ResourceIdentifier) String() string {
	var buffer bytes.Buffer
	var line string
	if resourceId.Key != "" {
		line = fmt.Sprintf("%s::%s", resourceId.ResourceType, resourceId.Key)
	} else {
		line = fmt.Sprintf("%s::<nil>", resourceId.ResourceType)
	}
	buffer.WriteString(line)
	return buffer.String()
}

func NewResourceIdentifier(resourceType model.ResourceType) *ResourceIdentifier {
	return &ResourceIdentifier{ResourceType: resourceType}
}

func NewResourceIdentifierWithKey(resourceType model.ResourceType,key string) *ResourceIdentifier {
	return &ResourceIdentifier{ResourceType: resourceType, Key: key}
}

// ----------------------------------------------------------------------

func NewEntityResource(resourceId *ResourceIdentifier) *EntityResource {
	return &EntityResource {
		ResourceId: resourceId,
		Props: make(map[model.MetricPropType]model.MetricValue),	//*MonitoredMetricProp),
		}
}

func (resource *EntityResource) SetMetricValue(propType model.MetricPropType, value model.MetricValue) {
	resource.Props[propType] = value
}

func (resource *EntityResource) GetMetricValue(propType model.MetricPropType) (model.MetricValue, error) {
	prop, exists := resource.Props[propType]
	if !exists {
		glog.V(4).Infof("Cannot find metrics for type %s\n", propType)
		return model.MetricValue(0), fmt.Errorf("missing metrics for type %:%s", resource.ResourceId, propType)
	}
	//return prop.Value, nil
	return prop, nil
}

func (resource *EntityResource) String() string {
	var buffer bytes.Buffer
	var line string
	line = fmt.Sprintf("%s", resource.ResourceId)
	buffer.WriteString(line)
	for _, prop := range resource.Props {
		line := fmt.Sprintf("\t%s", prop)
		buffer.WriteString(line)
	}
	line = fmt.Sprintf("\n")
	buffer.WriteString(line)
	return buffer.String()
}

//// The property of a resource whose value is monitored and obtained from environment
//type MonitoredMetricProp struct {
//	PropType model.MetricPropType
//	Value model.MetricValue
//}
//
//func (metricProp *MonitoredMetricProp) String() string {
//	var buffer bytes.Buffer
//	var line string
//	line = fmt.Sprintf("%s::%f", string(metricProp.PropType), metricProp.Value)
//	buffer.WriteString(line)
//	return buffer.String()
//}

// ----------------------------------------------------------------------
// SetMetricValue sets the metric value in the EntityResourceMap for the given resource type and the metric property type
func (resourceMap EntityResourceMap) SetMetricValue(resourceId *ResourceIdentifier,
							propType model.MetricPropType, value model.MetricValue) {
	resource, exists := resourceMap[resourceId]
	if !exists {
		resource = NewEntityResource(resourceId)
		resourceMap[resourceId] = resource
	}
	resource.Props[propType] = value
}

// GetMetricValue retrieves the metric value from the MetricMap for the given resource type and the metric property type
func (resourceMap EntityResourceMap) GetMetricValue(resourceId *ResourceIdentifier,
						propType model.MetricPropType) (model.MetricValue, error) {
	resource, exists := resourceMap[resourceId]
	if !exists {
		glog.V(4).Infof("Cannot find resource %s\n", resourceId)
		return model.MetricValue(0), fmt.Errorf("missing resource %s", resourceId)
	}
	prop, exists := resource.Props[propType]
	if !exists {
		glog.V(4).Infof("Cannot find metrics for type %s\n", propType)
		return model.MetricValue(0), fmt.Errorf("missing metrics for type %s:%s", resourceId, propType)
	}
	return prop, nil
}

func (resourceMap EntityResourceMap) String() string {
	var buffer bytes.Buffer
	for _, resource := range resourceMap {
		line := fmt.Sprintf("\t%s", resource.String())
		buffer.WriteString(line)
	}
	return buffer.String()
}

// --------------------------------------------------------------------------------

// Map of Entity Resources and the identifier for the resource
type SoldResourceMap EntityResourceMap

// Map of Entity Resources and the identifier for the resource
type BoughtResourceMap map[string]EntityResourceMap

