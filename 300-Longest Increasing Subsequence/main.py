class Solution:
    def lengthOfLIS(self, nums: List[int]) -> int:

        def bsearch(arr, v):
            l, r = 0, len(arr)-1
            while l < r:
                m = (l+r) >> 1
                if arr[m] < v:
                    l = m+1
                else:
                    r = m
            return l

        if len(nums) == 0:
            return 0
        if len(nums) == 1:
            return 1

        dp = [nums[0]]
        for n in nums[1:]:
            if n == dp[-1]:
                continue
            if n > dp[-1]:
                dp.append(n)
                continue
            p = bsearch(dp, n)
            dp[p] = n
        return len(dp)
