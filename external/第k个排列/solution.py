# coding: utf-8


class Solution(object):
    def getPermutation(self, n, k):
        """
        :type n: int
        :type k: int
        :rtype: str
        """

        def get_factorial(n):
            ret = [1]
            for i in range(1, n):
                r = 1
                for a in range(1, i+1):
                    r *= a
                ret.append(r)
            return ret

        k -= 1
        nums = list(range(1, n+1))
        factorial = get_factorial(n+1)[::-1]
        if k > factorial[0]:
            return ''

        ret = []
        used = [False for _ in range(n)]
        for i, count in enumerate(factorial[1:]):
            offset = k // count
            k = k if k < count else k - (count*offset)
            for pos in range(n):
                if used[pos]:
                    continue
                if offset == 0:
                    break
                offset -= 1
            ret.append(nums[pos])
            used[pos] = True

        return ''.join([str(i) for i in ret])


def test1():
    assert Solution().getPermutation(3, 3) == '213'


def test2():
    assert Solution().getPermutation(4, 9) == '2314'


def test3():
        assert Solution().getPermutation(3, 5) == '312'


def test4():
        assert Solution().getPermutation(4, 1) == '1234'


def test5():
        assert Solution().getPermutation(4, 2) == '1243'


def main():
    print(Solution().getPermutation(3, 5))


if __name__ == '__main__':
    main()
