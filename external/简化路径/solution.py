# coding: utf-8


class Solution(object):
    def simplifyPath(self, path):
        """
        :type path: str
        :rtype: str
        """
        ret = ['']
        for d in path.split('/'):
            if not d or d == '.':
                continue
            elif d == '..':
                if ret[-1] != '':
                    ret.pop()
                continue
            ret.append(d)

        if len(ret) <= 1:
            return '/'
        return '/'.join(ret)


def test1():
    path = '/home/'
    assert Solution().simplifyPath(path) == '/home'


def test2():
    path = '/../'
    assert Solution().simplifyPath(path) == '/'


def test3():
    path = '/home//foo/'
    assert Solution().simplifyPath(path) == '/home/foo'


def test4():
    path = '/a/./b/../../c/'
    assert Solution().simplifyPath(path) == '/c'


def test5():
    path = '/a/./b/../../c/'
    assert Solution().simplifyPath(path) == '/c'


def test6():
    path = '/a//b////c/d//././/..'
    assert Solution().simplifyPath(path) == '/a/b/c'


def main():
    path = ''
    print(Solution().simplifyPath(path))


if __name__ == '__main__':
    main()
