// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Destination struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Slug  string  `json:"slug"`
	Parks []*Park `json:"parks"`
}

type Park struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Query struct {
}
