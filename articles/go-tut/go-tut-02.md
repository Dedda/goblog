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

Okay, now let us use this new knowledge in our previous example. We will start by creating a constant containing the
name of the greeted person and using it in our print statement. We will be splitting the old message into two and
rebuild it with the `+` operator.

```go
package main

import "fmt"

func main() {
	const name = "World"
	fmt.Println("Hello, " + name + "!")
}
```

Of course this is still quite useless at the moment as we just make everything more complicated. But as soon as we want
to work with the name a bit more, like using it multiple times or making different versions of it (like uppercase or
lowercase), this constant will come in handy. If you want to be able to change the value of the name in your code, you
can try out `var` instead of `const` or use the short hand version of the declaration by just writing `name := "World"`
and removing the `const` or `var`. This will automatically create a new variable. Hint: if the variable already exists
and you just want to change the value, use `=` instead of `:=` or your compiler will complain because that variable
already exists and cannot be declared a second time in the same scope (basically a block of code).

As a small exercise, try to declare `name` as a variable and use it multiple times to greet multiple people with
different names. You will find a possible solution at the beginning of the next lesson.
