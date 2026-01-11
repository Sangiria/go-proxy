package models

type State struct {
	ActiveNodeId	string						`json:"active_node"`
	Subscriptions	map[string]*Subscription	`json:"subscriptions"`
	Nodes			map[string]*Node			`json:"nodes"`
}