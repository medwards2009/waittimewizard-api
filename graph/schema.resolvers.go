package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.57

import (
	"context"

	"github.com/medwards2009/waittimewizard-api/fetchData"
	"github.com/medwards2009/waittimewizard-api/graph/model"
)

// Destinations is the resolver for the destinations field.
func (r *queryResolver) Destinations(ctx context.Context) ([]*model.Destination, error) {
	fetchedData, err := fetchData.FetchDestinations()
	if err != nil {
		return nil, err
	}

	// Create a new slice to hold the results
	var result []*model.Destination

	for _, dest := range fetchedData.Destinations {
		var parks []*model.Park
		for _, park := range dest.Parks {
			parks = append(parks, &model.Park{
				ID:   park.Id,
				Name: park.Name,
			})
		}

		result = append(result, &model.Destination{
			ID:    dest.Id,
			Name:  dest.Name,
			Slug:  dest.Slug,
			Parks: parks,
		})
	}

	return result, nil
}

// LiveData is the resolver for the liveData field.
func (r *queryResolver) LiveData(ctx context.Context, id string, entityType *model.EntityType) (*model.LiveData, error) {
	fetchedData, err := fetchData.FetchLiveData(id)
	if err != nil {
		return nil, err
	}

	var filteredLiveData []*model.LiveDataListItem
	if entityType == nil {
		filteredLiveData = fetchedData.LiveData
	} else {
		for _, entity := range fetchedData.LiveData {
			if entity.EntityType == *entityType {
				filteredLiveData = append(filteredLiveData, entity)
			}
		}
	}

	var result = &model.LiveData{
		ID:         fetchedData.ID,
		Name:       fetchedData.Name,
		EntityType: fetchedData.EntityType,
		Timezone:   fetchedData.Timezone,
		LiveData:   filteredLiveData,
	}

	return result, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
