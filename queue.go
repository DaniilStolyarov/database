package database

import (
	"database/sql"
	"fmt"
)

type queueRow struct {
	user_id    sql.NullInt32
	book_id    sql.NullInt32
	created_at sql.NullInt32
	status     sql.NullInt32
	id         sql.NullInt32
}

type queueRowInsert struct {
	user_id    sql.NullInt32
	book_id    sql.NullInt32
	created_at sql.NullInt32
	status     sql.NullInt32
}

type QueueTable struct {
	_database *Database
}

func (queue_table QueueTable) Insert(row queueRowInsert) {
	_, err := queue_table._database.Sql.Exec(`insert into queue (user_id, book_id, created_at, status) values 
	($1::int, $2::int, $3::int, $4::int)`,
		row.user_id, row.book_id, row.created_at, row.status)
	if err != nil {
		panic(err)
	}
}
func (queue_table QueueTable) Show(row queueRow) {
	rows, err := queue_table._database.Sql.Query("select * from queue")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	fmt.Printf("id\tuser_id\tbook_id\tcreated_at\tstatus")
	isEmpty := true
	for rows.Next() {
		row := queueRow{}
		err := rows.Scan(&row.id, &row.user_id, &row.book_id, &row.created_at, &row.status)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d\t%d\t%d\t%d\t%d", row.id.Int32, row.user_id.Int32,
			row.book_id.Int32, row.created_at.Int32, row.status.Int32)
		isEmpty = false
	}
	if isEmpty {
		fmt.Printf("\t<empty table>\t\n")
	}
}

// create queue table
func (queue_table QueueTable) Create() {
	_, err := queue_table._database.Sql.Exec(`create table IF NOT EXISTS queue
	(
		id serial PRIMARY KEY,
		user_id int,
		book_id int,
		created_at int,
		status int
	)`)
	if err != nil {
		panic(err)
	}
}

// truncate queue table
func (queue_table QueueTable) Truncate() {
	_, err := queue_table._database.Sql.Exec(`truncate table queue`)
	if err != nil {
		panic(err)
	}
}
