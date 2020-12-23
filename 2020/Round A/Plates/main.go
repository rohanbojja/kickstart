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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	defer writer.Flush()
	var cases int
	scanf("%d\n", &cases)

	for testCase := 0; testCase < cases; testCase++ {
		var stacks, stackLength, plates int

		scanf("%d %d %d\n", &stacks, &stackLength, &plates)

		sum := make([][]int, stacks+1)
		for i := range sum {
			sum[i] = make([]int, stackLength+1)
		}

		dp := make([][]int, stacks+1)
		for i := range dp {
			dp[i] = make([]int, plates+1)
		}

		for i := 1; i <= stacks; i++ {
			for j := 1; j <= stackLength; j++ {
				scanf("%d ", &sum[i][j])
				sum[i][j] += sum[i][j-1]
			}
		}

		for stackN := 1; stackN <= stacks; stackN++ {
			for plateN := 1; plateN <= plates; plateN++ {
				for cont := 0; cont <= min(stackLength, plateN); cont++ {
					//printf("%d %d %d\n",stackN, plateN, cont)
					dp[stackN][plateN] = max(dp[stackN][plateN], dp[stackN-1][plateN-cont]+sum[stackN][cont])
				}
			}
		}
		ans := dp[stacks][plates]
		/*
			Answer here
		*/
		printf("Case #%d: %d\n", testCase+1, ans)
	}
}
