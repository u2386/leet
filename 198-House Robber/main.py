class Solution:
    def rob(self, nums):
        dp = [[0] * 2 for _ in range(len(nums))]
        dp[0][0] = 0
        dp[0][1] = nums[0]

        for i in range(1, len(nums)):
            dp[i][0] = max(dp[i-1][0], dp[i-1][1])
            dp[i][1] = dp[i-1][0] + nums[i]
            print(dp)
        return max(dp[-1][0], dp[-1][1])


def main():
    s = Solution()
    print(s.rob([1,2,3,1]))
    print(s.rob([2,7,9,3,1]))
    print(s.rob([2,1,1,2]))

main()