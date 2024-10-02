# Variables and Constants

Like in any programming language (with a few exceptions), in Go you'll use variables and constants that hold the values
you're working with.

```go
const (
    one_half = 0.5
    four = 4
)

var (
    five = 5
)
```

Like the names already suggest, variables can be changed during execution of the program while constants are immutable.
Inside functions variables can also be defined like this:

```go
var a int // define a variable of type int. it will by default be initialized to 0
var b = 2
c := 3
```