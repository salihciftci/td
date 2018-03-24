package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"strconv"
)

var owner = "td"

type td struct {
	Id    int    `json:"id"`
	TdId  int    `json:"tdId"`
	Td    string `json:"td"`
	Owner string `json:"owner"`
}

func list(db *sql.DB) {
	results, err := db.Query("Select tdId,td FROM td Where owner = ?", owner)
	if err != nil {
		log.Println(err.Error())
	}

	for results.Next() {
		var tds td
		err = results.Scan(&tds.TdId, &tds.Td)
		if err != nil {
			log.Println(err.Error())
		}

		strId := strconv.Itoa(tds.TdId)
		fmt.Println(strId + ": " + tds.Td)
	}

}

func add(db *sql.DB, args []string) {
	var tds td

	err := db.QueryRow("SELECT MAX(tdId) FROM td").Scan(&tds.TdId)
	if err != nil {
		//No error handling for getting 0(zero) value.
	}

	desc := ""
	for i, e := range args {
		if i == 0 {
			desc = e
		} else {
			desc = desc + " " + e
		}
	}

	strTdId := strconv.Itoa(tds.TdId + 1)
	insert, err := db.Query("Insert Into td(tdId,td,owner) VALUES (" + strTdId + ",'" + desc + "','" + owner + "')")

	if err != nil {
		log.Println(err.Error())
	}

	defer insert.Close()
}

func done(db *sql.DB, which string) {
	insert, err := db.Query("Delete from td Where tdId = " + which)

	if err != nil {
		log.Println(err.Error())
	}

	defer insert.Close()
}

func reset(db *sql.DB) {
	insert, err := db.Query("Delete FROM td")

	if err != nil {
		log.Println(err.Error())
	}

	defer insert.Close()
}

func help() {
	fmt.Println("td is a tool for managing tasks.")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("\ttd        \tlist all of your tasks.")
	fmt.Println("\ttd [desc] \tadd a new task.")
	fmt.Println("\t-d [arg]  \tcomplete your task.")
	fmt.Println("\t-r        \tcomplete all of your tasks.")
	fmt.Println("")
}

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/td")

	if err != nil {
		log.Println(err.Error())
	}

	defer db.Close()

	arg := os.Args[1:]
	if len(arg) < 1 {
		list(db)
	} else {
		switch arg[0] {
		case "-d":
			if len(arg) < 2 {
				help()
			} else {
				done(db, arg[1])
			}
		case "-r":
			reset(db)
		case "-h":
			help()
		default:
			add(db, arg[0:])
		}
	}
}
