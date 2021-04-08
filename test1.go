package processor1

import (
	"bufio"
	"fmt"
	"os"
)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Printf(f string, a ...interface{}) {
	fmt.Fprintf(writer, f, a...)
}
