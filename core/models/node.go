package models

import (
	"encoding/json"
)

type SourceType string
const (
	SourceManual = "manual"
	SourceSubscription = "subcription"
)

type Parsed struct {
	Type			string				`json:"type"`
	Address			string				`json:"address"`
	Port			uint16				`json:"port"`
	UUID			string				`json:"uuid"`
	Transport		string				`json:"transport"`
	Security		string				`json:"security,omitempty"`
	Sni				string				`json:"sni,omitempty"`
	Fp				string				`json:"fp,omitempty"`
	Pbk				string				`json:"pbk,omitempty"`
	Sid 			string				`json:"sid,omitempty"`
	Flow			string				`json:"flow,omitempty"`
	Host			string				`json:"host,omitempty"`
	Path			string				`json:"path,omitempty"`
	XHTTPMode		string				`json:"mode,omitempty"`
	XHTTPExtra		json.RawMessage		`json:"extra,omitempty"`	
}

type Node struct {
	Name		string		`json:"name"`
	Source		Source		`json:"source"`
	Parsed		Parsed		`json:"parsed"`
}

type Source struct {
	Type			SourceType	`json:"type"`
	SubscriptionID	string		`json:"subscription_id,omitempty"`
}