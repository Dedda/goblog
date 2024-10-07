# Conditionals

Conditionals allow you to direct the flow of your program depending on the values of variables. So far all our functions
were just executing one line after another in the exact order we wrote them in our source file. That's nice but also a
little boring since nothing ever changes. As a simple test, let's try to alter the program to randomly greet only of two
possible people. For this we will need another package called `math/rand` for our random number that will determine
whether we want to greet Alice or Bob. We will also need a new language concept, the `if` conditional. It will allow us
check a condition and only execute a block of code if that condition is true. We will also define an alternative path if
the statement is false.

```go
package main

import "fmt"
import "math/rand"

func main() {
	var message string
	number := rand.Float32()
	if number < 0.5 {
		message = greeting("Alice")
	} else {
		message = greeting("Bob")
	}
	fmt.Println(message)
}

func greeting(name string) string {
	return "Hello, " + name + "!"
}
```

We start by creating a new `message` variable of type `string`. This will hold the message to the selected person.
`rand.Float32` is a function that will provide us with a floating point number between 0 (inclusive) and 1 (exclusive).
To get a roughly 50% chance of greeting either person, we will check if this `number` is smaller than `0.5`. If so, we
say hello to Alice and will not execute the part where we will greet Bob. This block is only run when `number` was
greater or equal to `0.5`. If you want to check multiple cases, you can combine `if` with as many `else if` branches as
you like. Let's extend the example to greet Alice, Bob and people we don't know with individual messages:

```go
package main

import "fmt"
import "math/rand"

func main() {
	var message string
	number := rand.Float32()
	if number < 0.3 {
		message = greeting("Alice")
	} else if number < 0.6 {
		message = greeting("Bob")
	} else {
		message = "Welcome, stranger!"
	}
	fmt.Println(message)
}

func greeting(name string) string {
	return "Hello, " + name + "!"
}
```

First we check if `number` is smaller than `0.3` which gives a 30% chance to greet Alice. We do the same for Bob but
notice how we check `number` against `0.6` instead of `0.3`? This is because every time, `number` is smaller than `0.3`,
the Alice branch is already selected and the `else` makes us skip the next comparisons. We thereby end up with a chance
of `0.6 - 0.3` which is `0.3` or 30% just like for Alice. For all remaining cases where `number` is at least `0.6`, we
want to greet a stranger whose name we don't know so we just say "Welcome, stranger!". These `if` statements can also
handle more complicated checks by combining the single conditions with the
following logical operators:

| Operator | Description                                                              |
|----------|--------------------------------------------------------------------------|
| `&&`     | Logical `AND`: Both arguments have to be `true`                          |
| `||`     | Logical `OR`: Either one or both arguments have to be `true`             |
| `!`      | Logical `NOT`: This negates the argument. We now expect it to be `false` |

In case you are wondering what types of comparisons you can make, here is a small list of them

| Operator | Description                                                                               |
|----------|-------------------------------------------------------------------------------------------|
| `==`     | `EQUALS`: Both arguments have to be equal                                                 |
| `!=`     | `NOT EQUALS`: Both arguments cannot be equal                                              |
| `\>`     | `GREATER THAN`: The left argument has to be greater than the right                        |
| `\<`     | `LESS THAN`: The left argument has to be less than the right                              |
| `\>=`    | `GREATER THAN OR EQUAL`: The left argument has to be greater than or equal to the right   |
| `\<=`    | `LESS THAN OR EQUAL`: The left argument has to be less than or equal to the right         |

Another type of conditional is the `switch` statement. It makes it easier to define different paths if you are checking
on the same value for each of them. Let me show you in a simple example:

```go
i := 2
switch i {
    case 1: fmt.Println("i is 1")
    case 2: fmt.Println("i is 2")
    default: fmt.Println("i is neihter 1 nor 2")
}


```

They can also be used with normal logical expressions like in `if` statements and even evaluate the types of variables
like this:

```go
switch i.(type) {
    case int: fmt.Println("i is of type int")
    case string: fmt.Println("i is of type string")
    default: fmt.Println("i is of unknown type")
}
```
