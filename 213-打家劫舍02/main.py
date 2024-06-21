class Solution:
    def rob(self, nums):
        return max(self._rob(nums[:-1]), self._rob(nums[1:]))

    def _rob(self, nums):
        dp = [[0] * 2 for _ in range(len(nums))]
        dp[0][0] = 0
        dp[0][1] = nums[0]

        for i in range(1, len(nums)):
            dp[i][0] = max(dp[i-1][0], dp[i-1][1])
            dp[i][1] = dp[i-1][0] + nums[i]
        return max(dp[-1][0], dp[-1][1])


if __name__ == '__main__':
    s = Solution()
    print(s.rob([2,3,2]))
    print(s.rob([1,2,3,1]))
    print(s.rob([1,2,3]))
    print(s.rob([2,7,9,3,1]))
    print(s.rob([2,1,1,2]))
