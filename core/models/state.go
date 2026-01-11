package models

type State struct {
	ActiveNodeId	string			`json:"active_node"`
	Subscriptions	[]*Subscription	`json:"subscriptions"`
	Nodes			[]*Node			`json:"nodes"`
}