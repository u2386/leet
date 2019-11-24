# coding: utf-8


class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution(object):
    def zigzagLevelOrder(self, root):
        """
        :type root: TreeNode
        :rtype: List[List[int]]
        """
        if not root:
            return []

        ret = [[]]

        flip = 1
        q, p = [root], []
        while q:
            n, q = q[0], q[1:]
            ret[-1].append(n.val)

            if n.left:
                p.append(n.left)
            if n.right:
                p.append(n.right)

            if not q:
                q, p = p, q
                ret[-1] = ret[-1][::flip]
                flip *= -1
                if q or p:
                    ret.append([])
        return ret


def main():
    pass


if __name__ == '__main__':
    main()
