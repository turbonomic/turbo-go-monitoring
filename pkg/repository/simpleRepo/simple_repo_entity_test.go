package simpleRepo

import (
	"testing"
	"github.com/turbonomic/turbo-go-monitoring/pkg/model/entity"
	"github.com/turbonomic/turbo-go-monitoring/pkg/model"
	"github.com/turbonomic/turbo-go-monitoring/pkg/repository"
	"github.com/turbonomic/turbo-go-monitoring/pkg/model/resource"
	"github.com/turbonomic/turbo-go-monitoring/pkg/model/property"
)

var TestEntities = []struct {
	entityType model.EntityType
	entityId   model.EntityId
}{
	{entity.NODE, "1.2.3.4"},
	{entity.NODE, "192.168.99.100"},
	{entity.POD, "abc"},
	{entity.APP, "xyz"},
}

func TestSimpleMetricRepoEntity_GetId_GetType(t *testing.T) {

	for _, testEntity := range TestEntities {
		repoEntity := NewSimpleMetricRepoEntity(testEntity.entityType, testEntity.entityId)
		if repoEntity.GetType() != testEntity.entityType {
			t.Errorf("Retrieved type %v from repo entity %v is not the same as input %v",
				repoEntity.GetType(), repoEntity, testEntity.entityType)
		}
		if repoEntity.GetId() != testEntity.entityId {
			t.Errorf("Retrieved id %v from repo entity %v is not the same as input %v",
				repoEntity.GetId(), repoEntity, testEntity.entityId)
		}
	}
}

func TestSimpleMetricRepoEntity_GetSold(t *testing.T) {
	var cpuCap, cpuUsed, memCap, memUsed model.MetricValue
	cpuCap = 100.0
	cpuUsed = 50.0
	memCap = 4.0
	memUsed = 2.0

	resourceId1 := &repository.ResourceIdentifier{
		ResourceType: resource.CPU,
		Key: "",
	}
	prop1 := make(map[model.MetricPropType]*repository.MonitoredMetricProp)
	prop1[property.CAP] = &repository.MonitoredMetricProp{PropType: property.CAP, Value: cpuCap,}
	prop1[property.USED] = &repository.MonitoredMetricProp{PropType: property.USED, Value: cpuUsed,}

	resourceId2 := &repository.ResourceIdentifier{
		ResourceType: resource.MEM,
		Key: "",
	}
	prop2 := make(map[model.MetricPropType]*repository.MonitoredMetricProp)
	prop2[property.CAP] = &repository.MonitoredMetricProp{PropType: property.CAP, Value: memCap,}
	prop2[property.USED] = &repository.MonitoredMetricProp{PropType: property.USED, Value: memUsed,}

	resourceId3 := &repository.ResourceIdentifier{
		ResourceType: resource.CPU_PROV,
		Key: "node1",
	}
	prop3 := make(map[model.MetricPropType]*repository.MonitoredMetricProp)
	prop3[property.USED] = &repository.MonitoredMetricProp{PropType: property.USED, Value: cpuUsed,}

	resourceId4 := &repository.ResourceIdentifier{
		ResourceType: resource.MEM_PROV,
		Key: "node1",
	}
	prop4 := make(map[model.MetricPropType]*repository.MonitoredMetricProp)
	prop4[property.USED] = &repository.MonitoredMetricProp{PropType: property.USED, Value: memUsed,}

	repoEntity := NewSimpleMetricRepoEntity(entity.POD, "pod")

	providerType1 := entity.NODE
	provider1 := "Node1"

	repoEntity.SetBoughtResourceValue(providerType1, model.EntityId(provider1), resourceId3, property.USED, cpuUsed)
	repoEntity.SetBoughtResourceValue(providerType1, model.EntityId(provider1), resourceId4, property.USED, memUsed)
	resource31, _ := repoEntity.GetBoughtResource(providerType1, model.EntityId(provider1), resourceId3)

	used, _ := resource31.GetMetricValue(property.USED)
	if used != cpuUsed {
		t.Errorf("Retrieved bought resource metric %v from repo entity %v is not the same as input %v",
			used, repoEntity, cpuUsed)
	}

	resource41, _ := repoEntity.GetBoughtResource(providerType1, model.EntityId(provider1), resourceId4)

	used, _ = resource41.GetMetricValue(property.USED)
	if used != memUsed {
		t.Errorf("Retrieved bought resource metric %v from repo entity %v is not the same as input %v",
			used, repoEntity, memUsed)
	}

	repoEntity.SetSoldResourceValue(resourceId1, property.CAP, cpuCap)
	repoEntity.SetSoldResourceValue(resourceId1, property.USED, cpuUsed)

	cap, _ := repoEntity.soldMetricMap.GetMetricValue(resourceId1, property.CAP)
	if cap != cpuCap {
		t.Errorf("Retrieved sold resource metric %v from repo entity %v is not the same as input %v",
			cap, repoEntity, cpuCap)
	}

	used, _ = repoEntity.soldMetricMap.GetMetricValue(resourceId1, property.USED)
	if used != cpuUsed {
		t.Errorf("Retrieved sold resource metric %v from repo entity %v is not the same as input %v",
			used, repoEntity, cpuUsed)
	}

	repoEntity.SetSoldResourceValue(resourceId2, property.CAP, memCap)
	repoEntity.SetSoldResourceValue(resourceId2, property.USED, memUsed)

	cap, _ = repoEntity.soldMetricMap.GetMetricValue(resourceId2, property.CAP)
	if cap != memCap {
		t.Errorf("Retrieved sold resource metric %v from repo entity %v is not the same as input %v",
			cap, repoEntity, memCap)
	}

	used, _ = repoEntity.soldMetricMap.GetMetricValue(resourceId2, property.USED)
	if used != memUsed {
		t.Errorf("Retrieved sold resource metric %v from repo entity %v is not the same as input %v",
			used, repoEntity, memUsed)
	}
}
