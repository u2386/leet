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

struct ListNode *partion(struct ListNode *head)
{
    struct ListNode *slow, *fast, *ret;
    slow = head;
    fast = slow->next;

    while (fast && fast->next)
    {
        slow = slow->next;
        fast = fast->next->next;
    }
    ret = slow->next;
    slow->next = NULL;
    return ret;
}

struct ListNode *sortList(struct ListNode *head)
{
    if (!head || !head->next)
        return head;

    struct ListNode *l = head;
    struct ListNode *r = partion(head);

    l = sortList(l);
    r = sortList(r);

    struct ListNode *L = NULL, *rear;
    if (l->val < r->val)
    {
        rear = l;
        l = l->next;
    }
    else
    {
        rear = r;
        r = r->next;
    }
    L = rear;
    rear->next = NULL;

    while (l && r)
    {
        if (l->val < r->val)
        {
            rear->next = l;
            l = l->next;
        }
        else
        {
            rear->next = r;
            r = r->next;
        }
        rear = rear->next;
    }

    if (l)
        rear->next = l;
    if (r)
        rear->next = r;

    r = NULL;
    l = NULL;
    return L;
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
    int arr[] = {4, 2, 1, 3};
    struct ListNode *L = createLinkedList(arr, sizeof(arr) / sizeof(int));

    /* while (L)
     * {
     *     printf("%d ", L->val);
     *     L = L->next;
     * }
     * printf("\n\n");
     * */

    L = sortList(L);
    while (L)
    {
        printf("%d ", L->val);
        L = L->next;
    }
    printf("\n");

    return 0;
}
