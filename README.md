# Go Programming

My Go learning repository: language exercises as I learn, plus small projects I
build along the way. Everything lives in one Go module.

## Structure

```
.
├── go.mod                 # one module for the whole repo
├── learning/              # language exercises, one folder per topic
│   ├── basics/            # variables, loops, if/switch mixed practice
│   ├── variables/
│   ├── for-loop/
│   ├── if-else/
│   ├── pointers/
│   ├── pointer-challenge/ # Employee Salary Manager exercise
│   ├── slices/
│   ├── maps/
│   ├── interfaces/
│   ├── enums/
│   ├── error/
│   ├── strings-runes/
│   └── user-input/
└── projects/              # small projects (see projects/README.md)
```

Each folder is its own `package main` program.

## Running

Run any exercise or project by pointing `go run` at its folder:

```bash
go run ./learning/pointers
go run ./learning/pointer-challenge
go run ./projects/<name>
```

Build all packages to check they compile:

```bash
go build ./...
```

## Adding a new project

```bash
mkdir projects/my-project        # create the folder
# add projects/my-project/main.go with package main + func main()
go run ./projects/my-project     # run it
```

See [`projects/README.md`](projects/README.md) for the full pattern.

## Prerequisites

- [Go](https://golang.org/) installed (this repo uses Go 1.26).

## Resources

- [A Tour of Go](https://go.dev/tour/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)

---

Happy coding!
