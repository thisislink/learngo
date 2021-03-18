# Program.go Explained

**Code**
```
package main

import "fmt"

func main() {
	fmt.Println("I'm a hard coded sentence in main, the entry point for Go programs")
}
```

## What does it all mean?

```
package main
```
Every Go program is made up of packages and programs start running in package main.

```
import "fmt"
```
You can import a package into other packages with the import keyword. 

The "fmt" package is a standard package that comes with Go that allows you to print to the screen and accept user input.

```
func main() {
	fmt.Println("I'm a hard coded sentence in main, the entry point for Go programs")
}
```
If you're not writing a library, a main() function needs to be created to be the entry point to your application. 