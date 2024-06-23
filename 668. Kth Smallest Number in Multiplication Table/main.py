class Solution:
    def findKthNumber(self, m: int, n: int, k: int) -> int:

        def index(x):
            ret = 0
            for i in range(1, m+1):
                ret += min(n, x//i)
            return ret

        l, r = 1, n*m+1
        while l < r:
            x = (l + r) >> 1
            if k > index(x):
                l = x + 1
            else:
                r = x
        return l


if __name__ == '__main__':
    s = Solution()
    print(s.findKthNumber(3, 3, 5))