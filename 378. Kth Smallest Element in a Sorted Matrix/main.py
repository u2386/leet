class Solution:
    def kthSmallest(self, matrix, k):
        n = len(matrix)

        def index(x):
            count = 0
            for i in range(n):
                if matrix[i][0] > x:
                    break
                for j in range(n):
                    if matrix[i][j] <= x:
                        count +=1
                        continue
                    break
            return count

        l, r = matrix[0][0], matrix[-1][-1]
        while l < r:
            m = (l+r) >> 1
            if k <= index(m):
                r = m
            else:
                l = m + 1
        return l


if __name__ == '__main__':
    s = Solution()
    print(s.kthSmallest([[1,5,9],[10,11,13],[12,13,15]], 8))
    print(s.kthSmallest([[-5]], 1))
    print(s.kthSmallest([[1,2],[1,3]], 1))
    print(s.kthSmallest([[1,4],[2,5]], 2))