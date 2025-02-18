package models

type Lookup struct {
	Fqdn string `json:"fqdn"`
	Dns  string `json:"dns"`
	Type string `json:"type"`
}
