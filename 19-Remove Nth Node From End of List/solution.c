#include <stdlib.h>
#include <stdio.h>

struct ListNode
{
    int val;
    struct ListNode *next;
};

struct ListNode *removeNthFromEnd(struct ListNode *head, int n)
{
    if (n <= 0)
        return  head;

    struct ListNode *dummy = (struct ListNode *)malloc(sizeof(struct ListNode));
    dummy->next = head;

    struct ListNode *p = head;
    while (n > 0)
    {
        if (p == NULL)
            return head;
        p = p->next;
        n--;
    }

    struct ListNode *q = dummy;
    while (p != NULL)
    {
        p = p->next;
        q = q->next;
    }

    p = q->next;
    q->next = p->next;
    p->next = NULL;

    head = dummy->next;
    dummy->next = NULL;
    free(dummy);
    return head;
}

int main()
{
    return 0;
}
