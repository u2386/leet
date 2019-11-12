#include <stdio.h>

int max(int a, int b)
{
    return a > b ? a : b;
}

int findLengthOfLCIS(int *nums, int numsSize)
{
    if (numsSize <= 1)
    {
        return numsSize == 0 ? 0 : 1;
    }

    int max_len = 0, len = 1;

    int i = 1;
    while (i < numsSize - 1)
    {
        if (nums[i] > nums[i - 1])
        {
            len++;
            i++;
            continue;
        }
        max_len = max(max_len, len);
        len = 1;
        i += 2;
    }
    return max_len;
}

int main()
{
    int nums[] = {1, 3, 5, 7};
    printf("%d\n", findLengthOfLCIS(nums, sizeof(nums) / sizeof(nums[0])));
    return 0;
}
