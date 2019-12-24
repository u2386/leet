# coding: utf-8


class TreeNode(object):

   def __init__(self, x):
       self.val = x
       self.left = None
       self.right = None


class Solution(object):
    def kthSmallest(self, root, k):
        """
        :type root: TreeNode
        :type k: int
        :rtype: int
        """
        stack = []
        while True:
            while root:
                stack.append(root)
                root = root.left

            if not stack:
                break

            root = stack.pop()
            k -= 1
            if k == 0:
                return root.val

            root = root.right



def main():
    pass


if __name__ == '__main__':
    main()
