package literal

const WelcomeMessage = `Type 'quit' to exit at any time. Select search options:
 	* Enter '1' to search.	
 	* Enter '2' to view a list of searchable fields.
`

const ChooseModelMessage = `Please choose model you want to search. Select options:
	 * Enter '1' for User.
	 * Enter '2' for Organization.
	 * Enter '3' for Ticket.
`

const InputFieldMessage = `Enter search term: `
const InputValueMessage = `Enter search value: `

const (
	OrganizationJSONFile    = "github.com/pawarchetan/zendesk-search/data/organizations.json"
	TicketJSONFile          = "github.com/pawarchetan/zendesk-search/data/tickets.json"
	UserJSONFile            = "github.com/pawarchetan/zendesk-search/data/users.json"

	//tables
	UserTableName = "user"
	OrganizationTableName = "organization"
	TicketTableName = "ticket"

	// Organization indexed attributes
	OrganizationIndexID = "Id"
	OrganizationIndexURL = "URL"
	OrganizationIndexExternalID = "ExternalID"
	OrganizationIndexName = "Name"
	OrganizationIndexDomainName = "DomainName"
	OrganizationIndexDetail = "Detail"
	OrganizationIndexSharedTicket = "SharedTicket"
	OrganizationIndexTags = "Tags"

	// Organization fields
	OrganizationIndexFieldID = "_id"
	OrganizationIndexFieldURL = "url"
	OrganizationIndexFieldExternalID = "external_id"
	OrganizationIndexFieldName = "name"
	OrganizationIndexFieldDomainName= "domain_names"
	OrganizationIndexFieldDetail = "details"
	OrganizationIndexFieldSharedTicket = "shared_tickets"
	OrganizationIndexFieldTags = "tags"




)