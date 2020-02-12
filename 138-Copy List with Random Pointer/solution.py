# coding: utf-8


class Node:

    def __init__(self, x, next=None, random=None):
        self.val = int(x)
        self.next = next
        self.random = random


class Solution(object):

    def copyRandomList(self, head):
        """
        :type head: Node
        :rtype: Node
        """
        if not head:
            return

        tmp = head
        while tmp:
            n = Node(tmp.val, tmp.next)
            tmp.next = n
            tmp = n.next

        tmp = head
        while tmp:
            tmp.next.random = tmp.random.next if tmp.random else None
            tmp = tmp.next.next

        clone = tmp = Node(-1)
        while head:
            tmp.next = head.next
            head.next = head.next.next
            head = head.next
            tmp = tmp.next
        return clone.next


def main():
    pass


if __name__ == "__main__":
    main()
