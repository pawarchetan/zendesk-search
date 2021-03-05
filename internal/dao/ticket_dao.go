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

type TicketDao struct {
	db *db.InMemoryDB
}

func InitTicketDao(db *db.InMemoryDB) *TicketDao {
	return &TicketDao{
		db: db,
	}
}

func (t *TicketDao) Import(filePath string) error {
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

	txn := t.db.Transaction()
	for decoder.More() {
		entity := &model.Ticket{}
		err = decoder.Decode(entity)
		if err != nil {
			return err
		}
		if err := txn.Insert(literal.TicketTableName, entity); err != nil {
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

func (t *TicketDao) Search(field string, value string) (db.ResultIterator, error) {
	transaction := t.db.Transaction()
	finalValue, err := util.ConvertFieldValue(field, value, model.Ticket{})
	iterator, err := transaction.Get(literal.TicketTableName, field, finalValue)
	return iterator, err
}

