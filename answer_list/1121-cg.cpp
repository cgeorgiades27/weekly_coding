// week 11, 2021

#include <iostream>
#include <map>
#include <stack>
#include <string>

bool isValid(std::string);

int main()
{
 
  std::string s1, s2, s3;
  
  s1 = "()(){(())";
  // should return False
  std::cout << "Result for " << s1 << " is: " << isValid(s1) << '\n';
  
  s2 = "";
  // should return True
  std::cout << "Result for " << s2 << " is: " << isValid(s2) << '\n';
  
  s3 = "([{}])()";
  // should return True
  std::cout << "Result for " << s3 << " is: " << isValid(s3) << '\n';

  return 0;
}

bool isValid(std::string s)
{
  // no brackets
  if (s.size() == 0)
    return true;
  
  std::stack<char> stack_;
  std::map<char, char> map_;
  map_[']'] = '[';
  map_[')'] = '(';
  map_['}'] = '{';
  
  for (size_t i = 0; i < s.size(); ++i)
    {
      if ((s[i] == ')' || s[i] == '}' || s[i] == ']') && !stack_.empty())
        {
	  if (stack_.top() != map_[s[i]] || stack_.empty())
	    return false;
	  stack_.pop();
        }
      else
	stack_.push(s[i]);
    }
  // leftovers...
  if (!stack_.empty())
    return false;
  return true;
}
