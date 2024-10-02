# Variables and constants

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

Look at this example. Here the variable `v` is first set to 3 and in the next line changed to 5.
Keep in mind that variables cannot change their type.
This is not possible with the constant `c` since its value cannot be changed.

```go
var v = 3
v = 5
v = "six" // This will result in a compiler error

const c = 3
c = 5 // This will result in a compiler error
```