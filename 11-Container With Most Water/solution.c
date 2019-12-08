#include <stdlib.h>
#include <stdio.h>

static inline int max(int a, int b)
{
    return a > b ? a : b;
}

static inline int min(int a, int b)
{
    return a < b ? a : b;
}

int maxArea(int *height, int heightSize)
{
    int i = 0, j = heightSize - 1;
    int max_water = 0;
    int left, right, h;
    while (i < j)
    {
        left = height[i];
        right = height[j];
        h = min(left, right);
        max_water = max(max_water, (j - i) * h);

        if (h == left)
        {
            i++;
            while (i < j && height[i] < height[i - 1])
                i++;
        }
        else
        {
            j--;
            while (i < j && height[j] < height[j + 1])
                j--;
        }
    }
    return max_water;
}

int main()
{
    int arr[] = {1, 8, 6, 2, 5, 4, 8, 3, 7};
    /* int arr[] = {1, 2, 4, 3}; */
    int w = maxArea(arr, sizeof(arr) / sizeof(arr[0]));
    printf("%d\n", w);
    return 0;
}
