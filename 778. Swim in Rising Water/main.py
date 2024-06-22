dirs = [(-1, 0), (1, 0), (0, -1), (0, 1)]

class Solution:
    def swimInWater(self, grid):

        def bsearch(q, v):
            l, r = 0, len(q)
            while l < r:
                m = (l+r) >> 1
                if v <= q[m][0]:
                    r = m
                else:
                    l = m + 1
            return l

        n = len(grid[0])
        seen = [0 for _ in range(n*n)]
        q = [(grid[0][0], (0, 0))]
        seen[0] = True
        while q:
            (t0, (x, y)), q = q[0], q[1:]
            if x == n-1 and y == n-1:
                return t0
            for (dx, dy) in dirs:
                xi = x+dx
                yi = y+dy
                if not (0<=xi<n and 0<=yi<n):
                    continue
                if seen[n*xi+yi]:
                    continue
                seen[n*xi+yi] = True
                t = max(t0,grid[xi][yi])
                p = bsearch(q, t)
                q = q[:p] + [(t, (xi, yi))] + q[p:]

        return -1


if __name__ == '__main__':
    s = Solution()
    print(s.swimInWater([[0,3],[1,2]]))