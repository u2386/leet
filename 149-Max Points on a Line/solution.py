# coding: utf-8

from collections import defaultdict


class Solution(object):
    def maxPoints(self, points):
        """
        :type points: List[List[int]]
        :rtype: int
        """

        def _gcd(a, b):
            while b != 0:
                tmp = a % b
                a = b
                b = tmp
            return a

        if len(points) <= 2:
            return len(points)

        ret = 0
        for i in range(len(points)):
            (x0, y0) = points[i]
            dup = 0
            mlen = 0
            slopes = defaultdict(int)
            for j in range(i+1, len(points)):
                x = points[j][0] - x0
                y = points[j][1] - y0
                if x == 0 and y == 0:
                    dup += 1
                    continue
                gcd = _gcd(x, y)
                x //= gcd
                y //= gcd
                slopes[(x, y)] += 1
                mlen = max(mlen, slopes[(x, y)])
            ret = max(ret, mlen + dup + 1)
        return ret


def main():
    arr = [[1, 1], [2, 2], [3, 3]]
    print(Solution().maxPoints(arr))


if __name__ == '__main__':
    main()
