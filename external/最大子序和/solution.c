#include <stdlib.h>
#include <stdio.h>
#include <limits.h>

static inline int max(int a, int b, int c)
{
    return a > b
               ? a > c ? a : c
               : b > c ? b : c;
}

int _maxSubArray(int *nums, int start, int end)
{
    if (end == start)
        return nums[start];

    int mid = (start + end) >> 1;
    int leftMax = _maxSubArray(nums, start, mid);
    int rightMax = _maxSubArray(nums, mid + 1, end);

    int crossingLeftMax = INT_MIN;
    int prev = 0;
    for (int i = mid; i > start - 1; i--)
    {
        prev += nums[i];
        crossingLeftMax = crossingLeftMax > prev ? crossingLeftMax : prev;
    }
    int crossingRightMax = INT_MIN;
    prev = 0;
    for (int i = mid + 1; i < end + 1; i++)
    {
        prev += nums[i];
        crossingRightMax = crossingRightMax > prev ? crossingRightMax : prev;
    }

    return max(leftMax, (crossingLeftMax + crossingRightMax), rightMax);
}

int maxSubArray(int *nums, int numsSize)
{
    if (numsSize == 1)
        return nums[0];
    else if (numsSize == 2)
    {
        int sum = nums[0] + nums[1];
        return max(sum, nums[0], nums[1]);
    }
    return _maxSubArray(nums, 0, numsSize - 1);
}

int main()
{
    int arr[] = {-2, 1, -3, 4, -1, 2, 1, -5, 4};
    int m = maxSubArray(arr, sizeof(arr) / sizeof(arr[0]));
    printf("%d\n", m);
    return 0;
}
