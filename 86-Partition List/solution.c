#include <stdio.h>
#include <stdlib.h>

struct ListNode
{
    int val;
    struct ListNode *next;
};

struct ListNode *partition(struct ListNode *head, int x)
{
    if (head == NULL || head->next == NULL)
        return head;

    struct ListNode *dummy = (struct ListNode *)malloc(sizeof(struct ListNode));
    dummy->next = head;

    struct ListNode *p = dummy;
    while (p->next != NULL && p->next->val < x)
        p = p->next;

    head = p->next;

    struct ListNode *q = p;

    while (p->next != NULL)
    {
        if (p->next->val < x)
        {
            q->next = p->next;
            p->next = p->next->next;
            q = q->next;
            q->next = NULL;
        }
        else
        {
            p = p->next;
        }
    }
    q->next = head;

    head = dummy->next;
    dummy->next = NULL;
    free(dummy);
    return head;
}

int main()
{
    int arr[] = {1, 4, 3, 2, 5, 2};
    int target = 3;

    struct ListNode *dummy = (struct ListNode *)malloc(sizeof(struct ListNode));
    struct ListNode *p = dummy, *q;
    for (int i = 0; i < sizeof(arr) / sizeof(int); i++)
    {
        q = (struct ListNode *)malloc(sizeof(struct ListNode));
        q->val = arr[i];
        q->next = NULL;
        p->next = q;
        p = p->next;
    }

    struct ListNode *head = partition(dummy->next, target);
    p = head;
    q = head;
    while (p != NULL)
    {
        printf("%d ", p->val);
        p = p->next;
        q->next = NULL;
        free(q);
        q = p;
    }
    printf("\n");

    dummy->next = NULL;
    free(dummy);
    return 0;
}
