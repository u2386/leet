# coding: utf-8


class TreeNode(object):

    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class BSTIterator(object):

    def __init__(self, root):
        """
        :type root: TreeNode
        """
        self._next = None

        stack = []
        while root:
            stack.append(root)
            root = root.left
        self.stack = stack

    def next(self):
        """
        @return the next smallest number
        :rtype: int
        """
        if not self.hasNext():
            return None

        stack = self.stack

        node = stack.pop()
        v = node.val
        if not node.right:
            return v

        cur = node.right
        while cur:
            stack.append(cur)
            cur = cur.left

        return v

    def hasNext(self):
        """
        @return whether we have a next smallest number
        :rtype: bool
        """
        return bool(self.stack)


def main():
    pass


if __name__ == '__main__':
    main()
