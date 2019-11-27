# coding: utf-8


class Solution(object):
    def maximalSquare(self, matrix):
        """
        :type matrix: List[List[str]]
        :rtype: int
        """

        def spread(pos):
            i, j = pos
            r = 0

            x, y = i, j
            while True:
                x, y = i+r, j+r
                if not (x < len(matrix) and y < len(matrix[x])):
                    break

                row = [matrix[x][c] for c in range(j, y+1)
                       if matrix[x][c] == '1']
                col = [matrix[c][y] for c in range(i, x+1)
                       if matrix[c][y] == '1']
                if (
                    len(row) == len(col) == r+1
                ):
                    r += 1
                    continue
                break

            return r*r

        max_area = 0
        for i in range(len(matrix)):
            for j in range(len(matrix[i])):
                if matrix[i][j] == '0':
                    continue
                max_area = max(max_area, spread((i, j)))
        return max_area


def main():
    matrix = [
        ["1", "0", "1", "0", "0"],
        ["1", "0", "1", "1", "1"],
        ["1", "1", "1", "1", "1"],
        ["1", "0", "0", "1", "0"],
    ]
    print(Solution().maximalSquare(matrix))


if __name__ == '__main__':
    main()
