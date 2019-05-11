package rob

import "fmt"

func rob(nums []int) int {
	length := len(nums)

	if length == 0 {
		return 0
	} else if length == 1 {
		return nums[0]
	} else if length == 2 {
		return max(nums[0], nums[1])
	}

	return max(doRob(nums[:length - 1]), doRob(nums[1:]))
}

func doRob(nums []int) int {
	length := len(nums)
	dp := make([]int, length)
	for i := 0; i < length; i++ {
		dp[i] = 0;
	}
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i, n := range(nums[2:]) {
		i += 2
		dp[i] = max(dp[i-1], n + dp[i-2])
	}
	return dp[len(dp) - 1]
}


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1,2,3,1}
	fmt.Println(rob(nums))
}
