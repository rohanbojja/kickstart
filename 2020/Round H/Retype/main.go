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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	defer writer.Flush()
	var cases int
	scanf("%d\n", &cases)

	for testCase := 0; testCase < cases; testCase++ {
		var n, k, s int
		scanf("%d %d %d\n", &n, &k, &s)
		ans := min(n+k, k+(k-s)+(n-s))
		printf("Case #%d: %d\n", testCase+1, ans)
	}
}
