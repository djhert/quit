# Quitters
A basic package that implements an easy to use `Quit()` function wrapper to organize these functions.

## Example
```go
package main

import (
	"fmt"
	"os"

	"github.com/hlfstr/signals/quitters"
)

func quit() {
	fmt.Println("Bye")
}

func joke() {
	fmt.Printf("exit status -1\nNot really, haha\n")
}

func shutdown() {
	fmt.Println("Ok really this time")
	os.Exit(1)
}

func main() {
	quitter.AddQuit(quit)
	quitter.AddQuit(joke)
	quitter.AddQuit(shutdown)
	defer quitter.Quit()
	fmt.Printf("I am alive\nOk, now I'm closing\n")
}
```

When `main()` ends, the `quit()` function will be run which will run all functions added.  

## Intended Use
The **quitters** package can be used by itself, however the intention for writing it was for its parent package **quit**.  Below is an expanded version of the example above to accept signals and user input to exit.

```go
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/hlfstr/quit"
	"github.com/hlfstr/quit/quitters"
)

func end() {
	fmt.Println("Bye")
}

func joke() {
	fmt.Printf("exit status -1\nNot really, haha\n")
}

func shutdown() {
	fmt.Println("Ok really this time")
	os.Exit(1)
}

func restart() {
	fmt.Println("I'm restarting!")
}

func main() {
	quitters.AddQuit(end)
	quitters.AddQuit(joke)
	quitters.AddQuit(quit.Quit)
	quitters.AddQuit(shutdown)
	defer quit.Quit()
	fmt.Printf("I am alive\n")
	quit.Run(quitters.Quit, restart)
	fmt.Printf("Press 'Enter' to close, or send signal\n")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
```

## Functions
___
```go 
func AddQuit(f func())
```
Add a function to the quit list. Functions accept nothing and return nothing
___

```go
func Quit()
```
Run the quit list.
___