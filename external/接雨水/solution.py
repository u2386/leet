# coding: utf-8


class Solution(object):
    def trap(self, height):
        """
        :type height: List[int]
        :rtype: int
        """
        if not height:
            return 0

        water = 0
        stack = []
        cur = 0
        while cur < len(height):
            while stack and height[cur] > height[stack[-1]]:
                h = height[stack.pop()]
                if not stack:
                    break

                dist = cur - stack[-1] - 1
                water += dist * (min(height[cur], height[stack[-1]]) - h)

            stack.append(cur)
            cur += 1

        return water


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
