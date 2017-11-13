package repository

import (
	"github.com/turbonomic/turbo-go-monitoring/pkg/model"
)

// RepositoryEntity defines a set of interfaces to access a repository entity in the metric Repository.
// It has per-entity info such as id, type, and node ip, and interfaces to get and set metric values.
type RepositoryEntity interface {
	// GetId() returns the associated entity id
	GetId() model.EntityId
	GetDisplayName() string
	// GetType() returns the type of the associated entity
	GetType() model.EntityType

	// GetTypedId() returns the type and id of the associated entity
	GetTypedId() model.EntityTypedId

	GetSoldResource(resourceId *ResourceIdentifier) (*EntityResource, error)
	GetAllSoldResources() []*EntityResource

	SetSoldResource(soldResource *EntityResource)
	SetSoldResourceValue(resourceId *ResourceIdentifier, prop model.MetricPropType, value model.MetricValue)

	GetProviders() map[model.EntityType][]model.EntityId
	GetBoughtResource(providerType model.EntityType, provider model.EntityId, resourceId *ResourceIdentifier) (*EntityResource, error)
	GetBoughtResourcesByProvider(providerType model.EntityType, provider model.EntityId) ([]*EntityResource, error)
	GetBoughtResourcesByProviderType(providerType model.EntityType) (map[model.EntityId][]*EntityResource, error)

	SetBoughtResourceValue(providerType model.EntityType, provider model.EntityId, resourceId *ResourceIdentifier, prop model.MetricPropType, value model.MetricValue)
	SetBoughtResource(providerType model.EntityType, provider model.EntityId, boughtResource *EntityResource)



}


