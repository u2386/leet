#include <stdlib.h>
#include <stdio.h>

void reverse(int *nums, int start, int end)
{
    int i = start, j = end, temp;
    while (i < j)
    {
        temp = nums[i];
        nums[i++] = nums[j];
        nums[j--] = temp;
    }
}

void rotate(int *nums, int numsSize, int k)
{
    k = k % numsSize;
    reverse(nums, 0, numsSize - 1);
    reverse(nums, 0, k - 1);
    reverse(nums, k, numsSize - 1);
}

int main()
{
    int arr[] = {1, 2, 3, 4, 5, 6, 7}, k = 3;
    int size = sizeof(arr) / sizeof(arr[0]);

    rotate(arr, size, k);

    for (int i = 0; i < size; i++)
        printf("%d ", arr[i]);
    printf("\n");
    return 0;
}
