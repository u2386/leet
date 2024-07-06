class Solution:
    def combinationSum3(self, k: int, n: int):
        nums = [i for i in range(1, 10)]

        ret = []
        def comb(temp, start, seen):
            if len(temp) == k:
                if sum(temp) == n:
                    ret.append(temp[:])
                return

            for i in range(start, len(nums)):
                if seen[i]:
                    continue

                seen[i] = 1
                temp.append(nums[i])
                comb(temp[:], i+1, seen)
                temp = temp[:-1]
                seen[i] = 0

        comb([], 0, [0 for _ in range(10)])
        return ret

if __name__ == '__main__':
    s = Solution()
    print(s.combinationSum3(3, 7))
    print(s.combinationSum3(3, 9))
    print(s.combinationSum3(4, 1))