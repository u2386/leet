#include <string.h>
#include <stdio.h>

void revert(char *start, char *end)
{
    char c;
    while (start < end)
    {
        c = *start;
        *start = *end;
        *end = c;
        start++;
        end--;
    }
}

char *reverseWords(char *s)
{
    if (s == NULL)
        return NULL;

    char *i, *j;
    int l = strlen(s);

    /* Revert the string */
    revert(s, s + l - 1);

    i = s;
    j = s;
    while (*i != '\0')
    {
        if (*i == ' ')
        {
            i++;
            continue;
        }

        /* Move to one character after the word */
        while (!(*i == ' ' || *i == '\0'))
            i++;

        /* Revert one word */
        revert(j, i - 1);

        j = i - 1;
        while (*j == ' ')
            j--;
        j += 2;
    }

    if (j > s)
        j--;
    *j = '\0';
    return s;
}

int main()
{
    char s[] = "a good   example";
    char *p = reverseWords(s);

    for (int i = 0; i < strlen(p); i++)
    {
        printf("%c", p[i]);
    }
    printf("$");
    printf("\n");
    return 0;
}