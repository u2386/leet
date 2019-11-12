#coding: utf-8


class Solution(object):
    def restoreIpAddresses(self, s):
        """
        :type s: str
        :rtype: List[str]
        """
        def valid(seg):
            return int(seg) < 256 if seg[0] != '0' else len(seg) == 1

        def split(pos, count):
            if count == 0:
                if valid(s[pos:]):
                    segments.append(s[pos:])
                    ret.append('.'.join(segments))
                    segments.pop()
                return

            for i in range(pos+1, min(pos+4, len(s))):
                seg = s[pos:i]
                if not valid(seg):
                    return

                segments.append(seg)
                split(i, count-1)
                segments.pop()

        segments = []
        ret = []
        split(0, 3)
        return ret


def test01():
    assert Solution().restoreIpAddresses('25525511135') == [
        '255.255.11.135',
        '255.255.111.35',
    ]


def test02():
    assert Solution().restoreIpAddresses('0000') == [
        '0.0.0.0'
    ]


def test03():
    assert Solution().restoreIpAddresses('010010') == [
        '0.10.0.10',
        '0.100.1.0',
    ]


def main():
    print(Solution().restoreIpAddresses('0000'))


if __name__ == '__main__':
    main()
