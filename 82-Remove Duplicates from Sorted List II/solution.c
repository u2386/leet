#include <stdlib.h>
#include <stdio.h>

struct ListNode
{
    int val;
    struct ListNode *next;
};

struct ListNode *deleteDuplicates(struct ListNode *head)
{
    if (head == NULL)
        return head;

    struct ListNode *dummy = (struct ListNode *)malloc(sizeof(struct ListNode));
    dummy->next = head;
    struct ListNode *p = dummy;
    while (p->next && p->next->next)
    {
        if (p->next->val != p->next->next->val)
        {
            p = p->next;
            continue;
        }

        int t = p->next->val;
        while (p->next && p->next->val == t)
            p->next = p->next->next;
    }

    head = dummy->next;

    dummy->next = NULL;
    free(dummy);

    return head;
}

int main()
{
    return 0;
}
