package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

//return one-based index
func findSubString(s string, ss string, offset int) int {
	if strings.Index(s[offset:], ss) != -1 {
		return strings.Index(s[offset:], ss) + offset + 1
	}
	return -1
}

//Return zero-based indexes
func findAllSubStrings(s string, ss string) []int {
	var indexes []int
	i := 0
	for i != -1 {
		i = findSubString(s, ss, i)
		if i != -1 {
			indexes = append(indexes, i-1)
		}
	}
	return indexes
}

func main() {
	defer writer.Flush()
	var cases int
	scanf("%d\n", &cases)

	for i := 0; i < cases; i++ {

		var s string
		scanf("%s\n", &s)
		allKicks := findAllSubStrings(s, "KICK")
		allStarts := findAllSubStrings(s, "START")
		//printf("%v %v\n", allKicks, allStarts)
		ans := 0

		for _, kick := range allKicks {
			for _, start := range allStarts {
				if kick < start {
					ans++
				}
			}
		}
		printf("Case #%d: %d\n", i+1, ans)
	}
}
