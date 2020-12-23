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
	defer writer.Flush()
	var cases int
	scanf("%d\n", &cases)

	for testCase := 0; testCase < cases; testCase++ {
		var n, d int
		scanf("%d %d\n", &n, &d)
		arr := make([]int, n+1)
		for i := 1; i <= n; i++ {
			scanf("%d ", &arr[i])
		}
		prev := d
		for i := n; i > 0; i-- {
			prev = prev - (prev % arr[i])
		}
		ans := prev
		printf("Case #%d: %d\n", testCase+1, ans)
	}
}
