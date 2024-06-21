# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def rob(self, root):

        def _rob(root):
            if not root:
                return 0,0
            left = _rob(root.left)
            right = _rob(root.right)

            return max(left[0], left[1]) + max(right[0], right[1]), root.val+left[0]+right[0]
        return max(_rob(root))


if __name__ == '__main__':
    s = Solution()
    n1 = TreeNode(3)
    n2 = TreeNode(2)
    n3 = TreeNode(3)
    n4 = TreeNode(3)
    n5 = TreeNode(1)
    n1.left = n2
    n1.right = n3
    n2.right = n4
    n3.right = n5
    print(s.rob(n1))  # 7
    n1 = TreeNode(3)
    n2 = TreeNode(4)
    n3 = TreeNode(5)
    n4 = TreeNode(1)
    n5 = TreeNode(3)
    n6 = TreeNode(1)
    n1.left = n2
    n1.right = n3
    n2.left = n4
    n2.right = n5
    n3.right = n6
    print(s.rob(n1))  # 9
    n1 = TreeNode(4)
    n2 = TreeNode(1)
    n3 = TreeNode(2)
    n4 = TreeNode(3)
    n1.left = n2
    n2.left = n3
    n3.left = n4
    print(s.rob(n1))  # 7
    n1 = TreeNode(2)
    n2 = TreeNode(1)
    n3 = TreeNode(3)
    n1.left = n2
    n1.right = n3
    print(s.rob(n1))  # 4
    n1 = TreeNode(5)
    n2 = TreeNode(2)
    n3 = TreeNode(6)
    n4 = TreeNode(1)
    n5 = TreeNode(3)
    n1.left = n2
    n1.right = n3
    n2.left = n4
    n2.right = n5
    print(s.rob(n1))  # 10
    n1 = TreeNode(5)
    n2 = TreeNode(2)
    n3 = TreeNode(6)
    n4 = TreeNode(1)
    n5 = TreeNode(3)
    n6 = TreeNode(4)
    n1.left = n2
    n1.right = n3
    n2.left = n4
    n2.right = n5
    n5.right = n6
    print(s.rob(n1))  # 12