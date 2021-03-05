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

type UserDao struct {
	db *db.InMemoryDB
}

func InitUserRepo(db *db.InMemoryDB) *UserDao {
	return &UserDao{
		db: db,
	}
}

func (u *UserDao) Import(filePath string) error {
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

	txn := u.db.Transaction()
	for decoder.More() {
		entity := &model.User{}
		err = decoder.Decode(entity)
		if err != nil {
			return err
		}
		if err := txn.Insert(literal.UserTableName, entity); err != nil {
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

func (u *UserDao) Search(field string, value string) (db.ResultIterator, error) {
	transaction := u.db.Transaction()
	finalValue, err := util.ConvertFieldValue(field, value, model.User{})
	iterator, err := transaction.Get(literal.UserTableName, field, finalValue)
	return iterator, err
}
