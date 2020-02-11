#include <stdlib.h>
#include <stdio.h>

struct ListNode
{
    int val;
    struct ListNode *next;
};

struct ListNode *swapPairs(struct ListNode *head)
{
    struct ListNode *dummy = (struct ListNode *)malloc(sizeof(struct ListNode));
    dummy->next = head;
    struct ListNode *x = dummy, *y = dummy->next;
    while (x->next != NULL && y->next != NULL)
    {
        x->next = y->next;
        y->next = x->next->next;
        x->next->next = y;

        x = x->next->next;
        y = x->next;
    }

    head = dummy->next;
    free(dummy);
    return head;
}

int main()
{
    int arr[] = {1};
    struct ListNode *head = (struct ListNode *)malloc(sizeof(struct ListNode));
    struct ListNode *tmp = head;
    for (int i = 0; i < sizeof(arr) / sizeof(arr[0]); i++)
    {
        tmp->next = (struct ListNode *)malloc(sizeof(struct ListNode));
        tmp = tmp->next;
        tmp->next = NULL;
        tmp->val = arr[i];
    }
    tmp = head;
    head = head->next;
    free(tmp);

    head = swapPairs(head);

    tmp = head;
    while (tmp != NULL)
    {
        printf("%d\n", tmp->val);
        tmp = tmp->next;
    }

    while (head != NULL)
    {
        tmp = head;
        head = head->next;
        free(tmp);
    }
    return 0;
}
