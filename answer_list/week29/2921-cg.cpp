/*
## Weekly Coding Problem

### Week 29, 2021

This problem was asked by Microsoft.

Let's represent an integer in a linked list format by having each node represent a digit in the number. The nodes make up the number in reversed order.

For example, the following linked list:

`1 -> 2 -> 3 -> 4 -> 5`
is the number `54321`.

Given two linked lists in this format, return their sum in the same linked list format.

For example, given

`9 -> 9`
`5 -> 2`
return `124 (99 + 25)` as:

`4 -> 2 -> 1`
*/

#include <iostream>
#include <list>

typedef std::list<int> list;

int getNumFromList(list&);
list getListFromNum(int&);

int main()
{
  list l1 {9, 9};
  list l2 {5, 2};
  int sum = getNumFromList(l1) + getNumFromList(l2);
  list finalList = getListFromNum(sum);

  std::cout << "The list elements are: ";
  for (auto i : finalList)
    std::cout << i << ' ';
  std::cout << '\n';
  return 0;
}

int getNumFromList(list& l)
{
  int num = 0;
  int count = 1;
  for (auto i : l)
    {
      num += count * i;
      count *= 10;
    }
  return num;  
}

list getListFromNum(int& num)
{
  list l;
  int divisor = 1;
  
  for (int i = 0;  num / divisor >= 1; ++i, divisor *= 10)
      l.push_back(num / divisor % 10);
  
  return l;
}
