package model

type Ticket struct {
	Id string `json:"_id"`
	URL string `json:"url"`
	ExternalID string `json:"external_id"`
	Type string `json:"type"`
	Subject string `json:"subject"`
	Description string `json:"description"`
	Priority string `json:"priority"`
	Status string `json:"status"`
	SubmitterId int `json:"submitter_id"`
	AssigneeId int `json:"assignee_id"`
	Tags []string `json:"tags"`
	HasIncidents bool `json:"has_incidents"`
	Via string `json:"via"`
	OrganizationID int `json:"organization_id"`
}

type TicketDto struct {
	Id string `json:"_id"`
	URL string `json:"url"`
	ExternalID string `json:"external_id"`
	Type string `json:"type"`
	Subject string `json:"subject"`
	Description string `json:"description"`
	Priority string `json:"priority"`
	Status string `json:"status"`
	SubmitterId int `json:"submitter_id"`
	AssigneeId int `json:"assignee_id"`
	Tags []string `json:"tags"`
	HasIncidents bool `json:"has_incidents"`
	Via string `json:"via"`
	OrganizationName string `json:"organization_name"`
}
