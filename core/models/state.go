package models

type State struct {
	ActiveNodeId	string
	Subscriptions	[]*Subscription
	Nodes			[]*Node
}