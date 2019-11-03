# coding: utf-8


class Solution(object):

    def longestCommonPrefix(self, strs):
        """
        :type strs: List[str]
        :rtype: str
        """
        if not strs:
            return ''

        prefix = strs[0]
        for s in strs[1:]:
            prefix = self.find_prefix(prefix, s)
            if prefix == '':
                break
        return prefix


    def find_prefix(self, prefix, target):
        i = 0
        while i < len(prefix) and i < len(target):
            if prefix[i] == target[i]:
                i += 1
                continue
            break
        return prefix[:i]


def test1():
    s = ["flower","flow","flight"]
    assert Solution().longestCommonPrefix(s) == 'fl'


def test2():
    s = ["dog","racecar","car"]
    assert Solution().longestCommonPrefix(s) == ''


def main():
    pass


if __name__ == '__main__':
    main()
