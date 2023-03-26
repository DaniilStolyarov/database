package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var A string

type Database struct {
	connectionString string
	Sql              *sql.DB
	UserTable        *UsersTable
}

func MakeDatabase() *Database {
	var database *Database = &Database{}
	database.connectionString = "postgres://go_project:rIo3Fc@95.140.159.168:5433/lib_service?sslmode=disable"
	database.Sql = nil
	database.UserTable = &UsersTable{_database: database}
	return database
}

// Start database if not started
func (db *Database) Start() {
	if db.Sql != nil {
		fmt.Printf("ERROR::DATABASE::START: can`t start database, it`s already connected")
		return
	}
	sql_db, err := sql.Open("postgres", db.connectionString)
	db.Sql = sql_db
	if err != nil {
		panic(err)
	}
}

// Close connection with database
func (db *Database) Close() {
	defer db.Sql.Close()
}

type dbTable struct {
	schemaName  sql.NullString
	tableName   sql.NullString
	tableOwner  sql.NullString
	tableSpace  sql.NullString
	hasIndexes  sql.NullBool
	hasRules    sql.NullBool
	hasTriggers sql.NullBool
	rowSecurity sql.NullBool
}

func (db *Database) GetTablesList() []dbTable {
	rows, err := db.Sql.Query("SELECT * FROM pg_catalog.pg_tables where (tablename not like 'pg_%') and (tablename not like 'sql_%')")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var tables = []dbTable{}

	for rows.Next() {
		var table dbTable
		err := rows.Scan(&table.schemaName, &table.tableName, &table.tableOwner, &table.tableSpace,
			&table.hasIndexes, &table.hasRules, &table.hasTriggers,
			&table.rowSecurity)
		if err != nil {
			panic(err)
		}

		tables = append(tables, table)
	}

	return tables
}

// prints names of all databases by space
func (db *Database) ShowTablesList() {
	tables := db.GetTablesList()
	for i := range tables {
		fmt.Println(tables[i].tableName.String)
	}
}
