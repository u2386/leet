#include <stdlib.h>
#include <stdio.h>

#define EMPTY(heap) (heap->used == 0)
#define FULL(heap) (heap->used >= heap->capacity - 1)
#define PARENT(index) (index > 0 ? (index - 1) >> 1 : -1)
#define LCHILD(index) ((index << 1) + 1)
#define RCHILD(index) ((index << 1) + 2)

/**
 * Definition for singly-linked list.
 */
struct ListNode
{
    int val;
    struct ListNode *next;
};

typedef struct Item
{
    int val;
    struct ListNode *node;
} Item;

typedef struct Heap
{
    int used;
    int capacity;
    struct Item *items;
} Heap;

Heap *heap_init(int capacity);
int heap_push(Heap *heap, struct ListNode *node);
int heap_pop(Heap *heap, struct ListNode *node);
void shift_down(Heap *heap);
void shift_up(Heap *heap);

Item create_item(struct ListNode *node);
void swap(Item *n, Item *m);

void swap(Item *n, Item *m)
{
    Item temp = *n;
    *n = *m;
    *m = temp;
}

Heap *heap_init(int capacity)
{
    Heap *heap = (Heap *)malloc(sizeof(Heap));
    heap->used = 0;
    heap->capacity = capacity + 1;
    heap->items = (Item *)malloc(sizeof(Item) * heap->capacity);
    return heap;
}

Item create_item(struct ListNode *node)
{
    Item item = {
        .val = node->val,
        .node = node,
    };
    return item;
}

void shift_up(Heap *heap)
{
    int parent, lchild, rchild;
    Item *l = NULL, *r = NULL, *p = NULL;

    parent = PARENT(heap->used - 1);
    while (parent >= 0)
    {
        p = &heap->items[parent];

        rchild = RCHILD(parent);
        if (rchild < heap->used)
        {
            r = &heap->items[rchild];
            if (r->val < p->val)
                swap(p, r);
        }

        lchild = LCHILD(parent);
        l = &heap->items[lchild];
        if (l->val < p->val)
            swap(p, l);

        parent = PARENT(parent);
    }
}

void shift_down(Heap *heap)
{
    int parent = 0, lchild, rchild;
    int next = parent;
    Item *l = NULL, *r = NULL, *p = NULL;

    while (parent <= (heap->used << 1))
    {
        lchild = LCHILD(parent);
        rchild = RCHILD(parent);
        if (!(rchild < heap->used || lchild < heap->used))
            break;

        p = &heap->items[parent];
        if (rchild < heap->used)
        {
            r = &heap->items[rchild];
            if (p->val > r->val)
            {
                swap(p, r);
                next = rchild;
            }
        }

        if (lchild < heap->used)
        {
            l = &heap->items[lchild];
            if (p->val > l->val)
            {
                swap(p, l);
                next = lchild;
            }
        }

        if (parent == next)
            break;
        parent = next;
    }
}

int heap_push(Heap *heap, struct ListNode *node)
{
    if (FULL(heap))
        return -1;

    Item item = create_item(node);
    heap->items[heap->used++] = item;
    shift_up(heap);
    return 0;
}

int heap_pop(Heap *heap, struct ListNode *node)
{
    if (EMPTY(heap))
        return -1;

    *node = *heap->items[0].node;
    heap->items[0] = heap->items[--heap->used];
    shift_down(heap);
    return 0;
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

struct ListNode *create_node(int value)
{
    struct ListNode *node = (struct ListNode *)malloc(sizeof(struct ListNode));
    node->next = NULL;
    node->val = value;
    return node;
}

struct ListNode *create_linked(int *arr, int size)
{
    if (!size)
        return NULL;

    struct ListNode *head, *rear;
    head = create_node(arr[0]);
    rear = head;
    for (int i = 1; i < size; i++)
    {
        rear->next = create_node(arr[i]);
        rear = rear->next;
    }
    return head;
}

struct ListNode **create_node_collections(int **arr, int size)
{
    struct ListNode **ret = (struct ListNode **)malloc(sizeof(struct ListNode *) * size);
    for (int i = 0; i < size; i++)
    {
        ret[i] = create_linked(arr[i], sizeof(arr[i]) / sizeof(arr[i][0]));
    }
    return ret;
}

int main()
{
    int *arr[] = {
        (int[]){1, 4, 5},
        (int[]){1, 3, 4},
        (int[]){2, 6},
    };

    struct ListNode **data = create_node_collections(arr, 3);
    printf("%d\n", ((struct ListNode *)data)->val);

/* 
 * 
 *     int arr[] = {5, 4, 3, 2, 1};
 *     int size = sizeof(arr) / sizeof(arr[0]);
 *     Heap *h = heap_init(size);
 *     int i;
 *     struct ListNode *node;
 * 
 *     for (i = 0; i < size; i++)
 *     {
 *         node = (struct ListNode *)malloc(sizeof(struct ListNode));
 *         node->val = arr[i];
 * 
 *         if (heap_push(h, node) != 0)
 *         {
 *             printf("Heap is Full\n");
 *             goto free;
 *         }
 *     }
 * 
 *     struct ListNode n;
 *     while (heap_pop(h, &n) == 0)
 *     {
 *         printf("%d ", n.val);
 *     }
 *     printf("\n");
 * 
 * free:
 *     if (node)
 *         free(node);
 * 
 *     for (i = 0; i < h->used; i++)
 *     {
 *         free(h->items[i].node);
 *     }
 *     free(h->items);
 *     free(h);
 *     return 0;
 * 
 *  */
}
