package main

import (
	"database/sql"
	"fmt"
)

type usersRow struct {
	id            sql.NullString
	login         sql.NullString
	email         sql.NullString
	pass_hash     sql.NullString
	flags         sql.NullInt32
	confirm_token sql.NullString
}
type usersRowInsert struct {
	login         string
	email         string
	pass_hash     string
	flags         int
	confirm_token string
}
type UsersTable struct {
	_database *Database
}

func (users_table UsersTable) insert(row usersRowInsert) {
	_, err := users_table._database.Sql.Exec(`insert into users (login, email, pass_hash, flags, confirm_token) values 
	($1::varchar, $2::varchar, $3::varchar, $4::int, $5::varchar)`, row.login, row.email, row.pass_hash, row.flags, row.confirm_token)
	if err != nil {
		panic(err)
	}
}

// prints users table in console
func (users_table UsersTable) show() {
	rows, err := users_table._database.Sql.Query("select * from users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	fmt.Printf("id\tlogin\temail\tflags\n")
	isEmpty := true
	for rows.Next() {

		row := usersRow{}
		err := rows.Scan(&row.id, &row.login, &row.email, &row.pass_hash, &row.flags, &row.confirm_token)
		if err != nil {
			panic(err)
		}
		fmt.Printf(row.id.String + "\t" + row.login.String + "\t" + row.email.String + "\t" + fmt.Sprint(row.flags.Int32) + "\n")
		isEmpty = false
	}
	if isEmpty {
		fmt.Printf("\t<empty table>\t\n")
	}
}

// create users table
func (users_table UsersTable) create() {
	_, err := users_table._database.Sql.Exec(`create table IF NOT EXISTS users
	(	id serial PRIMARY KEY,
		login varchar,
		email varchar,
		pass_hash varchar,
		flags int,
		confirm_token varchar
	)`)
	if err != nil {
		panic(err)
	}
}

// truncate users table
func (users_table UsersTable) truncate() {
	_, err := users_table._database.Sql.Exec(`truncate table users`)
	if err != nil {
		panic(err)
	}
}

// drop users table
func (users_table UsersTable) drop() {
	_, err := users_table._database.Sql.Exec("drop table IF EXISTS users")
	if err != nil {
		panic(err)
	}
}
