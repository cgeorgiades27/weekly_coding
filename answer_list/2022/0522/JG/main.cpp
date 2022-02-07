
#include <iostream>
#include <vector>
#include <string>
#include <assert.h>
#include <cmath>
#include <memory>

/* ## Week 5, 2022

return the longest palindrom in a given stting


/**
 * Function that returns the best position to stand to be the last person standing.
 *
 * 
 * @param k     first to die i.e place to start 
 * @param n     amount of prisoners
 
 */

std::string findLongestPalindrome(std::string parseMe)
{
  int counter = 0;
  std::string returnMe = "";
  for (auto currentChar : parseMe)
  {
    std::cout << currentChar << std::endl;
    counter = 1; //start checking from the last char 
    if (currentChar != parseMe.at(parseMe.length() - counter))
      returnMe = "";

    do
    {
      counter++;
      if (currentChar == parseMe.at(parseMe.length() - counter))
        returnMe += currentChar;
      break;
    } while (parseMe.length() - counter != 1);
  }
  std::cout << returnMe;
  return returnMe;
}

int main()
{

  //# Test Program

  std::string test = "adabbae";

  assert((findLongestPalindrome(test) == "abba") && "aaa should be longest palindrome");

  return 0;
}
