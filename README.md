# td

`td` is a tool for managing tasks. It's easy to use and not complicated. Written with Golang.

[![Build Status](https://travis-ci.org/salihciftci/td.svg?branch=master)](https://travis-ci.org/salihciftci/td)

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

You can complete your task with `td done index`.

```
$ td done 2
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

Even you can complete all of your tasks. Just use `td reset`
```
$ td
1: Try to learn Golang.
2: Call Salih.
$ td reset
$ td

```

## Installing td

td is written with Golang but there is no need a Golang installed to use td. You can [download](https://github.com/salihciftci/td/releases/download/v0.1.0/td) the lastest version of td or just clone to repository. After that just use it.

To using globally.
If you are using bash, run:
```
$ echo alias td="path/to/td" >> ~/.bash_profile
```

If you are using zsh, run:
```
$ echo alias td="path/to/td" >> ~/.zshrc
```

Don't you like td? Name it what ever you want.
```
alias t="path/to/td"
alias todo="path/to/td"
```

## Contributing

td made in one night for me. I will improve td but feel free to contribute.




