# Hello, World!

Here is an example of the classic hello world program which is usually one of the first things to try when learning a
new programming language. All it does is write "Hello, World!" to the standard output (normally the console).

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

There are a few steps we have to take before we get here and make our first executable program in Go.

### Installation

The first thing we will do is installing the Go compiler. We need the compiler to translate the code we write from text
to a binary format the computer understands. You can find the official instructions on how to install Go on your
computer [here](https://go.dev/doc/install).
This will vary between different operating systems. For Mac and Windows, you will find an installer that will guide you
through the process. For linux you can install Go manually by downloading the archive provided on the site and
extracting it to the correct location before updating your PATH environment variable. But i would recommend you to just
use the package manager that comes with your distro which will do all that for you and also lets you easily install
updates. After installation, you should be able to run the following command and get a similar output:

```shell
$ go version
go version go1.23.1 linux/amd64
```

### Creating your first project

Creating Go modules (basically what you would call a Go project) is very simple. Open your shell in a new, empty
directory and run the following command. This will create a `go.mod` file which holds all the important information of
your module like module name and dependencies. Here we're just using `unit_calculator` as the module name (spoiler
alert, we're going to build a program that helps you to convert between different units of measurement). Usually here
you put the location of your repository (in most cases from you GitHub) so others can easily download and use your
module in their projects.

```shell
$ go mod init unit_calculator
```

This should create a file similar to this:

```
module unit_calculator

go 1.23.1
```

You are now ready to create your first Go file. For that, just create a file called `main.go` in the project directory
and copy in the example code from above. In theory it doesn't matter what you call the file as long as its name ends in
`.go` but `main` is a convention in most languages and also the name of the entry point of the program (the first
function that is called when you run your executable) so i think it is quite fitting here.
Go ahead and build your program to get your first executable!

```shell
$ go build .
```

You can also compile and run your program with `go run .`

Right now it will just print `Hello, World!` to the console. That is cool but try adjusting the code so it greets only
you instead of the whole world. And maybe add a second line to the output where it says `Goodbye!`. You have now
successfully created your first "complete" go program!

In the next lesson, we will have a closer look at the example above and i will explain to you what every line does.
