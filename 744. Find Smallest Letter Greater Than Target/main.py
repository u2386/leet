class Solution:
    def nextGreatestLetter(self, letters, target):
        l, r = 0, len(letters)
        while l < r:
            m = (l + r) >> 1
            if letters[m] <= target:
                l = m + 1
            else:
                r = m
        if l >= len(letters):
            return letters[0]
        return letters[l]


if __name__ == "__main__":
    s = Solution()
    print(s.nextGreatestLetter(["c","f","j"], "a"))
    print(s.nextGreatestLetter(["c","f","j"], "c"))
