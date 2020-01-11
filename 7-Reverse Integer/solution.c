#include <stdlib.h>
#include <stdio.h>
#include <limits.h>

int CEIL = INT_MAX / 10;
int FlOOR = INT_MIN / 10;

int reverse(int x)
{
    int ret = 0;
    int r;
    while (x != 0)
    {
        if (ret > CEIL || ret < FlOOR)
        {
            return 0;
        }
        r = x % 10;
        ret *= 10;
        ret += r;
        x /= 10;
    }
    return ret;
}

int main()
{
    printf("%d\n", reverse(120));
    return 0;
}
