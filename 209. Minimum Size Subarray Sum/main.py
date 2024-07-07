class Solution:
    def minSubArrayLen(self, target: int, nums):
        n = len(nums)
        ret = 1<<32
        i, j = 0, 0

        s = 0
        while i < n or j < n:
            if s < target and j < n:
                s += nums[j]
                j += 1
            elif s >= target and i < j:
                ret = min(ret, j-i)
                s -= nums[i]
                i +=1
            else:
                break
        if ret == 1<<32:
            return 0
        return ret