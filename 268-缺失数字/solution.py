# coding: utf-8

class Solution(object):
    def missingNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        ret = len(nums)
        for i in range(ret):
            ret ^= i
            ret ^= nums[i]
        return ret
