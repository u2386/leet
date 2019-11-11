# coding:utf-8


class Solution(object):
    def maxAreaOfIsland(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: int
        """

        directors = [(-1, 0), (1, 0), (0, -1), (0, 1)]

        def bfs(row, col):
            area = 0
            q = [(row, col)]
            footprint.add((row, col))
            while q:
                (row, col), q = q[0], q[1:]

                if not grid[row][col]:
                    continue

                area += 1

                for rd, cd in directors:
                    pos = (row+rd, col+cd)
                    new_row, new_col = pos
                    if not (0 <= new_row < len(grid)):
                        continue
                    if not (0 <= new_col < len(grid[new_row])):
                        continue

                    if pos in footprint:
                        continue

                    if grid[new_row][new_col]:
                        footprint.add(pos)
                        q.append(pos)
            return area

        footprint = set()
        max_area = 0
        for row, _ in enumerate(grid):
            for col, _ in enumerate(grid[row]):
                if not grid[row][col]:
                    continue
                if (row, col) in footprint:
                    continue
                max_area = max(max_area, bfs(row, col))
        return max_area


def test01():
    grid = [[0, 0, 0, 0, 0, 0, 0, 0]]
    assert Solution().maxAreaOfIsland(grid) == 0


def test02():
    grid = [
        [1, 1, 0, 0, 0],
        [1, 1, 0, 0, 0],
        [0, 0, 0, 1, 1],
        [0, 0, 0, 1, 1]
    ]
    assert Solution().maxAreaOfIsland(grid) == 4


def test03():
    grid = [
        [0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0],
        [0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0],
        [0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0],
        [0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0],
        [0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0],
        [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0],
        [0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0],
        [0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0]
    ]
    assert Solution().maxAreaOfIsland(grid) == 6


def main():
    grid = [
        [1, 1, 0, 0, 0],
        [1, 1, 0, 0, 0],
        [0, 0, 0, 1, 1],
        [0, 0, 0, 1, 1]
    ]
    print(Solution().maxAreaOfIsland(grid))


if __name__ == '__main__':
    main()
