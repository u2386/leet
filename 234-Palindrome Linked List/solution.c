#include <stdlib.h>
#include <stdio.h>
#include <stdbool.h>

/**
 * Definition for singly-linked list.
 */

struct ListNode
{
    int val;
    struct ListNode *next;
};

bool isPalindrome(struct ListNode *head)
{
    if (head == NULL || head->next == NULL)
        return true;

    struct ListNode *slow = head, *fast = head;
    struct ListNode *p0 = NULL, *p1 = NULL;
    while (fast != NULL && fast->next != NULL)
    {
        p0 = slow;
        slow = slow->next;
        fast = fast->next->next;
        p0->next = p1;
        p1 = p0;
    }

    if (fast != NULL)
        slow = slow->next;

    while (slow != NULL && p0 != NULL)
    {
        if (slow->val != p0->val)
            return false;
        slow = slow->next;
        p0 = p0->next;
    }

    return true;
}

int main()
{
    int arr[] = {1, 2, 3, 1};
    struct ListNode *head = (struct ListNode *)malloc(sizeof(struct ListNode));
    struct ListNode *prev = head, *tmp;
    for (int i = 0; i < 4; i++)
    {
        tmp = (struct ListNode *)malloc(sizeof(struct ListNode));
        tmp->val = arr[i];
        tmp->next = NULL;
        prev->next = tmp;
        prev = prev->next;
    }

    prev = head;
    for (int i = 0; i < 4; i++)
    {
        prev = prev->next;
        printf("%d ", prev->val);
    }
    printf("\n");

    int ret = isPalindrome(head->next) ? 1 : 0;
    printf("%d\n", ret);
    return 0;
}
