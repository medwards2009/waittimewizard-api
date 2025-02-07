// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Destination struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Slug  string  `json:"slug"`
	Parks []*Park `json:"parks"`
}

type Forecast struct {
	Time       *string `json:"time,omitempty"`
	WaitTime   *int    `json:"waitTime,omitempty"`
	Percentage *int    `json:"percentage,omitempty"`
}

type LiveData struct {
	ID         string              `json:"id"`
	Name       string              `json:"name"`
	EntityType EntityType          `json:"entityType"`
	Timezone   string              `json:"timezone"`
	LiveData   []*LiveDataListItem `json:"liveData"`
}

type LiveDataListItem struct {
	ID             string      `json:"id"`
	Name           string      `json:"name"`
	EntityType     EntityType  `json:"entityType"`
	Status         StatusType  `json:"status"`
	LastUpdated    string      `json:"lastUpdated"`
	Queue          *Queue      `json:"queue,omitempty"`
	ShowTimes      []*Time     `json:"showTimes"`
	OperatingHours []*Time     `json:"operatingHours"`
	Forecast       []*Forecast `json:"forecast"`
}

type Park struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Query struct {
}

type Queue struct {
	Standby *WaitTime `json:"standby,omitempty"`
}

type Time struct {
	Type      string `json:"type"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type WaitTime struct {
	WaitTime *int `json:"waitTime,omitempty"`
}

type EntityType string

const (
	// Attraction
	EntityTypeAttraction EntityType = "ATTRACTION"
	// Theme park primary name
	EntityTypeDestination EntityType = "DESTINATION"
	// A sub park to to a destination ex. Hollywood Studios to Walt Disney World
	EntityTypePark EntityType = "PARK"
	// Show
	EntityTypeShow EntityType = "SHOW"
	// Restaurant
	EntityTypeRestaurant EntityType = "RESTAURANT"
	// Hotel
	EntityTypeHotel EntityType = "HOTEL"
)

var AllEntityType = []EntityType{
	EntityTypeAttraction,
	EntityTypeDestination,
	EntityTypePark,
	EntityTypeShow,
	EntityTypeRestaurant,
	EntityTypeHotel,
}

func (e EntityType) IsValid() bool {
	switch e {
	case EntityTypeAttraction, EntityTypeDestination, EntityTypePark, EntityTypeShow, EntityTypeRestaurant, EntityTypeHotel:
		return true
	}
	return false
}

func (e EntityType) String() string {
	return string(e)
}

func (e *EntityType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EntityType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EntityType", str)
	}
	return nil
}

func (e EntityType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type StatusType string

const (
	StatusTypeClosed        StatusType = "CLOSED"
	StatusTypeOperating     StatusType = "OPERATING"
	StatusTypeDown          StatusType = "DOWN"
	StatusTypeRefurbishment StatusType = "REFURBISHMENT"
)

var AllStatusType = []StatusType{
	StatusTypeClosed,
	StatusTypeOperating,
	StatusTypeDown,
	StatusTypeRefurbishment,
}

func (e StatusType) IsValid() bool {
	switch e {
	case StatusTypeClosed, StatusTypeOperating, StatusTypeDown, StatusTypeRefurbishment:
		return true
	}
	return false
}

func (e StatusType) String() string {
	return string(e)
}

func (e *StatusType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = StatusType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid StatusType", str)
	}
	return nil
}

func (e StatusType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
