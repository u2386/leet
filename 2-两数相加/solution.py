# coding: utf-8

# Definition for singly-linked list.
class ListNode(object):
   def __init__(self, x):
       self.val = x
       self.next = None

class Solution(object):
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        head = r = ListNode(None)
        flag = False

        while (l1 and l2):
            s = l1.val + l2.val
            if flag:
                s += 1

            r.next = ListNode(s % 10)
            l1 = l1.next
            l2 = l2.next
            r = r.next

            flag = s >= 10

        while l1:
            s = l1.val
            if flag:
                s += 1
            r.next = ListNode(s % 10)

            r = r.next
            l1 = l1.next
            flag = s >= 10

        while l2:
            s = l2.val
            if flag:
                s += 1
            r.next = ListNode(s % 10)

            r = r.next
            l2 = l2.next
            flag = s >= 10

        if flag:
            r.next = ListNode(1)

        return head.next


def main():
    l1 = ListNode(2)
    l1.next = ListNode(4)
    l1.next.next = ListNode(3)
    l1.next.next.next = ListNode(5)

    l2 = ListNode(5)
    l2.next = ListNode(6)
    l2.next.next = ListNode(4)

    s = Solution()
    r = s.addTwoNumbers(l1, l2)
    while r:
        print(r.val, end='')
        r = r.next
    print()


main()
