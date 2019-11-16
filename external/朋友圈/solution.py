# coding: utf-8


class Solution(object):
    def findCircleNum(self, M):
        """
        :type M: List[List[int]]
        :rtype: int
        """

        def find_circle(row):
            q = [row]
            while q:
                col, q = q[0], q[1:]
                footprint[col] = 1
                for r in range(len(M[0])):
                    if not footprint[r] and M[r][col]:
                        q.append(r)

        if len(M) == 1:
            return 1

        footprint = [0 for _ in range(len(M[0]))]
        circle = 0
        for r in range(len(M)):
            if footprint[r]:
                continue

            footprint[r] = 1
            if not M[r][r]:
                continue
            find_circle(r)
            circle += 1
        return circle


def test1():
    assert Solution().findCircleNum([
        [1, 1, 0],
        [1, 1, 0],
        [0, 0, 1]]) == 2


def test2():
    assert Solution().findCircleNum([
        [1, 1, 0],
        [1, 1, 1],
        [0, 1, 1]]) == 1


def test3():
    assert Solution().findCircleNum([
        [1, 0, 0, 1],
        [0, 1, 1, 0],
        [0, 1, 1, 1],
        [1, 0, 1, 1]]) == 1


def main():
    print(Solution().findCircleNum([
        [1, 0, 0, 1],
        [0, 1, 1, 0],
        [0, 1, 1, 1],
        [1, 0, 1, 1]]))


if __name__ == '__main__':
    main()
