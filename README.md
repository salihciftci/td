 # td 
 [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com) [![Build Status](https://travis-ci.org/salihciftci/td.svg?branch=master)](https://travis-ci.org/salihciftci/td) [![Go Report Card](https://goreportcard.com/badge/github.com/salihciftci/td)](https://goreportcard.com/report/github.com/salihciftci/td)

td is a tool for managing tasks. It's easy to use and not complicated. Written with Golang.


## Using td

Using td is easy, not complicated and quick.

### Add a task

To add a task, use `td [desc]`.

```
$ td Try to learn Golang.
$ td Go to school.
$ td Call Salih.
```

### Listing Tasks

For Listing your task just type `td`.

```
$ td
1: Try to learn Golang.
2: Go to school.
3: Call Salih.
```

### Completing a task

You can complete your task with `td -d index`.

```
$ td -d 2
$ td
1: Try to learn Golang.
2: Call Salih.
```

### Searching

You can even search a task with little trick.

```
$ td | grep Call
2: Call Salih.
```

### Completing all of your tasks.

Even you can complete all of your tasks. Just use `td -r`
```
$ td
1: Try to learn Golang.
2: Call Salih.
$ td -r
$ td

```

## Installing td

First we need a MySQL database. Edit the line [100](https://github.com/salihciftci/td/blob/master/td.go#L100) for your own database.

After that basically we need a table to store all of our tds.
Here is the MySQL query for creating table;
```sql
CREATE TABLE `td` (
  `id` int NOT NULL AUTO_INCREMENT,
  `tdId` int NOT NULL,
  `td` text NOT NULL,
  `owner` varchar(64) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `tdId` (`tdId`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;
```

And you are done! Build the td and use it!

**Note:** td has a `owner` variable in line [12](http://https://github.com/salihciftci/td/blob/master/td.go#L12) for multi-user usage.


## Contributing

td made in one night for me. I will improve td but feel free to contribute.

