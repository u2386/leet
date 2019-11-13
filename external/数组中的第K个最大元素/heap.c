#include <stdio.h>
#include <stdlib.h>

#define PARENT(index) ((index - 1) >> 1)
#define LCHILD(index) ((index << 1) + 1)
#define RCHILD(index) ((index << 1) + 2)

#define FULL(heap) (heap->len >= heap->capacity)
#define EMPTY(heap) (heap->len == 0)

typedef struct heap
{
    int len;
    int capacity;
    int *arr;
} heap;

void shiftup(heap *h, int index);
void shiftdown(heap *h, int index);
heap *heap_init(int capacity);
int heap_push(heap *h, int item);
int heap_pop(heap *h, int *pop);

static inline int *parent(heap *h, int index)
{
    return PARENT(index) >= 0 ? &h->arr[PARENT(index)] : NULL;
}

static inline int *lchild(heap *h, int index)
{
    return LCHILD(index) < h->len ? &h->arr[LCHILD(index)] : NULL;
}

static inline int *rchild(heap *h, int index)
{
    return RCHILD(index) < h->len ? &h->arr[RCHILD(index)] : NULL;
}

heap *heap_init(int capacity)
{
    heap *h = (heap *)malloc(sizeof(heap));
    h->capacity = capacity;
    h->len = 0;
    h->arr = (int *)malloc(sizeof(int) * (capacity + 1));
    if (h->arr == NULL)
        return NULL;
    return h;
}

static inline void swap(int *a, int *b)
{
    int tmp = *a;
    *a = *b;
    *b = tmp;
}

int heap_push(heap *h, int item)
{
    if (FULL(h))
        return -1;
    h->arr[h->len] = item;
    shiftup(h, h->len++);
    return 0;
}

int heap_pop(heap *h, int *pop)
{
    if (EMPTY(h))
        return -1;
    *pop = h->arr[0];
    h->arr[0] = h->arr[--h->len];
    shiftdown(h, 0);
    return 0;
}

void shiftup(heap *h, int index)
{
    int *p = parent(h, index);
    while (p != NULL && *p > h->arr[index])
    {
        swap(p, &h->arr[index]);
        index = PARENT(index);
        p = parent(h, index);
    }
}

void shiftdown(heap *h, int index)
{
    int next;
    int *left = lchild(h, index), *right = rchild(h, index), *min;
    while (left != NULL || right != NULL)
    {
        if (right != NULL && *left < *right)
        {
            min = left;
            next = LCHILD(index);
        }
        else
        {
            min = right;
            next = RCHILD(index);
        }
        if (min == NULL || *min > h->arr[index])
            break;
        swap(min, &h->arr[index]);

        index = next;
        left = lchild(h, index);
        right = rchild(h, index);
    }
}

int *heapify(int *arr, int size)
{
    int i;
    heap *h = heap_init(size);
    for (i = 0; i < size; i++)
        heap_push(h, arr[i]);

    int pop;
    int *narr = (int *)malloc(sizeof(int) * size);
    for (i = 0; i < size; i++)
    {
        heap_pop(h, &pop);
        narr[i] = pop;
    }
    return narr;
}

int main()
{
    int arr[] = {90, 15, 10, 7, 12, 2, 7, 3};
    int size = sizeof(arr) / sizeof(int);
    int *n = heapify(arr, size);

    for (int i = 0; i < size; i++)
    {
        printf("%d ", n[i]);
    }
    printf("\n");
    return 0;
}
