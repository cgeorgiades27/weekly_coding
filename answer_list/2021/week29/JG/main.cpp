
#include <iostream>
#include <string>
#include <assert.h>
#include <cmath>
#include <list>

/* Week 29, 2021
This problem was asked by Microsoft.

Let's represent an integer in a linked list format by having each node represent a digit in the number. The nodes make up the number in reversed order.

For example, the following linked list:

1 -> 2 -> 3 -> 4 -> 5 is the number 54321.

Given two linked lists in this format, return their sum in the same linked list format.

For example, given

9 -> 9 5 -> 2 return 124 (99 + 25) as:

4 -> 2 -> 1*/

/**
 * Function that returns sum of two linked list in the same linked list format.
 *
 * 
 * @param list1     fist linked list 
 * @param list2     second linked list
 
 */

std::list<int> getSum(std::list<int> &list1, std::list<int> &list2)
{
  std::list<int> returnMe;

  std::string stringRepOfList1;
  std::string stringRepOfList2;

  for (std::list<int>::reverse_iterator it = list1.rbegin(); it != list1.rend(); it++) //will traverse list1, get the value, and convert it into a string representation
  {
    stringRepOfList1 += std::to_string(*it);
  }
  for (std::list<int>::reverse_iterator it = list2.rbegin(); it != list2.rend(); it++) //will traverse list2, get the value, and convert it into a string representation
  {
    stringRepOfList2 += std::to_string(*it);
  }

  std::string Sum = std::to_string(stoi(stringRepOfList1) + stoi(stringRepOfList2)); // will convert the string to ints to we can add them

  // Traverse the string and push the value into the final list
  for (int i = 0; i < (int)Sum.size(); i++)
  {

    // Print current character in list
    returnMe.push_front((int)Sum[i] - 48); //will push back an ascii char (cast into a int), this is why 48 is the offset
  }

  return returnMe;
}

int main()
{
  std::list<int> List1 = {9, 9};
  std::list<int> List2 = {5, 2};
  std::list<int> answer = {4, 2, 1};

  getSum(List1, List2);

  //# Test Program

  assert((getSum(List1, List2) == answer) && "answer should be 4 -> 2 -> 1");

  List1 = {2, 8, 5, 1};
  List2 = {2, 1};
  answer = {4, 9, 5, 1};
  assert((getSum(List1, List2) == answer) && "answer should be 4 -> 9 -> 5 -> 1");

  List1 = {3, 3, 4};
  List2 = {4, 5, 9};
  answer = {7, 8, 3, 1};
  assert((getSum(List1, List2) == answer) && "answer should be 7 -> 8 -> 3 -> 1");

  return 0;
}
