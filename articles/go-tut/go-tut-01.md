# Basic structure of a Go program

Let's take a closer look at the previous example of `Hello, World!`. As a reminder, this is what the code looked like:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

There are more lines of code than just the one which actually prints the message we see in our console. But what does
all this mean? Do we really need it all? Here is a short explanation of every line:

```go
package main
```

Each Go module (like your own project) consists of one or more packages. There are basically the different directories
you put your code into to keep it organized and group functions and data that belongs together and decide which of them
should be accessible from outside that package. You will learn more about packages later in this series as we will use
packages in our module as well.

```go
import "fmt"
```

This line tells the compiler that we will be using something from the `fmt` package of the Go standard library in this
file. When you use code from external libraries, these imports will usually start with the location of the library's
repository. Importing sub-packages of a module can be done with `/`. For example, the `http` package inside the `net`
package can be imported with `"net/http"`.

```go
func main() {
```

To tell the compiler where the program should start executing our code, we have to create a function called `main`.
The parenthesis after the function name normally contain the parameters. In this case we don't have any. We finish this
line with a `{` which starts the function "body" containing the code of the function. You will see those braces all
throughout your code as they are used to build structs to organize your data and they also define code blocks which are
used for example for conditionals and loops.

```go
fmt.Println("Hello, World!")
```

Finally, some something happens in our program! While all the previous lines were just information for the compiler (and
you), this line will actually result in something being done during execution. Here we print the text `Hello, World!` to
the console. More specifically, we call the function `Println` in the `fmt` package we previously imported and give it
the text we want it to print. Be aware that the quotation marks are important here since we are defining a `string`
value which for now you can just think of as text. If you omit the quotes, the compiler would first assume we want to
use the variables `Hello` and `World` which are not defined anywhere as arguments to the `Println` function and the `!`
would cause a syntax error since the compiler does not know what this is supposed to mean.

At the end of the function we just have to close its body with a `}` and we are done with our first program!

In the next lesson we will learn about variables and constants to make working with values a lot more comfortable in the
future.