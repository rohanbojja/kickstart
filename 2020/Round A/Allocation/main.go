package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	var cases, n, b int
	scanf("%d\n", &cases)

	for i := 0; i < cases; i++ {
		defer writer.Flush()
		scanf("%d %d\n", &n, &b)
		hc := make([]int, n)

		for j := 0; j < n; j++ {

			scanf("%d ", &hc[j])

		}

		sort.Ints(hc)

		ans := 0
		totalPrice := 0

		for j := 0; j < n; j++ {
			if totalPrice+hc[j] > b {
				break
			} else {
				ans++
				totalPrice += hc[j]
			}
		}
		printf("Case #%d: %d\n", i+1, ans)
	}
}
