#include <stdio.h>

int search(int *nums, int numsSize, int target)
{
    int i = 0, j = numsSize - 1;

    int mid, n;
    int x, y;
    while (i <= j)
    {
        if (i == j && nums[i] != target)
            return -1;

        x = nums[i];
        y = nums[j];
        if (target == x)
            return i;
        else if (target == y)
            return j;

        mid = (i + j) >> 1;
        n = nums[mid];
        if (target == n)
            return mid;

        if (x < y)
        {
            if (target < x || target > y)
                return -1;
            else if (target > x && target < n)
                j = mid - 1;
            else
                i = mid + 1;
        }
        else
        {
            if (target < x && target > y)
                return -1;
            else if ((x < n && target > x && target < n) || (x > n && (target > x || target < n)))
                j = mid - 1;
            else
                i = mid + 1;
        }
    }
    return -1;
}

int main()
{
    int nums[] = {4, 5, 6, 7, 0, 1, 2};
    int target = 4;

    int ret = search(nums, sizeof(nums) / sizeof(nums[0]), target);
    printf("%d\n", ret);
    return 0;
}