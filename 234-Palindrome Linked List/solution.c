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
        return false;

    struct ListNode *slow = head, *fast = head->next;
    while (fast != NULL && fast->next != NULL)
    {
        slow = slow->next;
        fast = fast->next->next;
    }
    return true;
}

int main()
{
    int arr[] = {1, 2, 3, 2, 1};
    struct ListNode *head = (struct ListNode *)malloc(sizeof(struct ListNode));
    struct ListNode *prev = head, *tmp;
    for (int i = 0; i < 5; i++)
    {
        tmp = (struct ListNode *)malloc(sizeof(struct ListNode));
        tmp->val = arr[i];
        tmp->next = NULL;
        prev->next = tmp;
        prev = prev->next;
    }

    prev = head;
    for (int i = 0; i < 5; i++)
    {
        prev = prev->next;
        printf("%d ", prev->val);
    }
    printf("\n");

    int ret = isPalindrome(head->next) ? 1 : 0;
    printf("%d\n", ret);
    return 0;
}
