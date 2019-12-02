#include <stdlib.h>
#include <stdio.h>

int mySqrt(int x)
{
    if (x <= 1)
        return x;

    long low = 1, high = x, mul, mid;
    while (low < high)
    {
        mid = (low + high) >> 1;
        mul = mid * mid;
        if (mul > x)
            high = mid -1;
        else if (mul < x)
            low = mid+1;
        else
            return mid;
    }
    if (high * high <= x)
        return high;
    return high - 1;
}

int main()
{
    printf("%d\n", mySqrt(3));
    return 0;
}
