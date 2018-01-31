package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var tds []string

func add(arg []string) {
	read()
	name := ""
	for i, e := range arg {
		if i == 0 {
			name = e
		} else {
			name = name + " " + e
		}
	}
	tds = append(tds, name)
	write(tds)
}

func list() {
	read()
	for i, e := range tds {
		fmt.Printf("%d: %s\n", i+1, e)
	}
}

func reset() {
	tds = tds[:0]
	write(tds)
}

func write(tds []string) {
	file, err := os.Create(os.Getenv("HOME") + "/.tddb")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range tds {
		fmt.Fprintln(w, line)
	}
	w.Flush()
}

func done(which string) {
	read()
	if i, err := strconv.Atoi(which); err == nil {
		if i-1 < len(tds) {
			tds = append(tds[:i-1], tds[i:]...)
		}
	}
	write(tds)
}

func read() {
	tds = tds[:0]
	file, err := os.Open(os.Getenv("HOME") + "/.tddb")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tds = append(tds, scanner.Text())
	}
}

func help() {
	fmt.Println("td is a tool for managing tasks.")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("\ttd        \tlist all of your tasks.")
	fmt.Println("\ttd [desc] \tadd a new task.")
	fmt.Println("\tdone [arg]\tcomplete your task.")
	fmt.Println("\treset     \tcomplete all of your tasks.")
	fmt.Println("")
}

func main() {
	if _, err := os.Stat(os.Getenv("HOME") + "/.tddb"); os.IsNotExist(err) {
		write(tds)
	}

	arg := os.Args[1:]
	if len(arg) < 1 {
		list()
	} else {
		switch arg[0] {
		case "done":
			if len(arg) < 2 {
				help()
			} else {
				done(arg[1])
			}
		case "reset":
			reset()
		case "help":
			help()
		default:
			add(arg[0:])
		}
	}
}
