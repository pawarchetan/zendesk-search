package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/pawarchetan/zendesk-db/pkg/db"
	"github.com/pawarchetan/zendesk-search/internal/dao"
	"github.com/pawarchetan/zendesk-search/internal/literal"
	"github.com/pawarchetan/zendesk-search/internal/services"
	"github.com/pawarchetan/zendesk-search/internal/util"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func main()  {
	dbSchema := createDBSchema()
	database, err := createDB(dbSchema)
	if err != nil {
		fmt.Println("Error: Error while creating db schema: ", err.Error())
	}

	organizationService := initOrganizationServicesDao(database)
	ticketService := initTicketServiceDao(database, organizationService)
	userService := initUserServicesDao(database, organizationService, ticketService)
	err = userService.Import(literal.UserJSONFile)
	if err != nil {
		fmt.Println("Error: Error while importing users: ", err.Error())
	}
	err = organizationService.Import(literal.OrganizationJSONFile)
	if err != nil {
		fmt.Println("Error: Error while importing organizations: ", err.Error())
	}
	err = ticketService.Import(literal.TicketJSONFile)
	if err != nil {
		fmt.Println("Error: Error while importing tickets: ", err.Error())
	}
	fmt.Println("1. Press 1 to search zendesk")
	fmt.Println("2. Press 2 to view list of searchable fields (indexed fields)")
	fmt.Println("3. Type `Q` to exit")

	for {
		fmt.Print(literal.WelcomeMessage)
		cmd, quit := readInput()
		if quit {
			break
		}

		if quit = commandHandler(database, cmd, userService, organizationService, ticketService); quit {
			break
		}
	}
}

func initUserServicesDao(db *db.InMemoryDB, organizationService *services.OrganizationService, ticketService *services.TicketService) *services.UserService {
	userDao := dao.InitUserRepo(db)
	userService := services.InitUserService(userDao, organizationService, ticketService)
	return userService
}

func initOrganizationServicesDao(db *db.InMemoryDB) *services.OrganizationService {
	organizationDao := dao.InitOrganizationDao(db)
	organizationService := services.InitOrganizationService(organizationDao)
	return organizationService
}

func initTicketServiceDao(db *db.InMemoryDB, organizationService *services.OrganizationService) *services.TicketService {
	ticketDao := dao.InitTicketDao(db)
	ticketService := services.InitTicketService(ticketDao, organizationService)
	return ticketService
}

func displaySearchableFields(db *db.InMemoryDB) {
	inmemory := db.TableSchema()
	tableSchema := inmemory.Tables
	for name, tableIndex := range tableSchema {
		fmt.Println("-----------------------------------------")
		fmt.Println("Search ", name, " with")
		for indexName:= range tableIndex.Indexes {
			fmt.Println(indexName)
		}
		fmt.Println()
	}
}

func commandHandler(db *db.InMemoryDB, cmd string, service *services.UserService, organizationService *services.OrganizationService,
	ticketService *services.TicketService) bool {
	switch cmd {
	case "1":
		return searchHandler(service, organizationService, ticketService)
	case "2":
		displaySearchableFields(db)
		return false
	}

	return false
}

func searchHandler(userService *services.UserService, organizationService *services.OrganizationService,
	ticketService *services.TicketService) bool {
	var (
		searchModel string
		searchField string
		searchValue string
		quit        bool
	)

	fmt.Print(literal.ChooseModelMessage)
	for {
		line, quit := readInput()
		if quit {
			return true
		}

		if line == "1" || line == "2" || line == "3" {
			searchModel = line
			break
		}
		fmt.Print("Please enter '1', '2' or '3':")
	}

	fmt.Print(literal.InputFieldMessage)
	searchField, quit = readInput()
	if quit {
		return true
	}

	fmt.Print(literal.InputValueMessage)
	searchValue, quit = readInput()
	if quit {
		return true
	}
	var err error
	var result interface{}
	switch searchModel {
	case "1":
		result, err = userService.Search(searchField, searchValue)
	case "2":
		result, err = organizationService.Search(searchField, searchValue)
	case "3":
		result, err = ticketService.Search(searchField, searchValue)
	}
	if err != nil {
		fmt.Println("-------------------------------------------")
		fmt.Println("Error occurred: ", err)
		return false
	}

	fmt.Println("-------------------------------------------")
	res, _ :=json.MarshalIndent(&result, "", "\t")
	fmt.Println(string(res))
	return false
}

func createDB(dbSchema *db.InMemoryDBSchema) (*db.InMemoryDB, error) {
	return db.Init(dbSchema)
}

func createDBSchema() *db.InMemoryDBSchema {
	return &db.InMemoryDBSchema{
		Tables: createTableSchema(),
	}
}

func createTableSchema() map[string]*db.TableSchema {
	return map[string]*db.TableSchema{
		literal.UserTableName: {
			Name:    literal.UserTableName,
			Indexes: util.IndexUsers(),
		},
		literal.OrganizationTableName: {
			Name:    literal.OrganizationTableName,
			Indexes: util.IndexOrganizations(),
		},
		literal.TicketTableName: {
			Name:    literal.TicketTableName,
			Indexes: util.IndexTickets(),
		},
	}
}

func readInput() (string, bool) {
	scanner.Scan()
	line := scanner.Text()
	if line == "quit" {
		return "", true
	}

	return line, false
}