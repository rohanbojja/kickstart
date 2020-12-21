package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

type person struct {
	position int
	times    int
}

func main() {
	defer writer.Flush()
	var cases int
	scanf("%d\n", &cases)

	for i := 0; i < cases; i++ {
		var n, x int
		scanf("%d %d\n", &n, &x)
		a := make([]person, n)

		for j := 0; j < n; j++ {
			scanf("%d ", &a[j].times)
			a[j].position = j + 1
			a[j].times = int(math.Ceil(float64(a[j].times) / float64(x)))
		}

		sort.Slice(a, func(i, j int) bool {
			if a[i].times == a[j].times {
				return a[i].position < a[j].position
			}
			return a[i].times < a[j].times
		})
		printf("Case #%d: ", i+1)

		for _, e := range a {
			printf("%d ", e.position)
		}
		printf("\n")
	}
}
