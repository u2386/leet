class Solution:
    def numMatchingSubseq(self, s, words):
        index = {}
        for i in range(len(s)):
            if not index.get(s[i]):
                index[s[i]] = []
            index[s[i]].append(i)

        def bsearch(arr, p):
            l, r = 0, len(arr)
            while l < r:
                m = (l+r) >> 1
                if arr[m] <= p:
                    l = m + 1
                else:
                    r = m
            if l >= len(arr):
                return -1
            return arr[l]

        count = 0
        for word in words:
            prev = -1
            for c in word:
                idx = index.get(c)
                if not idx:
                    prev = -1
                    break

                prev = bsearch(idx, prev)
                if prev == -1:
                    break
            if prev != -1:
                count += 1
        return count


if __name__ == '__main__':
    s = Solution()
    print(s.numMatchingSubseq("abcde", ["bb"]))