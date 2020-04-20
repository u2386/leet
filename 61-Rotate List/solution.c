#include <stdio.h>
#include <stdlib.h>

struct ListNode
{
    int val;
    struct ListNode *next;
};

struct ListNode *rotateRight(struct ListNode *head, int k)
{
    if (head == NULL || k == 0)
        return head;

    int length = 1;
    struct ListNode *tail = head;
    while (tail->next != NULL)
    {
        tail = tail->next;
        length++;
    }

    int diff = length - (k % length) - 1;
    struct ListNode *p = head;
    while (diff > 0)
    {
        p = p->next;
        diff--;
    }

    tail->next = head;
    head = p->next;
    p->next = NULL;
    return head;
}

int main()
{
    return 0;
}
