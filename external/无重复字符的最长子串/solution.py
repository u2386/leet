# coding: utf-8


class Solution(object):

    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        if len(s) <= 1:
            return len(s)

        m = {}
        max_length = 0
        i, j = 0, 1

        m[s[i]] = i
        while i != j and j < len(s):
            if s[j] not in m:
                m[s[j]] = j
                j += 1
                continue

            max_length = max(max_length, len(m))

            p = m[s[j]] + 1
            while i < p:
                del m[s[i]]
                i += 1

            m[s[j]] = j
            j += 1

        return max(max_length, len(m))


def test1():
    s = 'abcabcbb'
    assert Solution().lengthOfLongestSubstring(s) == 3


def test2():
    s = 'bbbbb'
    assert Solution().lengthOfLongestSubstring(s) == 1


def test3():
    s = 'pwwkew'
    assert Solution().lengthOfLongestSubstring(s) == 3


def test4():
    s = " "
    assert Solution().lengthOfLongestSubstring(s) == 1


def test5():
    s = "aab"
    assert Solution().lengthOfLongestSubstring(s) == 2


def test6():
    s = "dvdf"
    assert Solution().lengthOfLongestSubstring(s) == 3


def main():
    pass


if __name__ == '__main__':
    main()
