/*
Week 06, 2023 - Pascal's Triangle

Pascal's Triangle is a triangle where all numbers are the sum of the two numbers above it. Here's an example of the Pascal's Triangle of size `5`.

```
     1
    1 1
   1 2 1
  1 3 3 1
 1 4 6 4 1
1 5 10 10 5 1
```

Given an integer `n`, generate the n-th row of the Pascal's Triangle.

Here's an example and some starter code.

```python
def pascal_triangle_row(n):
  # Fill this in.

print(pascal_triangle_row(6))
# [1, 5, 10, 10, 5, 1]
*/

#include <stdio.h>
#include <stdlib.h>

void pascal_triangle_row(int n);

int main()
{
    int n = 10;
    for (int i = 0; i < n; i++)
        pascal_triangle_row(i);
    return 0;
}

void pascal_triangle_row(int n)
{
    int arr[n][n];
    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j <= i; j++)
        {
            if (j == 0 || j == i)
                arr[i][j] = 1;
            else
                arr[i][j] = arr[i - 1][j - 1] + arr[i - 1][j];
        }
    }

    for (int i = 0; i < n; i++)
        printf("%d ", arr[n - 1][i]);
    printf("\n");
}