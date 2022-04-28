/*
## Weekly Coding Problem

### Week 17, 2022 - (Easy)

This problem was asked by **Apple**.

Implement the function `fib(n)`, which returns the nth number in the Fibonacci sequence, using only `O(1)` space.
*/

#include <stdio.h>

int fib(int n);

int main()
{
int num;
  printf("enter a number: ");
  scanf("%d", &num);
  printf("num is: %d\n", fib(num));
  return 0;
}


int fib(int n)
{
  int prev2 = 1, prev1 = 1, temp = 0;

  if (n < 2)
    return 1;
  
  for (int i = 2; i <= n; ++i)
    {
      if (i == n)
        return prev1;

      temp = prev1;
      prev1 += prev2;
      prev2 = temp;
    }
}