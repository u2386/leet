#include <stdio.h>
#include <stdlib.h>

struct ListNode
{
    int val;
    struct ListNode *next;
};

struct ListNode *reverseList(struct ListNode *head)
{
    struct ListNode *dummy = (struct ListNode *)malloc(sizeof(struct ListNode));
    dummy->next = NULL;
    struct ListNode *r;
    while (head != NULL)
    {
        r = head;
        head = head->next;
        r->next = dummy->next;
        dummy->next = r;
    }

    r = dummy->next;
    dummy->next = NULL;
    free(dummy);
    return r;
}

struct ListNode *create_node(int val)
{
    struct ListNode *r = (struct ListNode *)malloc(sizeof(struct ListNode));
    r->val = val;
    r->next = NULL;
    return r;
}

int main()
{
    struct ListNode *head;
    head = create_node(1);
    head->next = create_node(2);
    head->next->next = create_node(3);

    struct ListNode *r = reverseList(head);
    struct ListNode *p;
    while (r != NULL)
    {
        p = r;
        r = r->next;
        printf("%d\n", p->val);
        p->next = NULL;
        free(p);
    }
    return 0;
}
