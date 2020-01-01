#include <stdio.h>

int strStr(char *haystack, char *needle)
{
    if (*needle == '\0')
        return 0;

    char *head = haystack;
    while (*haystack != '\0')
    {
        int offset = 0;

        while (1)
        {
            if (needle[offset] == '\0')
                return haystack - head;
            if (haystack[offset] == '\0')
                return -1;
            if (haystack[offset] != needle[offset])
                break;
            offset++;
        }

        haystack++;
    }
    return -1;
}

int main()
{
    char p[] = "hello";
    char q[] = "ll";
    printf("%d\n", strStr(p, q));
    return 0;
}
