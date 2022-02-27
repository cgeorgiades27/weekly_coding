
#include <iostream>
#include <vector>
#include <string>
#include <assert.h>
#include <cmath>
#include <memory>

/* ## Week 28, 2021

This problem was asked by Google.

Given two strings A and B, return whether or not A can be shifted some number of times to get B.

For example, if A is abcde and B is cdeab, return true. If A is abc and B is acb, return false
*/

struct charLink
{

  /* data */
  int position;
  char character;
  std::shared_ptr<charLink> next;
  charLink(int pos, char character)
  {
    position = pos;
    this->character = character;
  }
};

/**
 * Function that turns a string into a circular linked list and returns the head
 *
 *
 * @param string     string to be converted

 */

std::shared_ptr<charLink> stringToCircularCharLink(std::string string)
{
  std::shared_ptr<charLink> prev;
  std::shared_ptr<charLink> Head;
  for (int i = 0; i < string.size(); i++)
  {

    auto newLink = std::make_shared<charLink>(i, string[i]);
    if (prev != nullptr)
    {
      prev->next = newLink;
    }
    prev = newLink;

    if (i == 0)
    {
      Head = newLink;
    }
  }
  prev->next = Head;
  return Head;
}

/**
 * return whether or not s1 can be shifted some number of times to get s2.
 *
 *
 * @param s1     string to be checked
 * @param s2     string to be checked

 */

bool ifShiftedEnoughCouldTheyBeEqual(std::string s1, std::string s2)
{

  auto Head = stringToCircularCharLink(s1);
  auto Head2 = stringToCircularCharLink(s2);
  while (Head2->character != Head->character)
  {
    Head2 = Head2->next;
  }
  for (int i = 0; i < s1.size(); i++)
  {
    if (Head2->character != Head->character)
    {
      return false;
    }
    Head2 = Head2->next;
    Head = Head->next;
  }
  return true;
}

int main()
{
  std::string s1 = "abcdefg";
  std::string s2 = "defgabc";
  std::string s3 = "abcdefg";
  std::string s4 = "abcdfeg";

  //# Test Program

  assert((ifShiftedEnoughCouldTheyBeEqual(s1, s2) == true) && "should be true ");
  assert((ifShiftedEnoughCouldTheyBeEqual(s3, s4) == false) && "should be true ");

  return 0;
}
