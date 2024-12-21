package dataModels

type Destinations struct {
	Destinations []*Destination `json:"destinations"`
}

type Park struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Destination struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Slug  string  `json:"slug"`
	Parks []*Park `json:"parks"`
}
