# coding: utf-8


class Solution(object):
    def jump(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        steps = 0
        m = 0
        end = 0
        for i, n in enumerate(nums[:-1]):
            m = max(m, nums[i] + i)
            if i == end:
                end = m
                steps += 1
        return steps


def test01():
    arr = [2, 3, 1, 1, 4]
    assert Solution().jump(arr) == 2


def test02():
    arr = [1 for _ in range(2500)]
    assert Solution().jump(arr) == 2499


def test03():
    arr = [0]
    assert Solution().jump(arr) == 0


def test04():
    arr = [1 for _ in range(2500)]
    arr[0] = 2500
    assert Solution().jump(arr) == 1


def test05():
    arr = [3, 2, 1]
    assert Solution().jump(arr) == 1


def test06():
    arr = [2, 3, 1]
    assert Solution().jump(arr) == 1


def main():
    arr = [1 for _ in range(2500)]
    print(Solution().jump(arr))


if __name__ == '__main__':
    main()
