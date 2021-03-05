package services

import (
	"github.com/pawarchetan/zendesk-search/internal/dao"
	"github.com/pawarchetan/zendesk-search/internal/model"
	"strconv"
)

type TicketService struct {
	dao *dao.TicketDao
	organizationService *OrganizationService
}

func InitTicketService(dao *dao.TicketDao, organizationService *OrganizationService) *TicketService {
	return &TicketService{
		dao: dao,
		organizationService: organizationService,
	}
}

func (t *TicketService) Import(filePath string) error {
	err := t.dao.Import(filePath)
	return err
}

func (t *TicketService) Search(field string, value string) ([]*model.TicketDto, error) {
	iterator, err := t.dao.Search(field, value)
	if err != nil {
		return nil, err
	}
	var tickets []*model.TicketDto
	for obj := iterator.Next(); obj != nil; obj = iterator.Next() {
		p := obj.(*model.Ticket)
		orgName, err := t.GetOrganizationName(p.OrganizationID)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, mapTicketModelToDto(p, orgName))
	}
	return tickets, nil
}

func (t *TicketService) GetOrganizationName(orgID int) (string, error) {
	organizationID := strconv.Itoa(orgID)
	organizations, err := t.organizationService.Search("_id", organizationID)
	if err != nil {
		return "", err
	}
	if len(organizations) != 0 {
		return organizations[0].Name, nil
	}
	return "", err
}

func mapTicketModelToDto(ticket *model.Ticket, orgName string) *model.TicketDto {
	return &model.TicketDto{
		Id: ticket.Id,
		URL: ticket.URL,
		ExternalID: ticket.ExternalID,
		Type: ticket.Type,
		Subject: ticket.Subject,
		Description: ticket.Description,
		Priority: ticket.Priority,
		Status: ticket.Status,
		SubmitterId: ticket.SubmitterId,
		AssigneeId: ticket.AssigneeId,
		Tags: ticket.Tags,
		HasIncidents: ticket.HasIncidents,
		Via: ticket.Via,
		OrganizationName: orgName,
	}
}
