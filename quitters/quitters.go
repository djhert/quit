package quitters

var (
	quits []func()
)

func init() {
	quits = make([]func(), 0)
}

// AddQuit will add the function passed to the
// quits array
func AddQuit(f func()) {
	quits = append(quits, f)
}

// Quit runs all quits functions
func Quit() {
	for i := range quits {
		quits[i]()
	}
}
