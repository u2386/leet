# coding: utf-8


class Solution(object):
    def maxEnvelopes(self, envelopes):
        """
        :type envelopes: List[List[int]]
        :rtype: int
        """
        if not envelopes:
            return 0

        envelopes.sort()
        dp = [1] * len(envelopes)

        for i in range(1, len(envelopes)):
            (w, h) = envelopes[i]

            j = i - 1
            while j >= 0:
                (pw, ph) = envelopes[j]
                if pw < w and ph < h:
                    dp[i] = max(dp[i], dp[j] + 1)
                j -= 1

        return max(dp)


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
