

/*

# Week 03, 2023 - Convert to Base Two

Given a non-negative integer `n`, convert `n` to base 2 in string form. Do not use any built in base 2 conversion functions like bin.

Here's an example and some starter code:

```python
def base_2(n):
  # Fill this in.

print(base_2(123))
# 1111011
```
In the above example, 2^6 + 2^5 + 2^4 + 2^3 + 2^1 + 2^0 = 123.

*/

#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>

void base_2(char *, int);

int main(void)
{
    int n = 123;
    char *buff = malloc((int)sizeof(n) * 8 + 1); // max size is 32 bits
    base_2(buff, n);
    printf("%s\n", buff);
    free(buff);

    return 0;
}

void base_2(char *buff, int n)
{
    int sz = (int)sizeof(n) * 8;
    char *str = malloc(sz + 1);

    if (str == NULL)
    {
        printf("error: malloc failed\n");
        exit(0);
    }

    int i = sz - 1;
    int j = 0;
    bool leading_zero = true; // ignore leading zeros
    while (i >= 0)
    {
        int r = (n >> i) & 1;
        if (r == 1)
            leading_zero = false;

        if (!leading_zero)
        {
            char *c = malloc(1);
            sprintf(c, "%d", r);
            str[j] = *c;
            ++j;
        }
        --i;
    }
    str[j] = '\0';

    strncpy(buff, str, sz - 1);
    free(str);
}
