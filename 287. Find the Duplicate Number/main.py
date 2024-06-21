class Solution:
    def findDuplicate(self, nums: List[int]) -> int:
        ret = 0
        for b in range(32):
            a, e = 0, 0
            for i in range(len(nums)):
                if (nums[i] >> b) & 1:
                    a += 1
                if (i >> b) & 1:
                    e += 1
            if a > e:
                ret |= 1<<i
        return ret


class Solution2:
    def findDuplicate(self, nums: List[int]) -> int:
        l = 1
        r = len(nums)
        while l < r:
            m = (l+r) >> 1
            count = 0
            for n in nums:
                if n <= m:
                    count += 1
            if count <= m:
                l = m+1
                continue
            r = m
        return l