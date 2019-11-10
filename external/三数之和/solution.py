# coding:utf-8


class Solution(object):
    def threeSum(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        if len(nums) < 3:
            return []

        nums.sort()
        ret = []
        for i, n in enumerate(nums):
            if n > 0:
                return ret
            if i > 0 and n == nums[i-1]:
                continue

            l, r = i+1, len(nums) - 1
            while l < r:
                s = n + nums[l] + nums[r]
                if s == 0:
                    ret.append([n, nums[l], nums[r]])
                    while l < r and nums[l] == nums[l+1]:
                        l += 1
                    while l < r and nums[r] == nums[r-1]:
                        r -= 1

                    l += 1
                    r -= 1

                elif s > 0:
                    r -= 1
                else:
                    l += 1
        return ret


def test01():
    assert Solution().threeSum([1, 2, -2, -1]) == []


def test02():
    assert sorted(Solution().threeSum([-1, 0, 1, 2, -1, -4])) == sorted([
        [-1, 0, 1],
        [-1, -1, 2]
    ])


def test03():
    assert sorted(Solution().threeSum([3, 0, -2, -1, 1, 2])) == sorted([
        [-2, -1, 3], [-2, 0, 2], [-1, 0, 1]
    ])


def test04():
    assert sorted(Solution().threeSum([0, 0, 0])) == sorted([
        [0, 0, 0]
    ])


def main():
    print(Solution().threeSum([0, 0, 0]))


if __name__ == '__main__':
    main()
