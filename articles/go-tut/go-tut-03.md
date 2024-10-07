# Functions

Before we get started with this lesson, here is an example of what your code might look like after completing the
exercise in the last lesson:

```go
package main

import "fmt"

func main() {
	name := "Alice"
	fmt.Println("Hello, " + name + "!")
	name = "Bob"
	fmt.Println("Hello, " + name + "!")
	name = "Carol"
	fmt.Println("Hello, " + name + "!")
}
```

Okay, this is cool. But still it looks like a lot of work to basically print the same text three times just with
changing names. There has to be an easier way to do this. And there is! Let me introduce functions to you. Just like
`main` in this example is a functions, you can also define your own functions the same way. Let us create a function
called `greet` to make greeting people easier and only write the printing code once. Add this to your `main.go` file:

```go
func greet(name string) {
    fmt.Println("Hello, " + name + "!")
}
```

This defines a function that takes in one parameter called `name` or type `string`. Remember `string`? This basically
means the kind of values we can give to this functions have to be text and not something else like for example numbers
(note that `"5"` is of type `string` while `5` without the quotes is treated as a number). This parameter `name` can be
used like a normal variable inside the function. You could change it or pass it on to other functions as well as adding
it to our nice greeting message. The last step would be to clean up our `main` function by replacing the calls to
`fmt.Println` with calls to our new `greet` function. We can also remove the `name` variable in `main` and pass in the
names to `greet` directly. Your code should now look like this and greet all three of our new friends:

```go
package main

import "fmt"

func main() {
	greet("Alice")
	greet("Bob")
	greet("Carol")
}

func greet(name string) {
	fmt.Println("Hello, " + name + "!")
}
```

Nice, now we only have to write the part that generates the message once and use it multiple times in out code. But what
if we want to do other things with the greeting message than just print it? if we alter the code in the `greet`
function, this new behavior will affect all calls to that function which might not be what we want. For example, if in
some scenarios we want to write the greeting to a file before wen print it, it would be easier to generate the message
in a function and then decide later on what we want to do with it. For these cases, we can use `return` types of
functions. By declaring a `return` type on your function, you are telling the callee that they can expect the function
to give back a value of that type whenever it is called. For our example we will replace `greet` with `greeting` and
make it `return` the generated message of type `string` by adding the type name behind the function parameters. We can
then assign the returned value to a variable in `main` or directly pass it down to `fmt.Println` like this:

```go
package main

import "fmt"

func main() {
	greetAlice := greeting("Alice")
	fmt.Println(greetAlice)
	writeToLog(greetAlice)
	fmt.Println(greeting("Bob"))
	fmt.Println(greeting("Carol"))
}

func greeting(name string) string {
	return "Hello, " + name + "!"
}
```

In this example we pass the returned greeting for Bob and Carol directly into `fmt.Println` but we also write the
greeting for Alice to the log file by passing it into `writeToLog`. The call to `writeToLog` is just an example of how
to use the returned value for more than once, the function does not really exist but it is a possible scenario you might
encounter in other programs.

This is only the tip of the iceberg. Functions of real programs will usually be more complex and take multiple
parameters of different types and sometimes also `return` multiple values. Functions themselves can actually be
values or even parameters to other functions. But that is something we will see later on.

To summarize, unctions are essential building blocks of all programs and allow us to group, organize and orchestrate
much more complicated code bases. They are the first step at splitting your on long and hard to grasp source files into
more understandable and reusable parts. We can give useful names to functions that describe what they do to make it
easier for ourselves and others to read the code later on. Let's continue to the next lesson where we learn about
conditionals to alter the behavior of our program depending on its internal state.
