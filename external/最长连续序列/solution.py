class Solution(object):
    def longestConsecutive(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        s = set(nums)

        max_len = 0
        for n in nums:
            if n - 1 not in s:
                l = 0
                while n in s:
                    l += 1
                    n += 1
                max_len = max(max_len, l)
        return max_len


def test01():
    arr = [100, 4, 200, 1, 3, 2]
    assert Solution().longestConsecutive(arr) == 4


def main():
    print(Solution().longestConsecutive([100, 4, 200, 1, 3, 2]))


if __name__ == '__main__':
    main()
