class Solution:
    def splitArray(self, nums, k):
        def groups(m):
            g = 1
            s = 0
            for n in nums:
                if n+s > m:
                    s = n
                    g += 1
                    continue
                s += n
            return g

        l = max(nums)
        r = sum(nums) + 1
        while l < r:
            m = (l+r) >> 1
            if groups(m) <= k:
                r = m
            else:
                l = m + 1
        return l

if __name__ == '__main__':
    s = Solution()
    print(s.splitArray([7,2,5,10,8], 2))
    print(s.splitArray([1,2,3,4,5], 2))