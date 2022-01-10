// Your First C++ Program

#include <iostream>
#include <vector>
#include <string>
#include <assert.h>
#include <stack>
#include <cmath>
#define NDEBUG

// This problem was asked by Stripe.

// Given an integer n, return the length of the longest consecutive run of 1s in its binary representation.

// For example, given 156, you should return 3.

int getLognestRunOfOnes(int deci)
{
  std::cout << "deciimal = " << deci << "\n";

  std::string bin{};
  int globalMax = 0;
  int localMax = 0;

  while (deci > 0)
  {
    if (deci % 2 == 0)
      bin.insert(bin.begin(), '0');
    else
      bin.insert(bin.begin(), '1');

    deci >>= 1;
  }

  std::cout << "binary = " << bin << "\n";

  for (unsigned long int index = 0; index < bin.size(); index++)
  {
    if (bin[index] == '1')
    {
      localMax++;
    }
    else
    {
      localMax = 0;
    }

    if (globalMax < localMax)
    {
      globalMax = localMax;
    }
  }
  return globalMax;
}
int main()
{

  int n;

  //# Test Program

  //# should return False
  assert((getLognestRunOfOnes(156) == 3) && "This return 3");

  assert((getLognestRunOfOnes(5851) == 2) && "This return 2");

  assert((getLognestRunOfOnes(4465) == 3) && "This return 3");

  assert((getLognestRunOfOnes(144) == 1) && "This return 1");

  assert((getLognestRunOfOnes(521) == 1) && "This return 1");

  assert((getLognestRunOfOnes(6134) == 7) && "This return 7");
  //edge case
  assert((getLognestRunOfOnes(15798) == 4) && "This return 4");

  assert((getLognestRunOfOnes(1759) == 5) && "This return 5");

  return 0;
}
