#include <stdio.h>
#include <stdlib.h>

/* 
 * Definition for singly-linked list.
 *  */

struct ListNode
{
    int val;
    struct ListNode *next;
};

struct ListNode *getIntersectionNode(struct ListNode *headA, struct ListNode *headB)
{
    struct ListNode *l0, *l1;
    l0 = headA;
    l1 = headB;

    while (l0 != l1)
    {
        if (l0 != NULL)
            l0 = l0->next;
        else
            l0 = headB;
        if (l1 != NULL)
            l1 = l1->next;
        else
            l1 = headA;
    }
    return l0;
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

void intersect(struct ListNode *headA, struct ListNode *headB, int val)
{
    struct ListNode *l0 = headA, *l1 = headB;
    while (l0 && l0->next)
    {
        if (l0->next->val == val)
            break;
        l0 = l0->next;
    }

    if (!l0)
        return;

    struct ListNode *temp = l0->next;

    while (l1)
    {
        if (l1->val == val)
        {
            l0->next = l1;
            break;
        }
        l1 = l1->next;
    }

    struct ListNode *p;
    while (temp)
    {
        p = temp;
        temp = temp->next;
        p->next = NULL;
        free(p);
    }
}

int main()
{
    int arr0[] = {4, 1, 8, 4, 5}, arr1[] = {5, 0, 1, 8, 4, 5};
    int point = 8;
    struct ListNode *l0 = createLinkedList(arr0, sizeof(arr0) / sizeof(arr0[0]));
    struct ListNode *l1 = createLinkedList(arr1, sizeof(arr1) / sizeof(arr1[0]));
    intersect(l0, l1, point);

    struct ListNode *ret = getIntersectionNode(l0, l1);
    if (ret)
        printf("%d\n", ret->val);
    return 0;
}
