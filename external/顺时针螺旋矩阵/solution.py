# coding: utf-8

DIRS = [(0, 1), (1, 0), (0, -1), (-1, 0)]


def clockwise(matrix):
    footprint = []
    for row in matrix:
        footprint.append([0 for _ in row])

    ret = []
    pos = 0
    i, j = DIRS[pos]

    fin_x, fin_y = len(matrix) >> 1, len(matrix[0]) >> 1
    x, y = 0, 0
    while True:
        footprint[x][y] = 1
        ret.append(matrix[x][y])

        if (x == fin_x and y == fin_y):
            return ret

        if (
            x + i >= len(matrix)
            or y + j >= len(matrix[0])
            or footprint[x+i][y+j]
        ):
            # Update direction
            pos = (pos + 1) % 4
            i, j = DIRS[pos]

        x += i
        y += j


def test1():
    matrix = [
        [1, 2, 3],
        [4, 5, 6],
        [7, 8, 9],
    ]

    actual = clockwise(matrix)
    assert actual == [1, 2, 3, 6, 9, 8, 7, 4, 5]


def main():
    matrix = [
        [1, 2, 3, 4],
        [5, 6, 7, 8],
        [9, 10, 11, 12],
    ]
    print(clockwise(matrix))


if __name__ == '__main__':
    main()
