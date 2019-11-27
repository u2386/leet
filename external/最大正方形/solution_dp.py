# coding: utf-8


class Solution(object):
    def maximalSquare(self, matrix):
        """
        :type matrix: List[List[str]]
        :rtype: int
        """
        if not matrix:
            return 0

        dp = [0 for _ in range(len(matrix[0]) + 1)]

        max_len = 0
        prev = 0
        for i in range(len(matrix)):
            for j in range(len(matrix[0])):
                temp = dp[j+1]
                if matrix[i][j] == '0':
                    dp[j+1] = 0

                else:
                    dp[j+1] = min(dp[j], dp[j+1], prev) + 1
                    prev = dp[j+1]
                    max_len = max(max_len, dp[j+1])

                prev = temp

        return max_len * max_len


def test01():
    matrix = [
        ["1", "0", "1", "0"],
        ["1", "0", "1", "1"],
        ["1", "0", "1", "1"],
        ["1", "1", "1", "1"]
    ]
    assert Solution().maximalSquare(matrix) == 4


def test02():
    matrix = [
        ["1", "0", "1", "0", "0"],
        ["1", "0", "1", "1", "1"],
        ["1", "1", "1", "1", "1"],
        ["1", "0", "0", "1", "0"],
    ]
    assert Solution().maximalSquare(matrix) == 4


def main():
    matrix = [
        ["1", "0", "1", "0"],
        ["1", "0", "1", "1"],
        ["1", "0", "1", "1"],
        ["1", "1", "1", "1"]
    ]
    print(Solution().maximalSquare(matrix))


if __name__ == '__main__':
    main()
