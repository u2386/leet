#include <stdio.h>

double myPow(double x, int n)
{
    if (n == 0)
    {
        return 1.0;
    }
    if (x == 1 || x == 0 || n == 1)
    {
        return x;
    }

    long N = n;
    if (n < 0)
    {
        x = 1 / x;
        N *= -1;
    }
    double ret = 1;
    while (N > 0)
    {
        if ((N & 1) == 1)
        {
            ret *= x;
        }
        x *= x;
        N >>= 1;
    }
    return ret;
}

int main()
{
    printf("%f\n", myPow(2.00000, -2147483648));
    return 0;
}
