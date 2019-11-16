# coding: utf-8


class Solution(object):
    def merge(self, intervals):
        """
        :type intervals: List[List[int]]
        :rtype: List[List[int]]
        """
        if not intervals:
            return []

        intervals.sort()

        ret = [intervals[0]]
        for i, j in intervals[1:]:
            if i < ret[-1][0]:
                ret[-1][0] = i

            if i <= ret[-1][1]:
                ret[-1][1] = max(j, ret[-1][1])
                continue

            ret.append([i, j])
        return ret


def test01():
    arr = [[1, 3], [2, 6], [8, 10], [15, 18]]
    assert Solution().merge(arr) == [[1, 6], [8, 10], [15, 18]]


def test02():
    arr = [[1, 4], [4, 5]]
    assert Solution().merge(arr) == [[1, 5]]


def test03():
    arr = [[1, 4], [0, 0]]
    assert Solution().merge(arr) == [[0, 0], [1, 4]]


def main():
    arr = [[1, 4], [4, 5]]
    print(Solution().merge(arr))


if __name__ == '__main__':
    main()
