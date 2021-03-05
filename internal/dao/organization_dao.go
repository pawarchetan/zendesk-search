package dao

import (
	"bufio"
	"encoding/json"
	"github.com/pawarchetan/zendesk-db/pkg/db"
	"github.com/pawarchetan/zendesk-search/internal/literal"
	"github.com/pawarchetan/zendesk-search/internal/model"
	"github.com/pawarchetan/zendesk-search/internal/util"
	"os"
)

type OrganizationDao struct {
	db *db.InMemoryDB
}

func InitOrganizationDao(db *db.InMemoryDB) *OrganizationDao {
	return &OrganizationDao{
		db: db,
	}
}

func (org *OrganizationDao) Import(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer util.CloseFile(file)

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)

	_, err = decoder.Token()
	if err != nil {
		return err
	}

	txn := org.db.Transaction()
	for decoder.More() {
		entity := &model.Organization{}
		err = decoder.Decode(entity)
		if err != nil {
			return err
		}
		if err := txn.Insert(literal.OrganizationTableName, entity); err != nil {
			return err
		}
	}

	_, err = decoder.Token()
	if err != nil {
		return err
	}

	txn.Commit()
	return nil
}

func (org *OrganizationDao) Search(field string, value string) (db.ResultIterator, error) {
	transaction := org.db.Transaction()
	finalValue, err := util.ConvertFieldValue(field, value, model.Organization{})
	iterator, err := transaction.Get(literal.OrganizationTableName, field, finalValue)
	return iterator, err
}
