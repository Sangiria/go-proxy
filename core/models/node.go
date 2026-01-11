package models

import "encoding/json"

type Transport string
const (
	TransportTCP  Transport = "tcp"
	TransportXHTTP Transport = "xhttp"
)

type Security string
const (
	SecurityNone   Security = ""
	SecurityReality Security = "reality"
)

type SourceType string
const (
	SourceManual = "manual"
	SourceSubscription = "subscription"
)

type Node struct {
	ID			string		`json:"id"`
	Name		string		`json:"name"`
	Source		Source		`json:"source"`
	URL			string		`json:"url"`
	Parsed		Parsed		`json:"parsed"`
}

type Source struct {
	Type			SourceType	`json:"type"`
	SubscriptionID	string		`json:"subscription_id,omitempty"`
}

type Parsed struct {
	Address			string				`json:"address"`
	Port			uint16				`json:"port"`
	UUID			string				`json:"uuid"`
	Transport		Transport			`json:"transport"`
	Security		Security			`json:"security,omitempty"`
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