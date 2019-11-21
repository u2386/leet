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

struct ListNode *detectCycle(struct ListNode *head)
{
    if (!head || !head->next)
        return NULL;

    struct ListNode *slow, *fast;
    slow = head;
    fast = slow;

    while (fast && fast->next)
    {
        slow = slow->next;
        fast = fast->next->next;

        if (slow == fast)
        {
            fast = head;
            break;
        }
    }

    if (!(fast && fast->next))
        return NULL;

    while (slow != fast)
    {
        slow = slow->next;
        fast = fast->next;
    }
    return slow;
}

struct ListNode *createLinkedList(int *arr, int size)
{
    if (size < 1)
        return NULL;

    struct ListNode *head, *rear;
    head = (struct ListNode *)malloc(sizeof(struct ListNode));
    if (head == NULL)
        return NULL;
    head->val = arr[0];
    head->next = NULL;
    rear = head;

    for (int i = 1; i < size; i++)
    {
        rear->next = (struct ListNode *)malloc(sizeof(struct ListNode));
        if (rear == NULL)
            return NULL;
        rear = rear->next;
        rear->val = arr[i];
        rear->next = NULL;
    }
    return head;
}

int main()
{
    int pos = 1;
    int arr[] = {3, 2, 0, -4};
    struct ListNode *L = createLinkedList(arr, sizeof(arr) / sizeof(int));

    struct ListNode *rear = L, *p;
    for (int i = 1; i < sizeof(arr) / sizeof(int); i++)
    {
        rear = rear->next;
        if (i == pos)
            p = rear;
    }
    rear->next = p;

    struct ListNode *c = detectCycle(L);
    if (c)
        printf("%d\n", c->val);
    return 0;
}
