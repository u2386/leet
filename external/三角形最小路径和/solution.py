# coding: utf-8


class Solution(object):
    def minimumTotal(self, triangle):
        """
        :type triangle: List[List[int]]
        :rtype: int
        """
        dp = [0 for _ in range(len(triangle)+1)]

        for i in range(len(triangle)-1, -1, -1):
            for j in range(len(triangle[i])):
                dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
        return dp[0]


def main():
    arr = [
        [2],
        [3, 4],
        [6, 5, 7],
        [4, 1, 8, 3]
    ]

    print(Solution().minimumTotal(arr))


if __name__ == '__main__':
    main()
