#coding: utf-8

import pprint


class Solution:
    def longestCommonSubsequence(self, text1: str, text2: str) -> int:
        dp = [[0 for _ in range(len(text2)+1)] for _ in range(len(text1)+1)]
        for i in range(len(text1)):
            for j in range(len(text2)):
                if text1[i] == text2[j]:
                    dp[i+1][j+1] = 1 + dp[i][j]
                else:
                    dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
        pprint.pprint(dp)
        return max(dp[-1])

def test():
    cases = [
        ["abcde", "ace"],
        ["abc", "abc"],
        ["abc", "def"],
        ["abcde", "aceabcde"],
        ["ezupkr", "ubmrapg"],
        ["bsbininm", "jmjkbkjkv"],
    ]
    s = Solution()
    for case in cases:
        print(case)
        t0, t1 = case
        l = s.longestCommonSubsequence(t0, t1)
        print(l)


test()
