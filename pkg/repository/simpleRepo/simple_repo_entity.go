package simpleRepo

import (
	"github.com/turbonomic/turbo-go-monitoring/pkg/model"
	"github.com/turbonomic/turbo-go-monitoring/pkg/repository"
	"fmt"
	"bytes"
)

// SimpleMetricRepoEntity is a simple implementation of the RepositoryEntity
type SimpleMetricRepoEntity struct {
	entityType model.EntityType
	entityId   model.EntityId
	displayName string
	// Resources sold by the entity
	soldMetricMap  repository.EntityResourceMap
	// Resources bought by the entity from different provider entities
	//boughtMetricMap  map[model.EntityId]repository.EntityResourceMap
	boughtMetricMapByType  map[model.EntityType]map[model.EntityId]repository.EntityResourceMap
}

func NewSimpleMetricRepoEntity(entityType model.EntityType, entityId model.EntityId) *SimpleMetricRepoEntity {
	return &SimpleMetricRepoEntity{
		entityId: entityId,
		entityType: entityType,
		soldMetricMap: make(repository.EntityResourceMap),
		//boughtMetricMap: make(map[string]repository.EntityResourceMap),
		boughtMetricMapByType: make(map[model.EntityType]map[model.EntityId]repository.EntityResourceMap),
	}
}


func NewSimpleMetricRepoEntityWithDisplayName(entityType model.EntityType, entityId model.EntityId, displayName string) *SimpleMetricRepoEntity {
	return &SimpleMetricRepoEntity{
		entityId: entityId,
		entityType: entityType,
		displayName: displayName,
		soldMetricMap: make(repository.EntityResourceMap),
		//boughtMetricMap: make(map[string]repository.EntityResourceMap),
		boughtMetricMapByType: make(map[model.EntityType]map[model.EntityId]repository.EntityResourceMap),
	}
}

func (repoEntity SimpleMetricRepoEntity) GetDisplayName() string {
	return repoEntity.displayName
}

func (repoEntity SimpleMetricRepoEntity) GetId() model.EntityId {
	return repoEntity.entityId
}

func (repoEntity SimpleMetricRepoEntity) GetType() model.EntityType {
	return repoEntity.entityType
}

func (repoEntity SimpleMetricRepoEntity) GetTypedId() model.EntityTypedId {
	return model.EntityTypedId{EntityType: repoEntity.entityType, EntityId: repoEntity.entityId}
}

func (repoEntity SimpleMetricRepoEntity) String() string {
	var buffer bytes.Buffer
	var line string
	line = fmt.Sprintf("%s::%s\n", repoEntity.GetType(), repoEntity.GetId())
	buffer.WriteString(line)

	line = fmt.Sprintf("Sold:\n")
	buffer.WriteString(line)
	line = fmt.Sprintf("%s\n", repoEntity.soldMetricMap.String())
	buffer.WriteString(line)
	for providerType, providerMap := range repoEntity.boughtMetricMapByType {
		for provider, resourceMap := range providerMap {
			line = fmt.Sprintf("Bought: %s::%s\n", providerType, provider)
			buffer.WriteString(line)
			line = fmt.Sprintf("%s\n", resourceMap)
			buffer.WriteString(line)
		}
	}
	return buffer.String()
}

func (repoEntity SimpleMetricRepoEntity) GetAllSoldResources() []*repository.EntityResource {
	var soldResources []*repository.EntityResource

	for _, resource := range repoEntity.soldMetricMap {
		soldResources = append(soldResources, resource)
	}
	return soldResources
}

func (repoEntity SimpleMetricRepoEntity) GetSoldResource(resourceId *repository.ResourceIdentifier) (*repository.EntityResource, error) {
	resource, exists := repoEntity.soldMetricMap[resourceId]
	if !exists {
		return nil, fmt.Errorf("missing resource %s", resourceId)
	}
	return resource, nil
}

func (repoEntity SimpleMetricRepoEntity) SetSoldResourceValue(resourceId *repository.ResourceIdentifier,
						prop model.MetricPropType, value model.MetricValue) {
	repoEntity.soldMetricMap.SetMetricValue(resourceId, prop, value)
}

func (repoEntity SimpleMetricRepoEntity) SetSoldResource(soldResource *repository.EntityResource) {
	repoEntity.soldMetricMap[soldResource.ResourceId] = soldResource
}

func (repoEntity SimpleMetricRepoEntity) SetBoughtResource(providerType model.EntityType,
								provider model.EntityId,
								boughtResource *repository.EntityResource) {
	// first find the provider type map
	_, exists := repoEntity.boughtMetricMapByType[providerType]
	if !exists {
		repoEntity.boughtMetricMapByType[providerType] = make(map[model.EntityId]repository.EntityResourceMap)
	}
	providerMap, _ := repoEntity.boughtMetricMapByType[providerType]
	// first find the provider type map
	_, exists = providerMap[provider]
	if !exists {
		providerMap[provider] = make(repository.EntityResourceMap)
	}
	resourceMap, _ := providerMap[provider]
	resourceMap[boughtResource.ResourceId] = boughtResource
	providerMap[provider] = resourceMap
}

func (repoEntity SimpleMetricRepoEntity) SetBoughtResourceValue(providerType model.EntityType,
								provider model.EntityId,
								resourceId *repository.ResourceIdentifier,
								prop model.MetricPropType, value model.MetricValue) {
	// first find the provider type map
	_, exists := repoEntity.boughtMetricMapByType[providerType]
	if !exists {
		repoEntity.boughtMetricMapByType[providerType] = make(map[model.EntityId]repository.EntityResourceMap)
	}
	// next find the map by provider id
	providerMap, _ := repoEntity.boughtMetricMapByType[providerType]
	_, exists = providerMap[provider]
	if !exists {
		providerMap[provider] = make(repository.EntityResourceMap)
	}
	// next the resource map for a particular provider id
	resourceMap, _ := providerMap[provider]
	resourceMap.SetMetricValue(resourceId, prop, value)
}

func (repoEntity SimpleMetricRepoEntity) GetProviders() map[model.EntityType][]model.EntityId {
	providerMap := make(map[model.EntityType][]model.EntityId)

	// first find the provider type map
	for providerType, providers := range repoEntity.boughtMetricMapByType {
		providerIds := []model.EntityId{}
		for providerId, _ := range providers {
			providerIds = append(providerIds, providerId)
		}
		providerMap[providerType] = providerIds
	}
	return providerMap
}

func (repoEntity SimpleMetricRepoEntity) GetBoughtResourcesByProviderType(providerType model.EntityType) (map[model.EntityId][]*repository.EntityResource, error) {
	// first find the provider type map
	providerMap, exists := repoEntity.boughtMetricMapByType[providerType]
	if !exists {
		return nil, fmt.Errorf("missing provider type %s\n", providerType)
	}

	var boughtResources map[model.EntityId][]*repository.EntityResource

	for provider, resourceMap := range providerMap {
		var resources []*repository.EntityResource
		for _, resource := range resourceMap {
			resources = append(resources, resource)
		}
		boughtResources[provider] = resources
	}
	return boughtResources, nil
}

func (repoEntity SimpleMetricRepoEntity) GetBoughtResource(providerType model.EntityType,
						provider model.EntityId,
						resourceId *repository.ResourceIdentifier) (*repository.EntityResource, error) {
	// first find the provider type map
	providerMap, exists := repoEntity.boughtMetricMapByType[providerType]
	if !exists {
		return nil, fmt.Errorf("missing provider type %s\n", providerType)
	}
	// next find the map by provider id
	resourceMap, exists := providerMap[provider]
	if !exists {
		return nil, fmt.Errorf("missing provider id %s\n", provider)
	}
	// next the resource map for a particular provider id
	resource, exists := resourceMap[resourceId]
	if !exists {
		return nil, fmt.Errorf("missing resource %s", resourceId)
	}
	return resource, nil

}

func (repoEntity SimpleMetricRepoEntity) GetBoughtResourcesByProvider(providerType model.EntityType,
								provider model.EntityId) ([]*repository.EntityResource, error) {
	// first find the provider type map
	providerMap, exists := repoEntity.boughtMetricMapByType[providerType]
	if !exists {
		return nil, fmt.Errorf("missing provider type %s\n", providerType)
	}
	// next find the map by provider id
	resourceMap, exists := providerMap[provider]
	if !exists {
		return nil, fmt.Errorf("missing provider id %s\n", provider)
	}

	var resourceList []*repository.EntityResource
	for _, resource := range resourceMap {
		resourceList = append(resourceList, resource)
	}
	return resourceList, nil
}


