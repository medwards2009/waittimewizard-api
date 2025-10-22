package fetchData

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchDestinations(t *testing.T) {
	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Mock response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"destinations": [
				{
					"id": "test-dest-1",
					"name": "Test Destination",
					"slug": "test-destination",
					"parks": [
						{
							"id": "test-park-1",
							"name": "Test Park"
						}
					]
				}
			]
		}`))
	}))
	defer server.Close()

	// Use the mock server URL as our base URL
	destinations, err := FetchDestinationsWithBaseURL(server.URL)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if destinations == nil {
		t.Fatal("Expected destinations to not be nil")
	}

	if len(destinations.Destinations) != 1 {
		t.Errorf("Expected 1 destination, got %d", len(destinations.Destinations))
	}

	dest := destinations.Destinations[0]
	if dest.Id != "test-dest-1" {
		t.Errorf("Expected destination ID to be 'test-dest-1', got '%s'", dest.Id)
	}

	if len(dest.Parks) != 1 {
		t.Errorf("Expected 1 park, got %d", len(dest.Parks))
	}

	park := dest.Parks[0]
	if park.Id != "test-park-1" {
		t.Errorf("Expected park ID to be 'test-park-1', got '%s'", park.Id)
	}
}
