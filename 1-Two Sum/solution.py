# coding: utf-8


class Solution(object):

    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        m = {}
        for i, n in enumerate(nums):
            d = target - n
            if d in m:
                return sorted([i, m[d]])
            m[n] = i


def test1():
    nums = [3,2,4]
    target = 6
    assert Solution().twoSum(nums, target) == [1, 2]


def main():
    nums = [2, 7 , 11, 15]
    target = 9

    s = Solution()
    print(s.twoSum(nums, target))


if __name__ == '__main__':
    main()
