# coding: utf-8


class Solution(object):
    def convert(self, s, numRows):
        """
        :type s: str
        :type numRows: int
        :rtype: str
        """
        if len(s) < numRows or numRows == 1:
            return s

        pat = ['' for _ in range(len(s))]

        x, d = 0, 1
        for c in s:
            pat[x] += c
            x += d

            if x == numRows - 1 or x == 0:
                d *= -1
        return ''.join(pat)


def main():
    print(Solution().convert('LEETCODEISHIRING', 3))


if __name__ == '__main__':
    main()
