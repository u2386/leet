#include <stdio.h>
#include <stdlib.h>

int maxProfit(int *prices, int pricesSize)
{
    if (pricesSize < 1)
        return 0;

    int profit = 0, min = 100000;
    int price;
    for (int i = 0; i < pricesSize; i++)
    {
        price = prices[i];
        if (price < min)
            min = price;
        else if (price - min > profit)
            profit = price - min;
    }
    return profit;
}

int main()
{
    int a[] = {7, 1, 5, 3, 6, 4};
    printf("%d\n", maxProfit(a, sizeof(a) / sizeof(a[0])));
    return 0;
}
