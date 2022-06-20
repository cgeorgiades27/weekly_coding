
/*****************
 * WEEK 24, 2022 *
 *****************/

#include <iostream>
#include <vector>

typedef std::vector<int> vec;

vec a1{1, 2, 3, 4, 5};
vec a2{3, 2, 1};

vec getArrSum(vec &);

int main()
{
  vec response = getArrSum(a1);
  for (auto i : response)
    std::cout << i << std::endl;

  return 0;
}

vec getArrSum(vec &v)
{
  int totalSum = 1;
  for (auto i : v)
    totalSum *= i;

  vec res;
  for (auto i : v)
    res.push_back(totalSum / i);

  return res;
}
