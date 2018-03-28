package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var (
	owner   = ""
	version = "0.1.1"
)

type td struct {
	ID    int    `json:"id"`
	TdID  int    `json:"tdID"`
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
		err = results.Scan(&tds.TdID, &tds.Td)
		if err != nil {
			log.Println(err.Error())
		}

		strID := strconv.Itoa(tds.TdID)
		fmt.Println(strID + ": " + tds.Td)
	}
}

func add(db *sql.DB, args []string) {
	var tds td

	err := db.QueryRow("SELECT MAX(tdId) FROM td WHERE owner = ?", owner).Scan(&tds.TdID)
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

	strTdID := strconv.Itoa(tds.TdID + 1)
	insert, err := db.Query("Insert Into td(tdId,td,owner) VALUES (" + strTdID + ",'" + desc + "','" + owner + "')")

	if err != nil {
		log.Println(err.Error())
	}

	defer insert.Close()
}

func done(db *sql.DB, which string) {
	_, err := strconv.Atoi(which)
	if err != nil {
		help()
	} else {
		insert, err := db.Query("Delete from td Where tdId = " + which + " AND owner = '" + owner + "'")

		if err != nil {
			log.Println(err.Error())
		}

		defer insert.Close()
	}

}

func reset(db *sql.DB) {
	insert, err := db.Query("Delete FROM td WHERE owner = ?", owner)

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
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	ip := os.Getenv("DB_IP")
	owner = os.Getenv("OWNER")

	db, err := sql.Open("mysql", user+":"+pass+"@tcp("+ip+")/lora")

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
