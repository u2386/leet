class Solution:
    def singleNonDuplicate(self, nums):
        l, r = 0, len(nums)-1
        while l < r:
            m = (l + r) >> 2 << 1
            if nums[m] == nums[m+1]:
                l = m + 2
            else:
                r = m
        return nums[l]

if __name__ == '__main__':
    s = Solution()
    print(s.singleNonDuplicate([1]))
