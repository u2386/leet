# coding: utf-8


class Solution(object):
    def maxEnvelopes(self, envelopes):
        """
        :type envelopes: List[List[int]]
        :rtype: int
        """
        if not envelopes:
            return 0

        def lis(nums):
            if not nums:
                return 0

            piles = 1
            tops = [nums[0]]
            for n in nums[1:]:
                left, right = 0, len(tops)

                while left < right:
                    mid = (left + right) >> 1
                    if tops[mid] < n:
                        left = mid + 1
                    else:
                        right = mid

                if left == piles:
                    piles += 1
                    tops.append(n)
                    continue
                tops[left] = n
            return piles

        envelopes.sort(key=lambda e: (e[0], -e[1]))
        heights = [e[1] for e in envelopes]
        return lis(heights)


def test1():
    arr = [[5, 4], [6, 4], [6, 7], [2, 3]]
    assert Solution().maxEnvelopes(arr) == 3


def test2():
    arr = [[2, 100], [3, 200], [4, 300], [5, 500],
           [5, 400], [5, 250], [6, 370], [6, 360], [7, 380]]
    assert Solution().maxEnvelopes(arr) == 5


def test3():
    arr = [[46, 89], [50, 53], [52, 68], [72, 45], [77, 81]]
    assert Solution().maxEnvelopes(arr) == 3


def main():
    arr = [[46, 89], [50, 53], [52, 68], [72, 45], [77, 81]]
    print(Solution().maxEnvelopes(arr))


if __name__ == '__main__':
    main()
