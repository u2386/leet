#include <stdio.h>
#include <stdlib.h>

/**
 * Definition for singly-linked list.
 */
struct ListNode
{
    int val;
    struct ListNode *next;
};

struct ListNode *removeElements(struct ListNode *head, int val)
{
    if (head == NULL)
    {
        return head;
    }

    struct ListNode *dummy = (struct ListNode *)malloc(sizeof(struct ListNode));
    dummy->next = head;

    struct ListNode *prev = dummy;
    while (prev != NULL && prev->next != NULL)
    {
        if (prev->next->val == val)
        {
            prev->next = prev->next->next;
            continue;
        }
        prev = prev->next;
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