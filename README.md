# go-aoc-template

This project aims to be a Go language Advent of Code template.

# Instructions

## Code

- For each day (1-25) you simply create a go file under `internal/day` with the number of the day
- For each exercise per day (a & b) a function needs to be created with the following signature: 

```go
func (d *D) DayXXY(input string) any { }
```

Where `XX` is the day number (2 digits, with a leading zero, like `01`) and `Y` is the day's exercise (either `a` or `b`). The `input string` is the exercise input text that will be injected to the function.

Note that the function is attached to the `D` struct. This is a trick to use reflection in Go, avoiding to manually call all functions under `day` package.

Example: [internal/day/day01.go](internal/day/day01.go)

## Input text

The text provided by Advent of Code website should be stored at `assets/input` folder with the name `XX.txt`, where `XX`is the day number (2 digits, with a leading zero, like `01`.

Example: [assets/input/01.txt](assets/input/01.txt)

## Do I need to care from anything else?

All functions + input text that comply with the rules described above will be automatically executed by the program.

## How-to run

The prerequisites are just having [Go](https://go.dev/doc/install) installed. Then you can run the following from the project root.

```sh
go run main.go
```
