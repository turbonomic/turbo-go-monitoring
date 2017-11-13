package repository

import (
	"github.com/turbonomic/turbo-go-monitoring/pkg/model"
	sdkbuilder "github.com/turbonomic/turbo-go-sdk/pkg/builder"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
	"github.com/turbonomic/turbo-go-monitoring/pkg/model/entity"
	"github.com/turbonomic/turbo-go-monitoring/pkg/model/resource"
	"github.com/turbonomic/turbo-go-monitoring/pkg/model/property"
)

var (
	sdkEntityTypes = map[model.EntityType]proto.EntityDTO_EntityType {
		entity.NODE: proto.EntityDTO_VIRTUAL_MACHINE,
	}

	sdkCommTypes = map[model.ResourceType]proto.CommodityDTO_CommodityType {
		resource.CPU: proto.CommodityDTO_CPU,
	}
)

type DTOBuilder struct {
	EntityType	model.EntityType
}

func NewDTOBuilder(entityType model.EntityType) *DTOBuilder {
	return  &DTOBuilder{EntityType: entityType,}
}

func (builder *DTOBuilder) buildDTO(repoEntity RepositoryEntity) *sdkbuilder.EntityDTOBuilder{
	// id.
	entityID := string(repoEntity.GetId())
	entityDTOBuilder := sdkbuilder.NewEntityDTOBuilder(sdkEntityTypes[repoEntity.GetType()], entityID)

	// display name.
	displayName := repoEntity.GetDisplayName()
	entityDTOBuilder.DisplayName(displayName)

	// commodities sold.
	resourcesSold := repoEntity.GetAllSoldResources()
	var commoditiesSold []*proto.CommodityDTO
	for _, entityResource := range resourcesSold {

		commSold := sdkbuilder.NewCommodityDTOBuilder(sdkCommTypes[entityResource.ResourceId.ResourceType])
		commSold.Key(entityResource.ResourceId.Key)
		capProp, err := entityResource.GetMetricValue(property.CAP)
		if err != nil {
			commSold.Capacity(float64(capProp))
		}
		usedProp, err := entityResource.GetMetricValue(property.USED)
		if err != nil {
			commSold.Used(float64(usedProp))
		}
	}
	entityDTOBuilder.SellsCommodities(commoditiesSold)

	// commodities bought
	providers := repoEntity.GetProviders()
	for providerType, _ := range providers {
		providerMap, _ := repoEntity.GetBoughtResourcesByProviderType(providerType)
		for providerID, resourceMap := range providerMap {
			commBoughtList := []*proto.CommodityDTO{}
			for _, resource := range resourceMap {
				commBought := sdkbuilder.NewCommodityDTOBuilder(sdkCommTypes[resource.ResourceId.ResourceType])
				commBought.Key(resource.ResourceId.Key)
				usedProp, err := resource.GetMetricValue(property.USED)
				if err != nil {
					commBought.Used(float64(usedProp))
				}
				resProp, err := resource.GetMetricValue(property.RESERVATION)
				if err != nil {
					commBought.Reservation(float64(resProp))
				}
				cb, err := commBought.Create()
				if err != nil {
					commBoughtList = append(commBoughtList, cb)
				}
			}
			providerDto := sdkbuilder.CreateProvider(sdkEntityTypes[providerType], string(providerID))
			entityDTOBuilder.Provider(providerDto)
			entityDTOBuilder.BuysCommodities(commBoughtList)
		}
	}

	return entityDTOBuilder
}
