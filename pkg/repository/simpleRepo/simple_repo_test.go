package simpleRepo

import (
	"testing"
	"github.com/turbonomic/turbo-go-monitoring/pkg/model"
	"github.com/turbonomic/turbo-go-monitoring/pkg/repository"
)

func TestSimpleMetricRepo(t *testing.T) {
	//
	// Construct a list of repo entities based on the test data
	//
	repoEntities := []repository.RepositoryEntity{}
	for _, testEntity := range TestEntities {
		repoEntity := NewSimpleMetricRepoEntity(testEntity.entityType, testEntity.entityId)
		repoEntities = append(repoEntities, repoEntity)
	}
	//
	// Construct a repo and add those repo entities to the repo
	//
	repo := NewSimpleMetricRepo()
	repo.SetEntities(repoEntities)
	//
	// Check GetEntity result
	//
	for _, testEntity := range TestEntities {
		repoEntity, err := repo.GetEntity(testEntity.entityType, testEntity.entityId)
		if err != nil {
			t.Errorf("No repo entity for type %v and id %v found in repo %v: %s", testEntity.entityType, testEntity.entityId, repo, err)
		} else if repoEntity.GetType() != testEntity.entityType {
			t.Errorf("Retrieved type %v from repo %v for entity type %v and id %v is not the same as entered %v",
				repoEntity.GetType(), repo, testEntity.entityType, testEntity.entityId, testEntity.entityType)
		} else if repoEntity.GetId() != testEntity.entityId {
			t.Errorf("Retrieved id %v from repo %v for entity type %v and id %v is not the same as entered %v",
				repoEntity.GetId(), repo, testEntity.entityType, testEntity.entityId, testEntity.entityId)
		}
	}
	//
	// Check GetEntityInstances result
	//
	fakeEntityType := model.EntityType("fakeEntityType")
	repoEntities = repo.GetEntitiesByType(fakeEntityType)
	if repoEntities == nil {
		t.Errorf("GetEntityInstances() for type %v on repo %v should not return nil; " +
			"an empty list is expected instead", fakeEntityType, repo)
	}
	for _, testEntity := range TestEntities {
		repoEntities = repo.GetEntitiesByType(testEntity.entityType)
		if repoEntities == nil || len(repoEntities) == 0 {
			t.Errorf("GetEntityInstances() for type %v on repo %v should not be nil or empty.", testEntity.entityType, repo)
		}
		for _, repoEntity := range repoEntities {
			if repoEntity == nil {
				t.Errorf("GetEntityInstances() for type %v on repo %v should not return nil instances: %v",
					testEntity.entityType, repo, repoEntities)
			}
		}
	}

}
