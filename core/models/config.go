package models

type LogLevel struct {
  	LogLevel string `json:"loglevel"`
}

type Inbound struct {
	Tag       		string      	`json:"tag"`
  	Listen     	 	string      	`json:"listen"`
  	Port      		uint16      	`json:"port"`
  	Protocol    	string      	`json:"protocol"`
  	InboundSettings InboundSettings `json:"settings"`
}

type InboundSettings struct {
  	Auth  string  	`json:"auth"`
  	Udp   bool  	`json:"udp"`
}

type Outbound struct {
  	Tag            string         `json:"tag"`
  	Protocol       string         `json:"protocol"`
  	Settings       VlessSettings  `json:"settings"`
  	StreamSettings StreamSettings `json:"streamSettings"`
}

type VlessSettings struct {
  	VNext []VNext `json:"vnext"`
}

type VNext struct {
  	Address string		`json:"address"`
  	Port    uint16		`json:"port"`
  	Users   []VlessUser	`json:"users"`
}

type VlessUser struct {
  	ID         string `json:"id"`
  	Encryption string `json:"encryption"`
  	Flow       string `json:"flow,omitempty"`
}

type StreamSettings struct {
  	Network         string          `json:"network"`
  	Security        string          `json:"security"`
  	RealitySettings RealitySettings `json:"realitySettings"`
}

type RealitySettings struct {
  	ServerName  string `json:"serverName"`
  	Fingerprint string `json:"fingerprint"`
  	PublicKey   string `json:"publicKey"`
  	ShortID     string `json:"shortId"`
}

type Config struct {
  	LogLevel    LogLevel	`json:"log"`
  	Inbounds   []Inbound	`json:"inbounds"`
  	Outbounds  []Outbound	`json:"outbounds"`
}