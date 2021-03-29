// Your First C++ Program

#include <iostream>
#include <vector>
#include <string>
#include <assert.h>
#include <stack>
#include <cmath>
#define NDEBUG

bool isValid(std::string s)
{

  std::stack<char> stack;

  if (s.size() % 2 != 0) //if string length is not even then false
  {
    return false;
  }
  for (long unsigned int i = 0; i < s.size(); i++)
  {

    if (s[i] == '(' || s[i] == '{' || s[i] == '[')
    {
      stack.push(s[i]);
    }
    if (stack.empty())
    {
      return false;
    }
    char tempChar = stack.top();
    //once we get to a closing char check if the item on top of stack is the oppening one if so pop it
    if (((s[i] == ')' && tempChar == '(')) || ((s[i] == '}' && tempChar == '{')) || ((s[i] == ']' && tempChar == '[')))
    {
      stack.pop();
    }
  }
  //if opening char remains in stack then it was unbalanced
  if (!stack.empty())
  {
    return false;
  }

  return true;
}

int main()
{

  std::string s;

  //# Test Program
  s = "()(){(())";

  //# should return False
  assert((isValid(s) == false) && "This should be false");

  s = "";
  //# should return True
  assert((isValid(s) == true) && "This should be true");

  s = "([{}])()";
  //# should return True
  assert((isValid(s) == true) && "This should be true");

  s = "(()()())";
  //# should return True
  assert((isValid(s) == true) && "This should be true");

  s = "())";
  //# should return False
  assert((isValid(s) == false) && "This should be false");

  s = "(*$*";
  //# should return False
  assert((isValid(s) == false) && "This should be false");

  s = "(()())()";
  //# should return True
  assert((isValid(s) == true) && "This should be true");

  return 0;
}
