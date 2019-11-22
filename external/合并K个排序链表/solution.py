# coding: utf-8

import heapq


class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def mergeKLists(self, lists):
        """
        :type lists: List[ListNode]
        :rtype: ListNode
        """
        ListNode.__lt__ = lambda self, o: True

        heap = [(n.val, n) for n in lists if n]
        if not heap:
            return
        heapq.heapify(heap)

        _, rear = heapq.heappop(heap)
        head = rear
        while heap:
            if rear.next:
                heapq.heappush(heap, (rear.next.val, rear.next))
            _, rear.next = heapq.heappop(heap)
            rear = rear.next
        return head


def main():
    l = [[1, 4, 5], [1, 3, 4], [2, 6]]
    nodes = []
    for values in l:
        nodes.append(ListNode(values[0]))
        rear = nodes[-1]
        for v in values[1:]:
            rear.next = ListNode(v)
            rear = rear.next

    head = Solution().mergeKLists(nodes)
    while head:
        print(head.val)
        head = head.next


if __name__ == '__main__':
    main()
