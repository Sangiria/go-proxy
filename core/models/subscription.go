package models

type Subscription struct {
	Name	string			`json:"name"`
	URL		string			`json:"url"`
	Nodes	map[string]Node	`json:"nodes"`
}