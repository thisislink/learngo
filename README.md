# learngo
Repository to teach myself [golang](https://golang.org/ref/spec) on Windows.

# Table of Contents
- [Setup](#setup)
	- [Useful Tools and Commands](#useful-tools-and-commands)
	- [Your First Go Program](#your-first-go-program)
- [Create a Server](#create-a-server)
- [Create an HTML Page](#create-an-html-page)

## Setup
1. [Download installer](https://golang.org/dl/) and follow the prompts to install Go.
2. Verify Go is installed properly: [Official Windows Install instructions](https://golang.org/doc/install#windows)

### Useful Tools and Commands
#### Using go mod tidy
Fetch all dependencies 
```
go mod tidy
```
[Go mod tidy Documentation](https://golang.org/ref/mod#go-mod-tidy)
#### Install Godoc
Godoc generates documentation for your Go programs
```
go get golang.org/x/tools/cmd/godoc
```
[Godoc Documentation](https://pkg.go.dev/golang.org/x/tools/cmd/godoc)

#### Using Vet
Find errors compliers might miss
```
go vet
```
Run the above command when you've created a module.

[Vet Documentation](https://golang.org/cmd/vet)

#### Install Golint
Finds and prints out style mistakes
```
go get golang.org/x/lint/golint
```
Run the above command when you've created a module.

[Golint Documentation](https://github.com/golang/lint)

#### Install Hot Reloading
Auto reloading/refresh when making changes in development

```
go get -u github.com/cosmtrek/air
```
[Air Documentation](https://github.com/cosmtrek/air)

[Back to TOC](#table-of-contents) | [Back to Top](#learngo)

## Your First Go Program
Create a src folder and a .go file. You can name the .go file anything you want. I used main.go.

### Example Folder Structure
```
- src
-- Program.go
.gitignore
README.md
```
Copy the following into your .go file:

```
package main

import "fmt"

func main() {
	fmt.Println("I'm a hard coded sentence in main, the entry point for Go programs")
}
```
Enable dependency tracking to manage and track modules that have imported packages.
src is the name of the location where your module is located. For me, my main is located in the src folder.
```
go mod init src
```
Change directories into the src folder, if you haven't already, and run the code
```
go run .
```
You've now setup and executed your first Go application.

[Back to TOC](#table-of-contents) | [Back to Top](#learngo)

## Create a Server

[Back to TOC](#table-of-contents) | [Back to Top](#learngo)

## Create an HTML Page

[Back to TOC](#table-of-contents) | [Back to TOp](#learngo)