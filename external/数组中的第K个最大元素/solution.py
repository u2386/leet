# coding: utf-8

import heapq


class Solution(object):
    def findKthLargest(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: int
        """
        q = nums[:k]
        heapq.heapify(q)

        for n in nums[k:]:
            if n < q[0]:
                continue
            heapq.heappushpop(q, n)
        return q[0]


def test01():
    assert Solution().findKthLargest([3, 2, 1, 5, 6, 4], 2) == 5


def test02():
    assert Solution().findKthLargest([3, 2, 3, 1, 2, 4, 5, 5, 6], 4) == 4


def main():
    nums = []
    k = 0
    print(Solution().findKthLargest(nums, k))


if __name__ == '__main__':
    main()
