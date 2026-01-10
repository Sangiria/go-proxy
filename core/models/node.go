package models

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
	Address		string	`json:"address"`
	Port		uint	`json:"port"`
	UUID		string	`json:"uuid"`
	Type		string	`json:"type"`
	Security	string	`json:"security"`
	Sni			string	`json:"sni"`
	Fp			string	`json:"fp"`
	Pbk			string	`json:"pbk"`
	Sid 		string	`json:"sid"`
}