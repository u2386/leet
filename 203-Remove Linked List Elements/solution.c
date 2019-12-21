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

    struct ListNode *prev = head;
    while (prev->val == val)
    {
        if (prev->next == NULL)
            return NULL;
        head = prev->next;
        prev = head;
    }
    while (prev->next != NULL)
    {
        if (prev->next->val == val)
        {
            prev->next = prev->next->next;
        }
        else
        {
            prev = prev->next;
        }
    }

    return head;
}

int main()
{
    return 0;
}