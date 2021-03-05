package services

import (
	"github.com/pawarchetan/zendesk-search/internal/dao"
	"github.com/pawarchetan/zendesk-search/internal/model"
)

type OrganizationService struct {
	dao *dao.OrganizationDao
}

func InitOrganizationService(dao *dao.OrganizationDao) *OrganizationService {
	return &OrganizationService{
		dao: dao,
	}
}

func (o *OrganizationService) Import(filePath string) error {
	err := o.dao.Import(filePath)
	return err
}

func (o *OrganizationService) Search(field string, value string) ([]*model.Organization, error) {
	iterator, err := o.dao.Search(field, value)
	if err != nil {
		return nil, err
	}
	var organizations []*model.Organization
	for obj := iterator.Next(); obj != nil; obj = iterator.Next() {
		p := obj.(*model.Organization)
		organizations = append(organizations, p)
	}
	return organizations, nil
}

