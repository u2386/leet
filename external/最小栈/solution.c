#include <stdlib.h>
#include <stdio.h>

#define INITIAL_CAP 10

typedef struct
{
    int *stack;
    int *min;
    int slen;
    int mlen;
    int capacity;
} MinStack;

/** initialize your data structure here. */

MinStack *minStackCreate()
{
    MinStack *s = (MinStack *)malloc(sizeof(MinStack));
    s->capacity = INITIAL_CAP;
    s->stack = (int *)malloc(s->capacity * sizeof(int));
    s->min = (int *)malloc(s->capacity * sizeof(int));
    s->slen = 0;
    s->mlen = 0;
    return s;
}

void minStackPush(MinStack *obj, int x)
{
    if (obj->slen == obj->capacity)
    {
        obj->capacity *= 2;
        obj->stack = (int *)realloc(obj->stack, obj->capacity);
        obj->min = (int *)realloc(obj->min, obj->capacity);
    }

    obj->stack[obj->slen++] = x;
    if ((obj->mlen == 0) || (x <= obj->min[obj->mlen - 1]))
    {
        obj->min[obj->mlen++] = x;
    }
}

void minStackPop(MinStack *obj)
{
    obj->slen--;
    int pop = obj->stack[obj->slen];
#ifdef DEBUG
    printf("%d\n", pop);
#endif

    if (pop == obj->min[obj->mlen - 1])
        obj->mlen--;
}

int minStackTop(MinStack *obj)
{
    return obj->stack[obj->slen - 1];
}

int minStackGetMin(MinStack *obj)
{
    return obj->min[obj->mlen - 1];
}

void minStackFree(MinStack *obj)
{
    free(obj->stack);
    free(obj->min);
    free(obj);
}

/**
 * Your MinStack struct will be instantiated and called as such:
 * MinStack* obj = minStackCreate();
 * minStackPush(obj, x);

 * minStackPop(obj);

 * int param_3 = minStackTop(obj);

 * int param_4 = minStackGetMin(obj);

 * minStackFree(obj);
*/

int main()
{
    MinStack *obj = minStackCreate();
    minStackPush(obj, 1);
    minStackPush(obj, 2);
    minStackPush(obj, 3);

    minStackPop(obj);
    minStackPop(obj);
    minStackFree(obj);
    return 0;
}
