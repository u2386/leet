# coding: utf-8


class Solution(object):
    def trap(self, height):
        """
        :type height: List[int]
        :rtype: int
        """
        if not height:
            return 0

        total = 0
        for h in range(1, max(height)+1):
            i = water = 0
            while i < len(height):
                if height[i] == 0:
                    i += 1
                    continue

                i += 1
                while i < len(height)-1 and height[i] == 0:
                    water += 1
                    i += 1
                if i < len(height) and height[i] == 0:
                    water = 0
                total += water
                water = 0
            height = [max(0, e-1) for e in height]
        return total


def test1():
    h = [4, 2, 3]
    assert Solution().trap(h) == 1


def test2():
    h = [0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]
    assert Solution().trap(h) == 6


def test3():
    h = [2, 1, 0, 2]
    assert Solution().trap(h) == 3


def main():
    h = [0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]
    print(Solution().trap(h))


if __name__ == '__main__':
    main()
