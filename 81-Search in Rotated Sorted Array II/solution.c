#include <stdlib.h>
#include <stdio.h>
#include <stdbool.h>

bool search(int *nums, int numsSize, int target)
{
    int left = 0, right = numsSize - 1, mid;
    while (left <= right)
    {
        if (nums[left] == target)
            return true;

        if (nums[0] == nums[right])
        {
            right--;
            continue;
        }

        while (left <= right && left < numsSize - 1 && nums[left] == nums[left + 1])
            left++;

        while (left <= right && right >= 0 && nums[right] == nums[right - 1])
            right--;

        mid = (left + right) >> 1;

        if (nums[mid] == target)
            return true;

        if (target < nums[mid])
        {
            if (nums[left] <= target || nums[mid] < nums[right])
                right = mid - 1;
            else
                left = mid + 1;
        }
        else
        {
            if (target <= nums[right] || nums[left] < nums[mid])
                left = mid + 1;
            else
                right = mid - 1;
        }
    }
    return false;
}

int main()
{
    int arr[] = {1, 3, 5};
    int target = 3;
    printf("%d\n", search(arr, sizeof(arr) / sizeof(int), target) ? 1 : 0);
    return 0;
}
