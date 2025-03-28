package main

import "fmt"

func canJump(n int) int {
	nums := make([]int, n+1)
	nums[0] = 1
	nums[1] = 1
	for i := 2; i <= n; i++ {
		nums[i] = nums[i-1] + nums[i-2]
	}
	return nums[n]
}

func main() {
	fmt.Println("This is demo02")
	fmt.Println("New feature")
	fmt.Println("Hot bug fix2")
	fmt.Println("bug fix")
	fmt.Println("bug fix2 twice")
	fmt.Println("New feature2")
	fmt.Println("New feature of function canJump")
	n := 1
	fmt.Println(canJump(n))
}
