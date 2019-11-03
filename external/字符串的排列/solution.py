# coding: utf-8


class Solution(object):

    def checkInclusion(self, s1, s2):
        """
        :type s1: str
        :type s2: str
        :rtype: bool
        """
        if not s1:
            return True
        if not s2:
            return False
        if len(s1) > len(s2):
            return False

        i, j = 0, len(s1) - 1
        s1_counter = self.get_counter(s1)
        s2_counter = self.get_counter(s2[i:j+1])

        while j < len(s2):
            if self.equals(s1_counter, s2_counter):
                return True

            s2_counter[s2[i]] -= 1
            i += 1

            j += 1
            if not j < len(s2):
                break
            s2_counter[s2[j]] = s2_counter.get(s2[j], 0) + 1

        return False

    def get_counter(self, s):
        ret = {}
        for c in s:
            ret[c] = ret.get(c, 0) + 1
        return ret

    def equals(self, c0, c1):
        for k, v in c0.items():
            if v == c1.get(k):
                continue
            return False
        return True


def test1():
    s1 = 'ab'
    s2 = 'eidbaooo'
    assert Solution().checkInclusion(s1, s2)


def test2():
    s1 = 'ab'
    s2 = 'eidboaoo'
    assert not Solution().checkInclusion(s1, s2)


def main():
    s1 = 'ab'
    s2 = 'eidboaoo'
    import pdb; pdb.set_trace()
    print(Solution().checkInclusion(s1, s2))


if __name__ == '__main__':
    main()
