# td

`td` is a tool for managing tasks. It's easy to use and not complicated. Written with Golang.

## Using td

Using td is easy, not complicated and quick.

### Add a task

To add a task, use `td add [desc]` or `td a [desc]`.

```
$ td add Try to learn Golang.
$ td add Go to school.
$ td a Call Salih.
```

### Listing Tasks

For Listing your task just use `td list` or `td l`.

```
$ td list
1: Try to learn Golang.
2: Go to school.
3: Call Salih.
```

### Completing a task

You can complete your task with `td done index` or `td d index`.

```
$ td done 2
$ td l
1: Try to learn Golang.
2: Call Salih.
```

### Searching

You can even search a task with little trick.

```
$ td l | grep Call
2: Call Salih.
```

### Completing all of your tasks.

Even you can complete all of your tasks. Just use `td reset`
```
$ td l
1: Try to learn Golang.
2: Call Salih.
$ td reset
$ td l

```

## Installing td

td is written with Golang but there is no need a Golang installed to use td. You can [download](http://) the lastest version of td or just clone to repository. After that just use it.

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
