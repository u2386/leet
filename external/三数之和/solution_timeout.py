# coding:utf-8


class Solution(object):
    def threeSum(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        if len(nums) < 3:
            return []

        def find_rest(n, nums):
            for m in nums:
                footprint.setdefault(-(m+n), []).append([m, n])

        footprint = {}
        ret = set()
        for i, n in enumerate(nums):
            if footprint.get(n):
                for x, y in footprint[n]:
                    ret.add(tuple(sorted([n, x, y])))
            find_rest(n, nums[:i])
        return [list(t) for t in ret]


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


def main():
    import pdb
    pdb.set_trace()
    print(Solution().threeSum([3, 0, -2, -1, 1, 2]))


if __name__ == '__main__':
    main()
