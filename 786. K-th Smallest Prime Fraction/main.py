class Solution:
    def kthSmallestPrimeFraction(self, arr, k):
        n = len(arr)
        l = 0
        r = 1

        while l < r:
            m = (l + r) / 2
            p, q = 0, 0
            mx = 0.0

            total = 0
            j = 0
            for i in range(j, n-1):
                while (j < n and arr[i] > arr[j] * m):
                    j += 1
                total += (n - j)
                if j == n:
                    break
                f = arr[i] / arr[j]
                if f > mx:
                    mx = f
                    p = i
                    q = j
            if total == k:
                return [arr[p], arr[j]]
            elif total < k:
                l = m + 1
            else:
                r = m
        return []


if __name__ == '__main__':
    s = Solution()
    print(s.kthSmallestPrimeFraction([1,7], 1))