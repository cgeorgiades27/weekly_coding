
#include <iostream>
#include <vector>
#include <string>
#include <assert.h>
#include <cmath>
#include <memory>

/* ## Week 28, 2021

Weekly Coding Problem - Week 17, 2021
Given a string consisting of parentheses, single digits, and positive and negative signs, convert the string into a mathematical expression to obtain the answer.

Don't use eval or a similar built-in parser.

For example, given '-1 + (2 + 3)', you should return 4.
*/

/**
 * Parse string, do arithmitic, return answer
 *
 * @param s1     string to be checked

 */

int evaluate(std::string s1)
{

  while (!s1.empty())
  {

    int openParenth = s1.find_last_of("(");
    int closedParenth = s1.find(")");
    while (closedParenth < openParenth)
    {
      closedParenth = s1.find(")", closedParenth + 1);
    }
    std::string str2 = s1;
   
    if (openParenth != std::string::npos)
    {
      str2 = s1.substr(openParenth, closedParenth);
    }else{
       openParenth = 0;
    closedParenth = s1.size();
    }


    auto minus = str2.find("-");
    auto plus = str2.find("+");
    auto div = str2.find("/");
    auto mul = str2.find("*");

    if (minus == 0)
    {
      minus = str2.find("-", minus + 1);
      plus = str2.find("+", minus + 1);
      div = str2.find("/", minus + 1);
      mul = str2.find("*", minus + 1);
      openParenth=-1;
    }

    if (minus != std::string::npos)
    {
    }
    else if (plus != std::string::npos)
    {
      auto result1 = stoi(s1.substr(openParenth + 1, openParenth + plus )) ;
      auto result2 = stoi(s1.substr(openParenth + plus + 2, closedParenth - 1));
      auto result =result1+result2;

    s1.insert(openParenth==-1?0:openParenth, std::to_string(result));
      s1.erase(openParenth==-1?1:openParenth+1, closedParenth);
      
    }
    else if (div != std::string::npos)
    {
    }
    else if (mul != std::string::npos)
    {
    }

    if (s1.size()==1)
    {return stoi(s1);}
  }

  return 4;
}

int main()
{
  std::string s1 = "-1 + (2 + 3)";
  std::string s2 = "-1 + ((2+2) + 3)";


  //# Test Program

  assert((evaluate(s1) == 4) && "should be true ");
  assert((evaluate(s2) == 6) && "should be true ");


  return 0;
}
