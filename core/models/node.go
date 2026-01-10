package models

import "encoding/json"

type Node struct {
	ID			string		`json:"id"`
	Name		string		`json:"name"`
	Source		Source		`json:"source"`
	URL			string		`json:"url"`
	NodeParsed	NodeParsed	`json:"parsed"`
}

type Source struct {
	Type			string	`json:"type"`
	SubscriptionID	string	`json:"subscription_id,omitempty"`
}

type NodeParsed struct {
	Address		string			`json:"address"`
	Port		uint			`json:"port"`
	UUID		string			`json:"uuid"`
	Transport	string			`json:"transport"`
	Security	string			`json:"security,omitempty"`
	Sni			string			`json:"sni,omitempty"`
	Fp			string			`json:"fp,omitempty"`
	Pbk			string			`json:"pbk,omitempty"`
	Sid 		string			`json:"sid,omitempty"`
	SpiderX		string			`json:"spiderX,omitempty"`
	Flow		string			`json:"flow,omitempty"`
	Host		string			`json:"host,omitempty"`
	Path		string			`json:"path,omitempty"`
	XHTTPMode	string			`json:"mode,omitempty"`
	XTTPExtra	json.RawMessage	`json:"extra,omitempty"`	
}