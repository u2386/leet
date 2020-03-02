#include <stdio.h>
#include <stdlib.h>

struct ListNode
{
    int val;
    struct ListNode *next;
};

struct ListNode *reverseList(struct ListNode *head)
{
    if (head == NULL || head->next == NULL)
        return head;
    struct ListNode *r = reverseList(head->next);
    head->next->next = head;
    head->next = NULL;
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
