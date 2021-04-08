package processor1

import (
	"bufio"
	"fmt"
	"os"
)

var Writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Printf(f string, a ...interface{}) {
	fmt.Fprintf(Writer, f, a...)
}
