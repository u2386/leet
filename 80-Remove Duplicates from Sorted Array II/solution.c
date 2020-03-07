#include <stdio.h>

int removeDuplicates(int *nums, int numsSize)
{
    if (numsSize <= 2)
        return numsSize;

    int counter = 0;
    int i = 1;
    for (int j = 1; j < numsSize; j++)
    {
        if (nums[j] != nums[j - 1])
        {
            counter = 0;
            nums[i++] = nums[j];
        }
        else if (counter < 1)
        {
            counter++;
            nums[i++] = nums[j];
        }
    }
    return i;
}

int main()
{
    int arr[] = {1, 1, 1, 2, 2, 3};
    int r = removeDuplicates(arr, sizeof(arr) / sizeof(arr[0]));
    for (int i = 0; i < r; i++)
    {
        printf("%d ", arr[i]);
    }
    printf("\n");
    return 0;
}
