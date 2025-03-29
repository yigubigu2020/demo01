package main

import "fmt"

func canJump(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func main() {
	fmt.Println("This is demo02")
	fmt.Println("New feature")
	fmt.Println("Hot bug fix2")
	fmt.Println("bug fix")
	fmt.Println("bug fix2 twice")
	fmt.Println("New feature2")
	fmt.Println("New feature of function canJump")
}
