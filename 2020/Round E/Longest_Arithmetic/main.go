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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var cases int
	scanf("%d\n", &cases)

	for i := 0; i < cases; i++ {
		defer writer.Flush()
		ans := 0
		var n int
		scanf("%d\n", &n)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			scanf("%d ", &a[i])
		}
		maxLength := 0
		curLength := 1
		prevDiff := a[0] - a[1]
		for i := 1; i < n; i++ {
			diff := a[i-1] - a[i]
			if prevDiff == diff {
				curLength++

			} else {
				curLength = 2
				prevDiff = diff
			}
			maxLength = max(maxLength, curLength)
		}
		ans = maxLength
		printf("Case #%d: %d\n", i+1, ans)
	}
}
