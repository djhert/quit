package quit

import "fmt"

const (
	//NAME of the library
	NAME = "quit"
	//MAJOR version of library
	MAJOR = "1"
	//MINOR version of library
	MINOR = "0"
	//RELEASE version of library
	RELEASE = "2"
)

//Info returns a string with the name of the library, and the version
func Info() string {
	return fmt.Sprintf("%s - v%s.%s.%s", NAME, MAJOR, MINOR, RELEASE)
}
