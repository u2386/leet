# coding: utf-8


class Solution(object):
    def calculate(self, s):
        """
        :type s: str
        :rtype: int
        """

        def scan_number(pos):
            ret = 0
            while pos < len(s) and '0' <= s[pos] <= '9':
                ret *= 10
                ret += int(s[pos])
                pos += 1
            return ret, pos

        def skip_whitespace(pos):
            while pos < len(s) and s[pos] == ' ':
                pos += 1
            return pos

        if not s:
            return

        stack = []
        pos = 0
        while pos < len(s):
            pos = skip_whitespace(pos)
            if not (pos < len(s)):
                break

            c = s[pos]

            if c not in '+-*/':
                num, pos = scan_number(pos)
                stack.append(num)
                continue

            pos = skip_whitespace(pos+1)
            if not (pos < len(s)):
                break

            opt = c
            num, pos = scan_number(pos)
            if opt == '+':
                stack.append(num)
            elif opt == '-':
                stack.append(-1 * num)
            elif opt == '*':
                prev = stack.pop()
                stack.append(prev * num)
            elif opt == '/':
                prev = stack.pop()
                if prev * num > 0:
                    stack.append(prev // num)
                else:
                    stack.append((prev + (-prev % num))//num)

        return sum(stack)


def test1():
    s = '3+2*2'
    assert Solution().calculate(s) == 7


def test2():
    s = ' 3/2 '
    assert Solution().calculate(s) == 1


def test3():
    s = ' 3+5 / 2 '
    assert Solution().calculate(s) == 5


def test4():
    s = "14-3/2"
    assert Solution().calculate(s) == 13


def main():
    s = "14-3/2"
    print(Solution().calculate(s))


if __name__ == '__main__':
    main()
