# coding: utf-8


class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

    def __eq__(self, o):
        return self.val == o.val


class Solution(object):
    def lowestCommonAncestor(self, root, p, q):
        """
        :type root: TreeNode
        :type p: TreeNode
        :type q: TreeNode
        :rtype: TreeNode
        """
        def find_path(root):
            if stop[0] or not root:
                return

            path.append(root)

            find_path(root.left)
            if root.val == p.val:
                p_path[:] = path[:]
            if root.val == q.val:
                q_path[:] = path[:]
            find_path(root.right)

            path.pop()

            if q_path and p_path:
                stop[0] = True

        stop = [False]
        path = []
        p_path, q_path = [], []
        find_path(root)

        n = None
        while p_path and q_path and p_path[0] == q_path[0]:
            n = p_path[0]
            p_path = p_path[1:]
            q_path = q_path[1:]
        return n


def test1():
    # [1,2,3,null,4]
    root = TreeNode(1)
    root.left = TreeNode(2)
    root.right = TreeNode(3)
    root.left.right = TreeNode(4)

    actual = Solution().lowestCommonAncestor(root, TreeNode(4), TreeNode(3))
    assert actual.val == 1


def test2():
    # [3,5,1,6,2,0,8,null,null,7,4]
    root = TreeNode(3)
    left = root.left = TreeNode(5)
    right = root.right = TreeNode(1)

    left.left = TreeNode(6)
    left.right = TreeNode(2)
    right.left = TreeNode(0)
    right.right = TreeNode(8)

    left.right.left = TreeNode(7)
    left.right.right = TreeNode(4)

    actual = Solution().lowestCommonAncestor(root, TreeNode(5), TreeNode(4))
    assert actual.val == 5


def main():
    root = TreeNode(3)
    left = root.left = TreeNode(5)
    right = root.right = TreeNode(1)

    left.left = TreeNode(6)
    left.right = TreeNode(2)
    right.left = TreeNode(0)
    right.right = TreeNode(8)

    left.right.left = TreeNode(7)
    left.right.right = TreeNode(4)

    actual = Solution().lowestCommonAncestor(root, TreeNode(5), TreeNode(4))
    print(actual.val)


if __name__ == '__main__':
    main()
