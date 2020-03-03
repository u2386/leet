#include <stdio.h>

int removeDuplicates(int *nums, int numsSize)
{
    if (numsSize <= 1)
        return numsSize;

    int i = 0;
    for (int j = 1; j < numsSize; j++)
    {
        if (nums[i] != nums[j])
        {
            nums[++i] = nums[j];
        }
    }
    return i + 1;
}

int main()
{
    int arr[] = {0, 0, 1, 1, 1, 2, 2, 3, 3, 4};
    int size = sizeof(arr) / sizeof(arr[0]);
    printf("%d\n", removeDuplicates(arr, size));
    for (int i = 0; i < size; i++)
    {
        printf("%d ", arr[i]);
    }
    printf("\n");
    return 0;
}
