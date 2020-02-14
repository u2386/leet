# coding: utf-8


class Node(object):
    def __init__(self, val=0, left=None, right=None, next=None):
        self.val = val
        self.left = left
        self.right = right
        self.next = next


class Solution(object):
    def connect(self, root):
        """
        :type root: Node
        :rtype: Node
        """
        if not root:
            return root

        q, p = [root], []
        while q or p:
            n, q = q[0], q[1:]
            if not n:
                continue

            n.next = q[0] if q else None
            if n.left:
                p.append(n.left)
            if n.right:
                p.append(n.right)
            if not q:
                q, p = p, q
        return root


def main():
    pass


if __name__ == "__main__":
    main()
