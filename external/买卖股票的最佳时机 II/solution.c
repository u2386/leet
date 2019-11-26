#include <stdio.h>
#include <stdlib.h>

int maxProfit(int *prices, int pricesSize)
{
    if (pricesSize < 2)
        return 0;

    int total = 0, i = 0, valley;
    while (i + 1 < pricesSize)
    {
        if (prices[i] >= prices[i + 1])
        {
            i++;
            continue;
        }

        valley = prices[i++];
        if (i >= pricesSize)
            break;

        while (i + 1 < pricesSize && prices[i] <= prices[i + 1])
            i++;
        total += prices[i++] - valley;
    }
    return total;
}

int main()
{
    int arr[] = {7, 1, 5, 3, 6, 4};
    printf("%d\n", maxProfit(arr, sizeof(arr) / sizeof(arr[0])));
    return 0;
}
