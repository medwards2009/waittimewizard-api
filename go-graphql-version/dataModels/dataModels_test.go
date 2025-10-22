package dataModels

import (
	"testing"
)

func TestDestinationStructs(t *testing.T) {
	// Test creating a new park
	park := &Park{
		Id:   "test-park-1",
		Name: "Test Park",
	}

	if park.Id != "test-park-1" {
		t.Errorf("Expected park ID to be 'test-park-1', got '%s'", park.Id)
	}

	if park.Name != "Test Park" {
		t.Errorf("Expected park name to be 'Test Park', got '%s'", park.Name)
	}

	// Test creating a new destination with parks
	destination := &Destination{
		Id:    "test-dest-1",
		Name:  "Test Destination",
		Slug:  "test-destination",
		Parks: []*Park{park},
	}

	if len(destination.Parks) != 1 {
		t.Errorf("Expected destination to have 1 park, got %d", len(destination.Parks))
	}

	// Test creating destinations container
	destinations := &Destinations{
		Destinations: []*Destination{destination},
	}

	if len(destinations.Destinations) != 1 {
		t.Errorf("Expected destinations to have 1 destination, got %d", len(destinations.Destinations))
	}
}
