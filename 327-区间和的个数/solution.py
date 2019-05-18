# coding: utf-8


class Solution(object):

    count = 0

    def countRangeSum(self, nums, lower, upper):
        """
        :type nums: List[int]
        :type lower: int
        :type upper: int
        :rtype: int
        """
        length = len(nums)
        if length == 0:
            return 0
        elif length == 1:
            return int(lower <= nums[0] <= upper)

        sums = [nums[0]]
        for i in range(1, length):
            sums.append(sums[i-1] + nums[i])
        self.mergeCount(sums, lower, upper)
        return self.count

    def mergeCount(self, sums, lower, upper):
        length = len(sums)
        if length <= 1:
            if length == 1 and lower <= sums[0] <= upper:
                self.count += 1
            return sums

        mid = length >> 1
        left = self.mergeCount(sums[:mid], lower, upper)
        right = self.mergeCount(sums[mid:], lower, upper)

        for l in left:
            i = j = 0
            while i < len(right):
                if right[i] - l >= lower:
                    break
                i += 1
            while j < len(right):
                if right[j] - l > upper:
                    break
                j += 1
            self.count += (j - i)

        ret = []
        i = j = 0
        while i < len(left) and j < len(right):
            if left[i] <= right[j]:
                ret.append(left[i])
                i += 1
            else:
                ret.append(right[j])
                j += 1

        if i < len(left):
            ret.extend(left[i:])
        if j < len(right):
            ret.extend(right[j:])
        return ret


def main():
    s = Solution()
    case = [-2, 5, -1]
    print(s.countRangeSum(case, -17, -10))


main()
