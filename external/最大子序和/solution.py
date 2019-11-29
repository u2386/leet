# coding: utf-8


class Solution(object):
    def maxSubArray(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        dp = [0 for _ in range(len(nums))]
        dp[0] = nums[0]

        for i in range(1, len(nums)):
            dp[i] = max(nums[i], dp[i-1]+nums[i])
        return max(dp)


def test1():
    arr = [-2, 1, -3, 4, -1, 2, 1, -5, 4]
    assert Solution().maxSubArray(arr) == 6


def test2():
    arr = [-2, -1]
    assert Solution().maxSubArray(arr) == -1


def test3():
    arr = [1]
    assert Solution().maxSubArray(arr) == 1


def main():
    pass


if __name__ == '__main__':
    main()
