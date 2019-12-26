#include <stdio.h>

/**
 * Definition for singly-linked list.
 */
struct ListNode
{
    int val;
    struct ListNode *next;
};

struct ListNode *deleteDuplicates(struct ListNode *head)
{
    struct ListNode *prev = head;
    while (prev && prev->next)
    {
        if (prev->val == prev->next->val)
        {
            prev->next = prev->next->next;
            continue;
        }
        prev = prev->next;
    }
    return head;
}

int main()
{
    return 0;
}
