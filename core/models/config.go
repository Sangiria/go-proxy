package models

import "encoding/json"

type LogLevel struct {
  	LogLevel string `json:"loglevel"`
}

type Inbound struct {
	Tag       		string      	`json:"tag"`
  	Listen     	 	string      	`json:"listen"`
  	Port      		uint16      	`json:"port"`
  	Protocol    	string      	`json:"protocol"`
  	InboundSettings InboundSettings `json:"settings"`
	Sniffing		Sniffing		`json:"sniffing"`
}

type Sniffing struct {
	Enabled			bool		`json:"enabled"`
	DestOverride	[]string	`json:"destOverride"`
}

type InboundSettings struct {
  	Auth  string  	`json:"auth"`
  	Udp   bool  	`json:"udp"`
}

type Outbound struct {
  	Tag            string         		`json:"tag"`
  	Protocol       string         		`json:"protocol"`
  	Settings       OutboundSettings  	`json:"settings"`
  	StreamSettings StreamSettings 		`json:"streamSettings"`
}

type OutboundSettings struct {
  	VNext []VNext `json:"vnext"`
}

type VNext struct {
  	Address string		`json:"address"`
  	Port    uint16		`json:"port"`
  	Users   []VlessUser	`json:"users"`
}

type VlessUser struct {
  	ID         	string 	`json:"id"`
  	Encryption 	string 	`json:"encryption"`
  	Flow       	string 	`json:"flow,omitempty"` 	//optional
	Level		uint	`json:"level,omitempty"`	//optional
}

type StreamSettings struct {
  	Network         string          `json:"network"`					//for now tcp or xhttp only
  	Security        string          `json:"security,omitempty"`			//for now reality only
  	RealitySettings RealitySettings `json:"realitySettings,omitzero"`
	XttpSettings	XttpSettings	`json:"xhttpSettings,omitzero"` 	//if transport is xhttp (tcp is the default)
}

type XttpSettings struct {
	Host		string			`json:"host,omitempty"`		//optional
	Path		string			`json:"path,omitempty"`		//optional
	Mode		string			`json:"mode,omitempty"`		//optional
	XttpExtra	json.RawMessage	`json:"extra,omitempty"`	//optional	
}

type RealitySettings struct {
  	ServerName  string 	`json:"serverName"` 		//sni
  	Fingerprint string 	`json:"fingerprint"`		//fp
  	PublicKey   string 	`json:"publicKey"`			//pbk
  	ShortID     string 	`json:"shortId,omitempty"`	//sid optional
	SpiderX		string	`json:"spiderX,omitempty"`	//optional
}

type Config struct {
  	LogLevel    LogLevel	`json:"log"`
  	Inbounds   []Inbound	`json:"inbounds"`
  	Outbounds  []Outbound	`json:"outbounds"`
}