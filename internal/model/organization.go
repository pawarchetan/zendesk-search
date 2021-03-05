package model

type Organization struct {
	Id           int      `json:"_id"`
	URL          string   `json:"url"`
	ExternalID   string   `json:"external_id"`
	Name         string   `json:"name"`
	DomainName   []string `json:"domain_names"`
	Detail       string   `json:"details"`
	SharedTicket bool     `json:"shared_tickets"`
	Tags         []string `json:"tags"`
}