package services

import (
	"github.com/pawarchetan/zendesk-search/internal/dao"
	"github.com/pawarchetan/zendesk-search/internal/model"
	"strconv"
)

type UserService struct {
	dao *dao.UserDao
	organizationService *OrganizationService
	ticketService *TicketService
}

func InitUserService(dao *dao.UserDao, organizationService *OrganizationService, ticketService *TicketService) *UserService {
	return &UserService{
		dao: dao,
		organizationService: organizationService,
		ticketService: ticketService,
	}
}

func (u *UserService) Import(filePath string) error {
	err := u.dao.Import(filePath)
	return err
}

func (u *UserService) Search(field string, value string) ([]*model.UserDto, error) {
	iterator, err := u.dao.Search(field, value)
	if err != nil {
		return nil, err
	}
	var users []*model.UserDto
	for obj := iterator.Next(); obj != nil; obj = iterator.Next() {
		p := obj.(*model.User)
		orgName, err := u.GetOrganizationName(p.OrganizationID)
		if err != nil {
			return nil, err
		}
		submittedTickets, err := u.GetTicketsByField("submitter_id", strconv.Itoa(p.Id))
		assignedTicket, err := u.GetTicketsByField("assignee_id", strconv.Itoa(p.Id))
		users = append(users, mapUserModelToDto(p, orgName, submittedTickets, assignedTicket))
	}
	return users, nil
}

func (u *UserService) GetOrganizationName(orgID int) (string, error) {
	organizationID := strconv.Itoa(orgID)
	organizations, err := u.organizationService.Search("_id", organizationID)
	if err != nil {
		return "", err
	}
	if len(organizations) != 0 {
		return organizations[0].Name, nil
	}
	return "", err
}

func (u *UserService) GetTicketsByField(field string, value string) ([]string, error) {
	tickets, err := u.ticketService.Search(field, value)
	if err != nil {
		return nil, err
	}
	var ticketsSubjects []string
	if len(tickets) != 0 {
		for _, ticket:= range tickets {
			ticketsSubjects = append(ticketsSubjects, ticket.Subject)
		}
		return ticketsSubjects, nil
	}
	return nil, nil
}

func mapUserModelToDto(user *model.User, organizationName string, submittedTickets []string,
	assignedTickets []string) *model.UserDto {
	return &model.UserDto{
		Id: user.Id,
		URL: user.URL,
		ExternalID: user.ExternalID,
		Alias: user.Alias,
		Name: user.Name,
		Verified: user.Verified,
		Shared: user.Shared,
		Locale : user.Locale,
		Timezone: user.Timezone,
		Email: user.Email,
		Phone: user.Phone,
		Signature: user.Signature,
		Tags: user.Tags,
		Suspended: user.Suspended,
		Role: user.Role,
		OrganizationName: organizationName,
		SubmittedTickets: submittedTickets,
		AssignedTickets: assignedTickets,
	}
}