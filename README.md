## Tutorials 

- [x] [Tutorial: Get started with Go](https://go.dev/doc/tutorial/getting-started)
- [ ] [Tutorial: Create a Go module](https://go.dev/doc/tutorial/create-module)
  - [x] Create a module
  - [x] Call your code from another module
  - [x] Return and handle an error
  - [x] Return a random greeting
  - [x] Return greetings for multiple people
  - [ ] Add a test
  - [ ] Compile and install the application
- [ ] [Tutorial: Getting started with multi-module workspaces](https://go.dev/doc/tutorial/workspaces.html)
- [ ] [Tutorial: Accessing a relational database](https://go.dev/doc/tutorial/database-access)
- [ ] [Tutorial: Developing a RESTful API with Go and Gin](https://go.dev/doc/tutorial/web-service-gin)
- [ ] [Tutorial: Getting started with generics](https://go.dev/doc/tutorial/generics)
- [ ] [Tutorial: Getting started with fuzzing](https://go.dev/doc/tutorial/fuzz)
- [ ] [Tutorial: Getting started with govulncheck](https://go.dev/doc/tutorial/govulncheck)
- [ ] [Tutorial: Find and fix vulnerable dependencies with VS Code Go](https://go.dev/doc/tutorial/govulncheck-ide)
- [ ] [Tutorial: A Tour of Go](https://go.dev/tour/)

## Extras

- [x] [How to Write Go Code](https://go.dev/doc/code)

## Useful Links

- The [main documentation page](https://go.dev/doc/) outlines a lot of the
  learning journeys and tutorial paths
- See [Effective Go](https://go.dev/doc/effective_go.html) for tips on writing
  clear, idiomatic Go code. 
- Take [A Tour of Go](https://go.dev/tour/) to learn the language proper. 
- Visit [the documentation page](https://go.dev/doc/#articles) for a set of
  in-depth articles about the Go language and its libraries and tools. 
- See the [testing package documentation](https://go.dev/pkg/testing/) for more
  detail on testing.
- [Go Maps in Action](https://go.dev/blog/maps)
- For more, see [The blank
  identifier](https://go.dev/doc/effective_go.html#blank) in Effective Go. 

## Learnings

### Creating Modules

Before writing any code, if you are writing a module you need to create the
directory which will host your module.

Run the go mod init command, giving it your module path -- here, use
example.com/greetings. If you publish a module, this must be a path from which
your module can be downloaded by Go tools. That would be your code's repository. 

```go
go mod init example.com/name_of_module
```

### Creating Applications

You go through the same motions as [creating modules](#creating-modules), and
then declare a main package. In Go, code executed as an application must be in a main
package. 

```go
package main

import (
  // packages
)

func main() {
    // do something using packages
}
```

### Creating variables

In Go, the := operator is a shortcut for declaring and initializing a variable
in one line (Go uses the value on the right to determine the variable's type).
Taking the long way, you might have written this as:

```go
var message string
message = fmt.Sprintf("Hi, %v. Welcome!", name)
```

## Creating Functions

In Go, a function whose name starts with a capital letter can be called by a
function not in the same package. This is known in Go as an exported name.

### Creating Arrays

#### Slices

Go has something called `slices`. These are like arrays, except their sizes
change dynamically in line with their contents.

When declaring a slice, you omit its size in the brackets, like this: 

```go
[]string.
```

This tells Go that the size of the array underlying the slice can be dynamically
changed. 

### Creating Maps

In Go, you initialize a map with the following syntax:
`make(map[key-type]value-type)`

```go
make(map[string]int)
```

Values can be added to the map, using either map literals:

```go
// this also creates the map

commits := map[string]int{
    "rsc": 3711,
    "r":   2138,
    "gri": 1908,
    "adg": 912,
}
```

Or separate statements:

```go
m = make(map[string]int)
m["route"] = 66
```

### Loops

#### For Loops

```go
for _, name := range names {
  message, err := Hello(name)
  if err != nil {
    return nil, err
  }
  // In the map, associate the retrieved message with
  // the name.
  messages[name] = message
}
```

The `range` returns two values: the index of the current item in the
loop and a copy of the item's value. You don't need the index, so you use the Go
blank identifier (an underscore) to ignore it. 