package database

import (
	"bytes"
	"fmt"
	"log"
	"strings"
)

func (db *database) Import(datas []byte) {

	lines := bytes.Split(datas, []byte{'\n'})

	if len(lines) <= 1 {
		log.Println("something wrong")
	}

	fields := db.importFieldsName(lines[0])

	db.importRows(lines[1:], fields)
}

func (db *database) importFieldsName(f []byte) []string {

	query, fields := createCreateTableQuery(string(f))

	err := db.DB.Exec(query).Error
	if err != nil {
		log.Panic(err)
	}

	return fields
}

func (db *database) importRows(rows [][]byte, fields []string) {

	for _, row := range rows {

		if len(row) == 0 {
			continue
		}

		splitted := strings.Split(string(row), "|")

		if len(splitted) != len(fields) {
			log.Println(splitted)
			log.Println(fields)
			log.Panic("Something wrong")

		}

		query := createInsertIntoQuery(fields, splitted)

		err := db.DB.Exec(query).Error
		if err != nil {
			log.Panic(err)
		}

	}
}

// TODO attention to special char (' % ...)
func createCreateTableQuery(fields string) (string, []string) {

	splitted := strings.Split(fields, "|")

	query := "CREATE TABLE clanndatabase ("

	for i, field := range splitted {
		field = strings.TrimSpace(field)
		field = strings.ReplaceAll(field, " ", "")
		if i != 0 {
			query += ","
		}
		query += fmt.Sprintf("%s varchar", field)
	}

	return query + ")", splitted
}

func createInsertIntoQuery(fields []string, row []string) string {

	query := "INSERT INTO clanndatabase ("

	for i, field := range fields {
		field = strings.TrimSpace(field)
		field = strings.ReplaceAll(field, " ", "")
		if i != 0 {
			query += ","
		}
		query += field
	}
	query += ") VALUES ("

	for i, r := range row {
		r = strings.TrimSpace(r)
		r = strings.ReplaceAll(r, " ", "")

		if i != 0 {
			query += ","
		}
		if r != "" {
			query += "'" + r + "'"
		} else {
			query += "''"
		}

	}

	return query + ")"
}
