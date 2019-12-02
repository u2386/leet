#include <stdlib.h>
#include <stdio.h>
#include <stdbool.h>

bool validUtf8(int *data, int dataSize)
{
    int arr[4][2] = {
        {0, 127},
        {192, 223},
        {224, 239},
        {240, 247}};

    int i = 0;
    while (i < dataSize)
    {
        int header = data[i];

        int j = 0;
        for (; j < 4; j++)
        {
            if (arr[j][0] <= header && header <= arr[j][1])
                break;
        }

        if (j == 0)
        {
            i++;
            continue;
        }
        else if (j == 4)
        {
            return false;
        }

        if (i + j >= dataSize)
        {
            return false;
        }

        i++;
        for (; j > 0; j--)
        {
            if ((data[i++] >> 6) ^ 2)
            {
                return false;
            }
        }
    }
    return true;
}

int main()
{
    int arr[] = {197, 130, 1};
    printf("%d\n", validUtf8(arr, sizeof(arr) / sizeof(arr[0])) ? 1 : 0);
    return 0;
}
