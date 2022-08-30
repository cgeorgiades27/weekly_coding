/*
  # Weekly Coding Problem
  
  ## Week 33, 2022
  
  This problem was recently asked by Uber:
  
  Imagine you are building a compiler. Before running any code, the compiler must check that the parentheses in the program are balanced. Every opening bracket must have a corresponding closing bracket. We can approximate this using strings.
  
  Given a string containing just the characters `'(', ')', '{', '}', '[' and ']'`, determine if the input string is valid.
  An input string is valid if:
  - Open brackets are closed by the same type of brackets.
  - Open brackets are closed in the correct order.
  - Note that an empty string is also considered valid.
  ```
  Example:
  Input: "((()))"
  Output: True
  
  Input: "[()]{}"
  Output: True
  
  Input: "({[)]"
  Output: False
  ```
  ``` python
  class Solution:
    def isValid(self, s):
      # Fill this in.
  
  # Test Program
  s = "()(){(())" 
  # should return False
  print(Solution().isValid(s))
  
  s = ""
  # should return True
  print(Solution().isValid(s))
  
  s = "([{}])()"
  # should return True
  print(Solution().isValid(s))
  ```

*/

#include <iostream>
#include <string>
#include <map>
#include <stack>
#include <assert.h>

bool CheckParen(std::string s);

int main()
{
  // Test Cases (no output is successful)
  assert(CheckParen("()(){(())")==false);
  assert(CheckParen("")==true);
  assert(CheckParen("([{}])()")==true);
  
  return 0;
}

bool CheckParen(std::string s)
{
    std::stack<char> l;
    std::map<char, char> m;
    m[')'] = '(';
    m['}'] = '{';
    m[']'] = '[';

    for (int i = 0; i < s.size(); ++i)
    {
        if (s[i] == '(' ||
            s[i] == '{' ||
            s[i] == '[')
            l.push(s[i]);

        if (s[i] == ')' ||
            s[i] == '}' ||
            s[i] == ']')
        {
            if (l.empty())
                return false;

            if (m[s[i]] == l.top())
                l.pop();
        }
    }

    return l.empty();
}