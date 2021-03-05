package util

import (
	"github.com/pawarchetan/zendesk-db/pkg/db"
	"github.com/pawarchetan/zendesk-db/pkg/index"
	"github.com/pawarchetan/zendesk-search/internal/literal"
	"os"
	"reflect"
	"strconv"
)

func ConvertFieldValue(field string, value string, model interface{}) (interface{}, error) {
	val := reflect.ValueOf(model)
	for i := 0; i < val.Type().NumField(); i++ {
		if field == val.Type().Field(i).Tag.Get("json") {
			fieldType := val.Type().Field(i).Type
			finalValue, err := getFieldType(fieldType.Kind(), value)
			return finalValue, err
		}
	}
	return nil, nil
}

func getFieldType(kind reflect.Kind, value string) (interface{}, error) {
	switch kind {
	case reflect.Int:
		valueSearch, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		return valueSearch, nil
	case reflect.String:
		return value, nil
	case reflect.Bool:
		valueSearch, err := strconv.ParseBool(value)
		if err != nil {
			return nil, err
		}

		return valueSearch, nil
	case reflect.Slice:
		return value, nil
	}
	return -1, nil
}

func CloseFile(f *os.File) {
	f.Close()
}

func IndexTickets() map[string]*db.IndexSchema {
	return map[string]*db.IndexSchema{
		"_id": {
			Name:    "_id",
			Indexer: &index.StringFieldIndex{Field: "Id"},
		},
		"url": {
			Name:    "url",
			Indexer: &index.StringFieldIndex{Field: "URL"},
		},
		"external_id": {
			Name:    "external_id",
			Indexer: &index.StringFieldIndex{Field: "ExternalID"},
		},
		"type": {
			Name:    "type",
			Indexer: &index.StringFieldIndex{Field: "Type"},
		},
		"subject": {
			Name:    "subject",
			Indexer: &index.StringFieldIndex{Field: "Subject"},
		},
		"description": {
			Name:    "description",
			Indexer: &index.StringFieldIndex{Field: "Description"},
		},
		"priority": {
			Name:    "priority",
			Indexer: &index.StringFieldIndex{Field: "Priority"},
		},
		"status": {
			Name:    "status",
			Indexer: &index.StringFieldIndex{Field: "Status"},
		},
		"has_incidents": {
			Name:    "has_incidents",
			Indexer: &index.BoolFieldIndex{Field: "HasIncidents"},
		},
		"via": {
			Name:    "via",
			Indexer: &index.StringFieldIndex{Field: "Via"},
		},
		"tags": {
			Name:    "tags",
			Indexer: &index.StringSliceFieldIndex{Field: "Tags"},
		},
		"organization_id": {
			Name:    "organization_id",
			Indexer: &index.IntFieldIndex{Field: "OrganizationID"},
		},
		"submitter_id": {
			Name:    "submitter_id",
			Indexer: &index.IntFieldIndex{Field: "SubmitterId"},
		},
		"assignee_id": {
			Name:    "assignee_id",
			Indexer: &index.IntFieldIndex{Field: "AssigneeId"},
		},

	}
}

func IndexOrganizations() map[string]*db.IndexSchema {
	return map[string]*db.IndexSchema{
		literal.OrganizationIndexFieldID: {
			Name:    literal.OrganizationIndexFieldID,
			Indexer: &index.IntFieldIndex{Field: literal.OrganizationIndexID},
		},
		literal.OrganizationIndexFieldURL: {
			Name:    literal.OrganizationIndexFieldURL,
			Indexer: &index.StringFieldIndex{Field: literal.OrganizationIndexURL},
		},
		literal.OrganizationIndexFieldExternalID: {
			Name:    literal.OrganizationIndexFieldExternalID,
			Indexer: &index.StringFieldIndex{Field: literal.OrganizationIndexExternalID},
		},
		literal.OrganizationIndexFieldName: {
			Name:    literal.OrganizationIndexFieldName,
			Indexer: &index.StringFieldIndex{Field: literal.OrganizationIndexName},
		},
		literal.OrganizationIndexFieldDomainName: {
			Name:    literal.OrganizationIndexFieldDomainName,
			Indexer: &index.StringSliceFieldIndex{Field: literal.OrganizationIndexDomainName},
		},
		literal.OrganizationIndexFieldDetail: {
			Name:    literal.OrganizationIndexFieldDetail,
			Indexer: &index.StringFieldIndex{Field: literal.OrganizationIndexDetail},
		},
		literal.OrganizationIndexFieldSharedTicket: {
			Name:    literal.OrganizationIndexFieldSharedTicket,
			Indexer: &index.BoolFieldIndex{Field: literal.OrganizationIndexSharedTicket},
		},
		literal.OrganizationIndexFieldTags: {
			Name:    literal.OrganizationIndexFieldTags,
			Indexer: &index.StringSliceFieldIndex{Field: literal.OrganizationIndexTags},
		},
	}
}

func IndexUsers() map[string]*db.IndexSchema {
	return map[string]*db.IndexSchema{
		"_id": {
			Name:    "_id",
			Indexer: &index.IntFieldIndex{Field: "Id"},
		},
		"url": {
			Name:    "url",
			Indexer: &index.StringFieldIndex{Field: "URL"},
		},
		"external_id": {
			Name:    "external_id",
			Indexer: &index.StringFieldIndex{Field: "ExternalID"},
		},
		"name": {
			Name:    "name",
			Indexer: &index.StringFieldIndex{Field: "Name"},
		},
		"alias": {
			Name:    "alias",
			Indexer: &index.StringFieldIndex{Field: "Alias"},
		},
		"active": {
			Name:    "active",
			Indexer: &index.BoolFieldIndex{Field: "Active"},
		},
		"verified": {
			Name:    "verified",
			Indexer: &index.BoolFieldIndex{Field: "Verified"},
		},
		"shared": {
			Name:    "shared",
			Indexer: &index.BoolFieldIndex{Field: "Shared"},
		},
		"locale": {
			Name:    "locale",
			Indexer: &index.StringFieldIndex{Field: "Locale"},
		},
		"timezone": {
			Name:    "timezone",
			Indexer: &index.StringFieldIndex{Field: "Timezone"},
		},
		"email": {
			Name:    "email",
			Indexer: &index.StringFieldIndex{Field: "Email"},
		},
		"signature": {
			Name:    "signature",
			Indexer: &index.StringFieldIndex{Field: "Signature"},
		},
		"suspended": {
			Name:    "suspended",
			Indexer: &index.BoolFieldIndex{Field: "Suspended"},
		},
		"role": {
			Name:    "role",
			Indexer: &index.StringFieldIndex{Field: "Role"},
		},
		"tags": {
			Name:    "tags",
			Indexer: &index.StringSliceFieldIndex{Field: "Tags"},
		},
		"organization_id": {
			Name:    "organization_id",
			Indexer: &index.IntFieldIndex{Field: "OrganizationID"},
		},
	}
}
