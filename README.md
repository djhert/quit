# Quit
Package **quit** is a cross-platform implementation to listen for process calls to your application to run	specified functions.  This is intended to allow for a clean-up function when your application is sent a signal to quit, or a function to run when signaled to restart.

Listens for all POSIX signals that would result in the termination of the application, and SIGUSR2 for restart.

**NOTE:** Only _INTERRUPT_ and _KILL_ quit are available on Windows

The following signals are listend for, and what function they run.

**Quit:**
```
SIGHUP  (Signal Hang-Up)
SIGABRT (Signal Abort)
SIGIOT	(Signal Abort)
SIGTERM (Signal Terminate)
SIGINT	(Signal Interrupt)
SIGKILL (Signal Kill)
```
**_Intentionally_** the _SIGQUIT_ signal is not implemented as this is defined to also provide a core dump according to POSIX. Go conforms to this by default and is better behavior.  :smile:

**Restart:**
```
SIGUSR2 (Signal User 2)
```

## Usage
**Example:**
```go
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/hlfstr/quit"
)

func myQuit() {
	quit.Quit()
	// Exit
	os.Exit(1)
}

func myRestart() {
	// Do something to restart
	fmt.Println("Restart")
}

func main() {
	// Defer the quit
	defer quit.Quit()
	//Start the quit and add myQuit and myRestart
	quit.Run(myQuit, myRestart)
	fmt.Printf("Press 'Enter' to close, or send signal\n")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
```

In the above example, upon receiving any POSIX signal that would result in termination of the application the	function `myQuit()` will be triggered.
**_OR_**
Upon receiving the _SIGUSR2_ signal, the function `myRestart()` will trigger.

The function `myQuit()` is also called by defering `quit.Quit()`.  This is so your `myQuit()` function is called if the application is quit via a signal or normal lifecycle.

**NOTE:** The _SIGUSR2_ signal is likely not available for Windows natively

**HINT:**
```go
//You can pass the "myQuit" function instead of "myRestart" to bypass this behavior.
go quit.Signal(myQuit, myQuit)
```

## Important Note
**It is important that your "quit" function results in the termination of your application.**
If it does not, the standard signals that can close your application may no longer work thus leaving a process that is very difficult to kill.

If you are stuck with a process that you cannot kill, issue the SIGQUIT signal to the application.

On Unix-Based:
```sh
kill -s SIGQUIT $PID
```

On Windows:
```
You may be able to kill it in the Task Manager.  If not, restart your computer.
```

## Functions
___

```go
func Run(quit func(), restart func()) 
```
Starts listening for the signals and starts the goroutine.  Expects funtion to call on quit, and restart. 

Functions `quit()` and `restart()` are functions that do not accept anyting or return anything.
```
TODO: Look into implementing a return interface for quit() and restart().
```
___
```go
func Quit()
```
Closes the goroutine created by `quit.Quit()` function above.  Can be called to start or as part of the ending cycle
___