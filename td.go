package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	clip "github.com/atotto/clipboard"
	_ "github.com/go-sql-driver/mysql"
)

var (
	version = "0.1.2"
	tds     td
)

type td struct {
	ID    int    `json:"id"`
	TdID  int    `json:"tdID"`
	Td    string `json:"td"`
	Owner string `json:"owner"`
}

//list listing all to-dos from database.
func list(db *sql.DB, owner string) {
	results, err := db.Query("Select tdId,td FROM td Where owner = ?", owner)
	if err != nil {
		log.Println(err.Error())
	}

	for results.Next() {
		err = results.Scan(&tds.TdID, &tds.Td)
		if err != nil {
			log.Println(err.Error())
		}

		strID := strconv.Itoa(tds.TdID)
		fmt.Println(strID + ": " + tds.Td)
	}
}

//add new task with given argument to database.
func add(db *sql.DB, args []string, owner string) {

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
	fmt.Println("==> '" + desc + "' added.")
}

//done completes single td with given argument from database.
func done(db *sql.DB, which string, owner string) {
	_, err := strconv.Atoi(which)
	if err != nil {
		help()
	} else {
		err := db.QueryRow("Select 1 from td Where owner='" + owner + "' AND tdID = " + which).Scan(&tds.TdID)
		if err != nil {
			//Given index doesn't exist
			fmt.Println("==> td[" + which + "] doesn't exist.")
		} else {
			insert, err := db.Query("Delete from td Where tdId = " + which + " AND owner = '" + owner + "'")
			if err != nil {
				log.Println(err.Error())
			}
			defer insert.Close()
			fmt.Println("==> td[" + which + "] completed.")
		}
	}
}

//reset deletes all tds from database.
func reset(db *sql.DB, owner string) {
	insert, err := db.Query("Delete FROM td WHERE owner = ?", owner)

	if err != nil {
		log.Println(err.Error())
	}

	defer insert.Close()
	fmt.Println("==> Completed all of your tds.")
}

func clipboard(db *sql.DB, which string, owner string) {
	_, err := strconv.Atoi(which)
	if err != nil {
		help()
	} else {
		err := db.QueryRow("Select 1 FROM td Where owner='" + owner + "' AND tdID = " + which).Scan(&tds.TdID)
		if err != nil {
			fmt.Println("==> td[" + which + "] doesn't exist.")
		} else {
			err := db.QueryRow("Select td FROM td Where owner='" + owner + "' AND tdID = " + which).Scan(&tds.Td)
			if err != nil {
				log.Println(err.Error())
			}
			clip.WriteAll(tds.Td)
			fmt.Println("==> td[" + which + "] has been copied to clipboard.")
		}
	}

}

//help stdout usage of td.
func help() {
	fmt.Println("td is a tool for managing tasks.")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("\ttd        \tlist all of your tasks.")
	fmt.Println("\ttd [desc] \tadd a new task.")
	fmt.Println("\t-d [index]  \tcomplete your task.")
	fmt.Println("\t-c [index]  \tcopy to clipboard.")
	fmt.Println("\t-r        \tcomplete all of your tasks.")
	fmt.Println("")
}

func main() {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	ip := os.Getenv("DB_IP")
	owner := os.Getenv("OWNER")

	//Connecting to MySQL
	db, err := sql.Open("mysql", user+":"+pass+"@tcp("+ip+")/td")

	if err != nil {
		log.Println(err.Error())
	}

	defer db.Close()

	//Checking td table exist or not
	err = db.QueryRow("SELECT 1 FROM td LIMIT 1").Scan(&tds.ID)
	if err != nil {
		log.Println(err.Error())
	}

	//Parsing Args
	arg := os.Args[1:]
	if len(arg) < 1 {
		list(db, owner)
	} else {
		switch arg[0] {
		case "-d":
			if len(arg) < 2 {
				help()
			} else {
				done(db, arg[1], owner)
			}
		case "-c":
			if len(arg) < 2 {
				help()
			} else {
				clipboard(db, arg[1], owner)
			}
		case "-r":
			reset(db, owner)
		case "-h":
			help()
		default:
			add(db, arg[0:], owner)
		}
	}
}
