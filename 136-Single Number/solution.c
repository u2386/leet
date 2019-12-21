#include <stdlib.h>
#include <stdio.h>

int singleNumber(int *nums, int numsSize)
{
    int ret = nums[0];
    for (int i = 1; i < numsSize; i++)
    {
        ret ^= nums[i];
    }
    return ret;
}

int main()
{
    return 0;
}