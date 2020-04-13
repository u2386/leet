#include <stdio.h>

int removeElement(int *nums, int numsSize, int val)
{
    int i = 0, ret = numsSize;
    for (int j = 0; j < numsSize; j++)
    {
        if (nums[j] != val)
        {
            nums[i++] = nums[j];
        }
        else
        {
            ret--;
        }
    }
    return ret;
}

int main()
{
    int arr[] = {0, 1, 2, 2, 3, 0, 4, 2};
    int val = 2;
    int len = removeElement(arr, sizeof(arr) / sizeof(int), val);
    for (int i = 0; i < len; i++)
    {
        printf("%d\n", arr[i]);
    }
    return 0;
}
