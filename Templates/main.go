package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	var cases int
	scanf("%d\n", &cases)

	for i := 0; i < cases; i++ {
		defer writer.Flush()
		ans := 0

		printf("Case #%d: %d\n", i+1, ans)
	}
}
