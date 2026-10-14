# Projects

Small Go projects I build while learning the language. Each project lives in its
own folder and is a self-contained program.

## Add a new project

```bash
# from the repo root
mkdir projects/my-project
```

Create `projects/my-project/main.go`:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello from my-project!")
}
```

Run it:

```bash
go run ./projects/my-project
```

## Notes

- Everything shares the one module declared in the repo's root `go.mod`, so you
  don't need a separate `go mod init` for each project.
- A project can grow to multiple files — just add more `.go` files to its folder
  (e.g. `projects/my-project/storage.go`) and they become part of the same program.
- If a project ever needs its own third-party dependencies or versioning, give
  *that folder* its own `go.mod` later. You don't need to decide that up front.
