#include <stdlib.h>
#include <stdio.h>
#include <assert.h>

/**
 * Definition for singly-linked list.
 */
struct ListNode
{
    int val;
    struct ListNode *next;
};

struct ListNode *mergeTwoLists(struct ListNode *l1, struct ListNode *l2)
{
    if (l1 == NULL)
        return l2;
    if (l2 == NULL)
        return l1;

    struct ListNode *head = NULL;
    head = (struct ListNode *)malloc(sizeof(struct ListNode));
    struct ListNode *current = NULL;
    current = head;

    while (l1 != NULL && l2 != NULL)
    {
        if (l1->val < l2->val)
        {
            current->next = l1;
            l1 = l1->next;
        }
        else
        {
            current->next = l2;
            l2 = l2->next;
        }

        if (head == NULL)
            head = current;
        current = current->next;
    }

    if (l1 != NULL)
        current->next = l1;
    if (l2 != NULL)
        current->next = l2;

    return head->next;
}

int main()
{
    return 0;
}
