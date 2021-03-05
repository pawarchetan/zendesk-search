package model

type User struct {
	Id int `json:"_id"`
	URL string `json:"url"`
	ExternalID string `json:"external_id"`
	Name string `json:"name"`
	Alias string `json:"alias"`
	Active bool `json:"active"`
	Verified bool `json:"verified"`
	Shared bool `json:"shared"`
	Locale string `json:"locale"`
	Timezone string `json:"timezone"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Signature string `json:"signature"`
	Tags []string `json:"tags"`
	Suspended bool `json:"suspended"`
	Role string `json:"role"`
	OrganizationID int `json:"organization_id"`
}

type UserDto struct {
	Id int `json:"_id"`
	URL string `json:"url"`
	ExternalID string `json:"external_id"`
	Name string `json:"name"`
	Alias string `json:"alias"`
	Active bool `json:"active"`
	Verified bool `json:"verified"`
	Shared bool `json:"shared"`
	Locale string `json:"locale"`
	Timezone string `json:"timezone"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Signature string `json:"signature"`
	Tags []string `json:"tags"`
	Suspended bool `json:"suspended"`
	Role string `json:"role"`
	OrganizationName string `json:"organization_name"`
	SubmittedTickets []string `json:"submitted_tickets"`
	AssignedTickets []string `json:"assigned_tickets"`
}