package repository

import (
	"github.com/turbonomic/turbo-go-monitoring/pkg/model"
	"github.com/turbonomic/turbo-go-monitoring/pkg/model/resource"
	"github.com/turbonomic/turbo-go-monitoring/pkg/model/property"
	"testing"
	"fmt"
)

var TestResource = []struct {
	resourceType model.ResourceType
	key string
	propType     model.MetricPropType
	value        float64
}{
	{resource.CPU, "", property.USED, 10.1},
	{resource.CPU, "", property.PEAK, 46.7},
	{resource.MEM, "", property.USED, 62.3},
	{resource.MEM, "", property.AVERAGE, 43.4},
	{resource.MEM_PROV, "mem1", property.CAP, 87.9},
	{resource.DISK, "disk1", property.CAP, 100.0},
	{resource.CPU_PROV, "cpu1", property.CAP, 90.5},
}

func TestResourceMap(t *testing.T) {

	resourceMap := &EntityResourceMap{}
	// Add all test metrics into the metric map
	//
	idMap := make(map[model.ResourceType]*ResourceIdentifier)
	for _, testResource := range TestResource {
		_, exists := idMap[testResource.resourceType]
		if !exists {
			var resourceId *ResourceIdentifier
			if testResource.key == "" {
				resourceId = NewResourceIdentifier(testResource.resourceType)
			} else {
				resourceId = NewResourceIdentifierWithKey(testResource.resourceType, testResource.key)
			}
			idMap[testResource.resourceType] = resourceId
		}
	}

	for _, testResource := range TestResource {
		resourceId := idMap[testResource.resourceType]
		resourceMap.SetMetricValue(resourceId, testResource.propType, model.MetricValue(testResource.value))
	}

	fmt.Printf("##### Resource Map\n %++v ##### \n", resourceMap)
	//
	// Retrieve the value for each metric and confirm it's the same as entered
	//
	for _, testResource := range TestResource {
		resourceId := idMap[testResource.resourceType]
		value, err := resourceMap.GetMetricValue(resourceId, testResource.propType)
		if err != nil {
			t.Errorf("Error while retrieving metric (%v, %v) from map %v: %s",
				testResource.resourceType, testResource.propType, resourceMap, err)
		}
		if value != model.MetricValue(testResource.value) {
			t.Errorf("Retrieved value %v of metric (%v, %v) from metric map %v is not the same as entered %v",
				value, testResource.resourceType, testResource.propType, resourceMap, testResource.value)

		}
	}
	//
	// Try to fetch the value of a non-existing metric
	//
	fakeResourceType := model.ResourceType("NOT_EXIST_RESOURCE_TYPE")
	resourceId := NewResourceIdentifier(fakeResourceType)
	fakeMetricPropType := model.MetricPropType("NOT_EXIST_PROP_TYPE")
	_, err := resourceMap.GetMetricValue(resourceId, fakeMetricPropType)
	if err == nil {
		t.Errorf("Expecting error but getting no error, when retrieving metric value of resource %v and "+
			"property %v from metric map %v", fakeResourceType, fakeMetricPropType, resourceMap)
	}
}
